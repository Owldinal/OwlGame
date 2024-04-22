// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/access/AccessControl.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Burnable.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721Burnable.sol";
import "@openzeppelin/contracts/utils/ReentrancyGuard.sol";

import "./Owldinal.sol";
import "./OwldinalGenOneBox.sol";
import "./Utils.sol";

contract OwlGame is AccessControl, ReentrancyGuard {
    struct OwlStakingInfo {
        uint256 tokenId;
        address owner;
        uint64 stakingTime;
        uint256[] buffedTargetIds;
    }

    struct TokenStakingInfo {
        uint256 tokenId;
        uint256 reward;
        address owner; // 20 bytes
        uint64 stakingTime; // 8 bytes
        BoxType boxType; // 1 byte
        uint16 buffLevel; // 2 bytes (total 31)
    }

    struct Rebate {
        uint256 totalRebateEarned;
        uint256 rebatePendingWithdrawal;
        uint256 unlockedRebateToClaim;
        uint256 mintedBoxCount;
    }

    // event MintBox(address indexed user, uint256 tokenId);
    // event OpenBox(address indexed user, uint256 tokenId, BoxType boxType);
    event PrizePoolIncreased(uint256 amount);
    event PrizePoolDecreased(uint256 amount);
    event JoinGame(address indexed user, uint32 inviteCode);
    event BindInvitation(address indexed invitee, address inviter);

    event RequestMint(address indexed user, uint256 requestId, uint256 count);

    event MintMysteryBox(
        address indexed user,
        uint256 requestId,
        uint256 count,
        uint256[] tokenId
    );
    event StakeOwldinalNft(address indexed user, uint256[] tokenId);
    event StakeMysteryBox(
        address indexed user,
        uint256[] tokenId,
        uint16 buffLevel
    );
    event UnstakeOwldinalNft(address indexed user, uint256[] tokenId);
    // This event cannot accept batch parameters, otherwise it would be impossible to
    // distinguish the rewards for each specific token.
    event UnstakeMysteryBox(
        address indexed user,
        uint256 tokenId,
        BoxType boxType,
        uint256 rewards
    );
    // If mintCount=0, it indicates that the burn occurred during the claim elf.
    event OwlTokenBurned(address user, uint256 mintCount, uint256 amount);

    event RebateRewardsIncreased(
        address indexed user,
        address invitee,
        uint256 mintCount,
        uint256 amount
    );
    event UnlockableRebateIncreased(
        address indexed user,
        uint256 mintCount,
        uint256 amount
    );
    event RebateClaimed(address indexed user, uint256 amount);

    event FruitRewardUpdated(
        uint256 amount,
        uint256 count,
        uint256 totalFruitCount,
        uint256 totalElfCount
    );
    event ElfRewardUpdated(uint256 amount, uint256 count);

    bytes32 public constant SERVER_ROLE = keccak256("SERVER_ROLE");

    ERC20Burnable public owlToken;
    Owldinal public owldinalNftContract;
    OwldinalGenOneBox public mysteryBoxContract;

    // 3 hours and 55 minutes
    uint256 public constant FRUIT_REWARD_INTERVAL = 3600 * 4 - 300;
    // The price that needs to be spent for each mint.
    uint256 public constant MINT_PRICE = 100000 * 10 ** 18;
    uint256 public constant MINT_REBATE = 10000 * 10 ** 18;

    // The reward proportion for Fruit, scaled by a factor of 1,000,000.
    // This value starts at 0 and is calculated by the _calculateFruitRewardsProportion
    // method based on the current staked amount of fruit. When the value is 0, it indicates
    // that the rewards proportion has not been set yet.
    uint256 private globalFruitRewardsProportion = 0;

    // Owl Token Id -> OwlStakingInfo
    mapping(uint256 => OwlStakingInfo) owlInfoMap;

    // owner address -> owl token id list
    mapping(address => uint256[]) stakedOwldinalsByOwner;
    mapping(uint256 => TokenStakingInfo) public tokenInfoMap;

    uint256[] fruitIdList;
    uint256[] elfIdList;

    // mystery box token id -> buffing owldinals id
    mapping(uint256 => uint256[]) buffingOwlsByMysteryBox;

    // inviter address -> invitees
    mapping(address => address[]) public inviterToInviteesMap;
    // invitees address -> inviter
    mapping(address => address) public inviteeToInviterMap;

    // invite code -> inviter (find invter from code)
    mapping(uint32 => address) public inviteCodeToInviterMap;
    // inviter -> invite code (check if first join game)
    mapping(address => uint32) public inviterToInviteCodeMap;

    // inviter -> unclaimed rebate
    mapping(address => Rebate) public inviterRebateMap;

    // prize pool
    uint256 public prizePool;

    struct MintRequest {
        address user;
        uint256 count;
    }
    uint256 mintRequestCounter;
    // request id => mint request
    mapping(uint256 => MintRequest) public requestMintMap;

    bool public isMoonBoostEnable;
    address[] private moonBoostWhiteList = [
        address(0x0012b4b8B7E8f5B69b3BDb5DE56Fb19A6464894C),
        address(0xD0004696e83f1D6614f4d7BF3392B6005B6DE3Ec),
        address(0xceAc49a86AdB77D322d3F9fE435397982C942c00),
        address(0xE87d3CE5740062aDfa6B09cc7A8Ef600A732A3bD),
        address(0x44ea107Df684040f89f42f250997901269EBAB68),
        address(0xDb4fcEE6212E9c16d099B2264E628A92044ceF5e),
        address(0x3C8a4EdfE35b4eEAAD9A06e5c93ad954A8672868),
        address(0xDcf358aE5A6B6E71cfC0756a0706978F1AEFa091),
        address(0xB93Cb37A3207c09a0dC1afbf6F2D1914450E6f2E),
        address(0xfa4C794e070FDA97d3A57e8b274845066Cc2c107)
    ];

    constructor(address server) {
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _grantRole(SERVER_ROLE, server);
    }

    // region ---- Admin ----

    function initialize(
        address owlTokenAddr,
        address payable owldinalNftAddr,
        address mysteryBoxAddr
    ) external onlyRole(DEFAULT_ADMIN_ROLE) {
        owlToken = ERC20Burnable(owlTokenAddr);
        owldinalNftContract = Owldinal(owldinalNftAddr);
        mysteryBoxContract = OwldinalGenOneBox(mysteryBoxAddr);
    }

    function setServer(address server) external onlyRole(DEFAULT_ADMIN_ROLE) {
        _grantRole(SERVER_ROLE, server);
    }

    function addPrize(
        uint256 prizeAmount
    ) external onlyRole(DEFAULT_ADMIN_ROLE) {
        require(
            owlToken.transferFrom(msg.sender, address(this), prizeAmount),
            "Transfer failed"
        );
        prizePool += prizeAmount;
        emit PrizePoolIncreased(prizeAmount);
    }

    function setGlobalFruitRewardsProportion(
        uint256 proportion
    ) external onlyRole(DEFAULT_ADMIN_ROLE) {
        globalFruitRewardsProportion = proportion;
    }

    function setMoonBoost(bool isEnable) external onlyRole(DEFAULT_ADMIN_ROLE) {
        isMoonBoostEnable = isEnable;
    }

    function withdraw(
        address recipient
    ) external onlyRole(DEFAULT_ADMIN_ROLE) returns (uint256) {
        require(recipient != address(0), "wrong recipient");
        uint256 balance = owlToken.balanceOf(address(this));
        require(balance > 0, "No funds");
        bool sent = owlToken.transfer(recipient, balance);
        require(sent, "Transfer failed");

        emit PrizePoolDecreased(prizePool);
        prizePool = 0;

        return balance;
    }

    // endregion ---- Admin ----

    function getTokenInfo(
        uint256 tokenId
    ) external view returns (TokenStakingInfo memory) {
        return tokenInfoMap[tokenId];
    }

    // region ---- Player ----

    // Handles the binding of referral relationships. Given an invite code, this function
    // finds the associated wallet address and sets the caller's wallet address as the invitee
    // in the relationship. It returns the inviter's wallet address upon successful binding.
    function handleInviteCode(
        uint32 inviteCode
    ) external returns (address inviter) {
        address sender = msg.sender;

        // check if is first in game
        bool isFirstInGame = inviterToInviteCodeMap[sender] == 0;
        require(isFirstInGame, "already in game");

        // check invite code
        inviter = inviteCodeToInviterMap[inviteCode];
        require(inviter != address(0), "wrong invite code");
        inviterToInviteesMap[inviter].push(sender);
        inviteeToInviterMap[sender] = inviter;

        // generate invite code
        uint32 newInviteCode = _generateInviteCode();
        inviteCodeToInviterMap[newInviteCode] = sender;
        inviterToInviteCodeMap[sender] = newInviteCode;

        emit JoinGame(sender, newInviteCode);
        emit BindInvitation(sender, inviter);
    }

    function getInviteCode(
        address inviter
    ) external view returns (uint32 inviteCode) {
        return inviterToInviteCodeMap[inviter];
    }

    function requestMint(uint256 count) external nonReentrant {
        require(count > 0, "wrong count");

        uint256 mintPrice = MINT_PRICE * count;
        require(
            owlToken.transferFrom(msg.sender, address(this), mintPrice),
            "Transfer failed"
        );

        _handleFirstInGame(msg.sender);
        address inviter = inviteeToInviterMap[msg.sender];

        // 50% of the cost incurred during each minting operation is allocated to the reward pool.
        uint256 prizeAmount = mintPrice / 2;
        prizePool = prizePool + prizeAmount;

        uint256 burnAmount;
        if (inviter == address(0)) {
            // If the minter's address does not have an inviter, 50% of the minting cost is burned
            burnAmount = mintPrice - prizeAmount;
        } else {
            // If the minter's address has an inviter, 40% of the minting cost is burned,
            // and 10% are rewarded back to the inviter.
            uint256 rebateAmount = MINT_REBATE * count;
            burnAmount = mintPrice - prizeAmount - rebateAmount;

            Rebate storage inviterRebate = inviterRebateMap[inviter];
            inviterRebate.totalRebateEarned += rebateAmount;
            inviterRebate.rebatePendingWithdrawal += rebateAmount;
            emit RebateRewardsIncreased(
                inviter,
                msg.sender,
                count,
                rebateAmount
            );
        }

        // burn cost
        owlToken.burn(burnAmount);
        emit OwlTokenBurned(msg.sender, count, burnAmount);

        mintRequestCounter++;
        requestMintMap[mintRequestCounter] = MintRequest({
            user: msg.sender,
            count: count
        });

        emit PrizePoolIncreased(prizeAmount);
        emit RequestMint(msg.sender, mintRequestCounter, count);
    }

    // GAS=329550: tokenIdList.length=3
    function stakeOwldinalNft(uint256[] calldata tokenIdList) external {
        require(tokenIdList.length > 0, "Param is empty");
        // check a maximum of 3 Owldinal NFTs can be staked.
        require(
            stakedOwldinalsByOwner[msg.sender].length + tokenIdList.length <= 3,
            "maximum is 3"
        );

        _handleFirstInGame(msg.sender);

        // check all the token is owner's
        uint256 length = tokenIdList.length;
        for (uint256 i = 0; i < length; i++) {
            uint256 tokenId = tokenIdList[i];
            require(
                owldinalNftContract.ownerOf(tokenId) == msg.sender,
                "Not owner"
            );

            owldinalNftContract.transferFrom(
                msg.sender,
                address(this),
                tokenId
            );

            stakedOwldinalsByOwner[msg.sender].push(tokenId);
            owlInfoMap[tokenId] = OwlStakingInfo({
                tokenId: tokenId,
                owner: msg.sender,
                stakingTime: uint64(block.timestamp),
                buffedTargetIds: new uint256[](0)
            });
        }

        emit StakeOwldinalNft(msg.sender, tokenIdList);
    }

    function stakeMysteryBox(uint256[] calldata tokenIdList) external {
        require(tokenIdList.length > 0, "Param is empty");
        // check a maximum of 3 Owldinal NFTs can be staked.
        _handleFirstInGame(msg.sender);

        // check all the token is owner's
        uint256 length = tokenIdList.length;
        uint16 buffLevel;
        for (uint256 i = 0; i < length; i++) {
            uint256 tokenId = tokenIdList[i];
            require(
                mysteryBoxContract.ownerOf(tokenId) == msg.sender,
                "Not owner"
            );

            mysteryBoxContract.transferFrom(msg.sender, address(this), tokenId);
            // update owl buffer
            uint256[] storage stakingOwlIds = stakedOwldinalsByOwner[
                msg.sender
            ];
            buffLevel = uint16(stakingOwlIds.length);
            for (uint256 j = 0; j < stakingOwlIds.length; j++) {
                owlInfoMap[stakingOwlIds[j]].buffedTargetIds.push(tokenId);
                buffingOwlsByMysteryBox[tokenId].push(stakingOwlIds[j]);
            }

            BoxType boxType = mysteryBoxContract.getBoxType(tokenId);
            tokenInfoMap[tokenId] = TokenStakingInfo({
                tokenId: tokenId,
                owner: msg.sender,
                stakingTime: uint64(block.timestamp),
                boxType: boxType,
                buffLevel: buffLevel,
                reward: 0
            });

            if (boxType == BoxType.FRUIT) {
                fruitIdList.push(tokenId);
            } else if (boxType == BoxType.ELF) {
                elfIdList.push(tokenId);
            }
        }

        emit StakeMysteryBox(msg.sender, tokenIdList, buffLevel);
    }

    function unstakeOwldinalNft(uint256[] calldata tokenIdList) external {
        require(tokenIdList.length > 0, "Param is empty");
        uint256 length = tokenIdList.length;
        for (uint256 i = 0; i < length; i++) {
            uint256 tokenId = tokenIdList[i];
            require(owlInfoMap[tokenId].owner == msg.sender, "Not owner");
            require(
                owlInfoMap[tokenId].buffedTargetIds.length == 0,
                "owl still in use"
            );

            owldinalNftContract.transferFrom(
                address(this),
                msg.sender,
                tokenId
            );
            delete owlInfoMap[tokenId];
            uint256[] storage ids = stakedOwldinalsByOwner[msg.sender];
            Utils.removeValue(ids, tokenId);
        }

        emit UnstakeOwldinalNft(msg.sender, tokenIdList);
    }

    function claimAndUnstakeMysteryBox(
        uint256[] calldata tokenIdList
    ) external {
        require(tokenIdList.length > 0, "Param is empty");

        uint256 totalRewardsCanClaim = 0;
        uint256 totalRewardsForElf = 0;
        uint256 totalRewardsToBurn = 0;

        bool hasMoonBoost = false;
        if (isMoonBoostEnable) {
            uint256 whiteListLength = moonBoostWhiteList.length;
            for (uint256 i = 0; i < whiteListLength; i++) {
                if (msg.sender == moonBoostWhiteList[i]) {
                    hasMoonBoost = true;
                    break;
                }
            }

            if (!hasMoonBoost) {
                uint256 nftBalance = owldinalNftContract.balanceOf(msg.sender);
                if (nftBalance > 0) {
                    hasMoonBoost = true;
                }
            }
        }

        uint256 length = tokenIdList.length;
        for (uint256 i = 0; i < length; i++) {
            uint256 tokenId = tokenIdList[i];
            require(tokenInfoMap[tokenId].owner == msg.sender, "Not owner");

            TokenStakingInfo storage tokenInfo = tokenInfoMap[tokenId];

            // calculate rewards
            if (tokenInfo.reward > 0) {
                if (tokenInfo.boxType == BoxType.FRUIT) {
                    uint256 rewardProportion;
                    if (tokenInfo.buffLevel >= 3) {
                        rewardProportion = 85;
                    } else {
                        rewardProportion = 75;
                    }
                    if (hasMoonBoost) {
                        rewardProportion += 7;
                    }

                    uint256 rewardsCanClaim = (tokenInfo.reward *
                        rewardProportion) / 100;
                    uint256 rewardsForElf = tokenInfo.reward - rewardsCanClaim;

                    totalRewardsCanClaim += rewardsCanClaim;
                    totalRewardsForElf += rewardsForElf;

                    emit UnstakeMysteryBox(
                        msg.sender,
                        tokenId,
                        tokenInfo.boxType,
                        rewardsCanClaim
                    );
                } else {
                    // tokenInfo.boxType == BoxType.ELF
                    uint256 rewardProportion;
                    if (tokenInfo.buffLevel >= 2) {
                        rewardProportion = 90;
                    } else {
                        rewardProportion = 85;
                    }
                    if (hasMoonBoost) {
                        rewardProportion += 7;
                    }
                    uint256 rewardsCanClaim = (tokenInfo.reward *
                        rewardProportion) / 100;
                    uint256 rewardsToBurn = tokenInfo.reward - rewardsCanClaim;
                    totalRewardsCanClaim += rewardsCanClaim;
                    totalRewardsToBurn += rewardsToBurn;

                    emit UnstakeMysteryBox(
                        msg.sender,
                        tokenId,
                        tokenInfo.boxType,
                        rewardsCanClaim
                    );
                }
            } else {
                emit UnstakeMysteryBox(
                    msg.sender,
                    tokenId,
                    tokenInfo.boxType,
                    0
                );
            }

            // unstake nfts
            mysteryBoxContract.transferFrom(
                address(this),
                tokenInfo.owner,
                tokenId
            );

            // remove from owl's buff list
            if (buffingOwlsByMysteryBox[tokenId].length > 0) {
                for (
                    uint256 j = 0;
                    j < buffingOwlsByMysteryBox[tokenId].length;
                    j++
                ) {
                    uint256 owlId = buffingOwlsByMysteryBox[tokenId][j];
                    Utils.removeValue(
                        owlInfoMap[owlId].buffedTargetIds,
                        tokenId
                    );
                }
                delete buffingOwlsByMysteryBox[tokenId];
            }

            // remove from token info
            if (tokenInfo.boxType == BoxType.FRUIT) {
                Utils.removeValue(fruitIdList, tokenId);
            } else {
                Utils.removeValue(elfIdList, tokenId);
            }
            delete tokenInfoMap[tokenId];
        }

        // handle rewards
        if (totalRewardsCanClaim > 0) {
            owlToken.transfer(msg.sender, totalRewardsCanClaim);
        }
        if (totalRewardsToBurn > 0) {
            owlToken.burn(totalRewardsToBurn);
            emit OwlTokenBurned(msg.sender, totalRewardsToBurn, 0);
        }
        if (totalRewardsForElf > 0 && elfIdList.length > 0) {
            uint256 eachElfRewards = totalRewardsForElf / elfIdList.length;
            for (uint256 i = 0; i < elfIdList.length; i++) {
                tokenInfoMap[elfIdList[i]].reward += eachElfRewards;
            }
            emit ElfRewardUpdated(eachElfRewards, elfIdList.length);
        }
    }

    function claimInviterReward() external {
        Rebate storage rebate = inviterRebateMap[msg.sender];
        require(rebate.rebatePendingWithdrawal > 0, "no reward");
        require(rebate.unlockedRebateToClaim > 0, "need more unlock");

        // It should be noted that the unlock provided by each mint is calculated as a percentage (for example,
        // the maximum reward is that one mint will unlock 850% of the rebate), so be careful to divide by 100
        // when using it.
        uint256 withdrawAmount;
        if (rebate.rebatePendingWithdrawal > rebate.unlockedRebateToClaim) {
            withdrawAmount = rebate.unlockedRebateToClaim;
            rebate.unlockedRebateToClaim = 0;
            rebate.rebatePendingWithdrawal -= withdrawAmount;
        } else {
            withdrawAmount = rebate.rebatePendingWithdrawal;
            rebate.rebatePendingWithdrawal = 0;
            rebate.unlockedRebateToClaim -= withdrawAmount;
        }

        owlToken.transfer(msg.sender, withdrawAmount);
        emit RebateClaimed(msg.sender, withdrawAmount);
    }

    // endregion ---- Player ----

    // region ---- Server ----

    // After staking, Owl rewards are distributed every 4 hours.
    function updateAllFruitRewards() external onlyRole(SERVER_ROLE) {
        require(fruitIdList.length > 0, "No fruit need rewards"); // TODO: use return instead

        // calculate how many fruit can get rewards. Loop twice to reduce the gas cost
        // brought by the length of rewardFruitIdList.
        uint256 rewardFruitCount = 0;
        uint256 length = fruitIdList.length;
        for (uint256 i = 0; i < length; i++) {
            uint256 fruitId = fruitIdList[i];
            TokenStakingInfo storage fruit = tokenInfoMap[fruitId];
            if (
                (block.timestamp - fruit.stakingTime) >= FRUIT_REWARD_INTERVAL
            ) {
                rewardFruitCount++;
            }
        }

        require(rewardFruitCount > 0, "No fruit need rewards"); // TODO: use return instead

        uint256[] memory rewardFruitIdList = new uint256[](rewardFruitCount);
        uint256 index = 0;
        for (uint256 i = 0; i < length; i++) {
            uint256 fruitId = fruitIdList[i];
            TokenStakingInfo storage fruit = tokenInfoMap[fruitId];
            if (
                (block.timestamp - fruit.stakingTime) >= FRUIT_REWARD_INTERVAL
            ) {
                rewardFruitIdList[index] = fruitId;
                index++;
            }
        }

        // fruitRewardsProportion should divied by 100_000_000
        uint256 fruitRewardsProportion = _calculateFruitRewardsProportion(
            rewardFruitCount
        );
        uint256 totalRewards = (prizePool * fruitRewardsProportion) /
            100_000_000;
        uint256 eachFruitRewards = totalRewards / rewardFruitCount;
        require(eachFruitRewards > 0, "rewards should not be zero");

        for (uint256 i = 0; i < rewardFruitCount; i++) {
            uint256 fruitId = rewardFruitIdList[i];
            TokenStakingInfo storage fruit = tokenInfoMap[fruitId];
            fruit.reward += eachFruitRewards;
            fruit.stakingTime = uint64(block.timestamp);
        }
        emit FruitRewardUpdated(
            eachFruitRewards,
            rewardFruitCount,
            fruitIdList.length,
            elfIdList.length
        );

        prizePool -= totalRewards;
        emit PrizePoolDecreased(totalRewards);
    }

    function mintMysteryBox(
        uint256 mintRequestId
    ) external nonReentrant onlyRole(SERVER_ROLE) {
        MintRequest storage request = requestMintMap[mintRequestId];
        require(request.user != address(0), "no request");

        // get owldinal buff level
        uint256[] storage stakingOwlIds = stakedOwldinalsByOwner[request.user];
        uint256 buffLevel = stakingOwlIds.length;

        // burn cost and do mint, mystery box will open directly (maybe burned)
        uint256[] memory tokenIdList = mysteryBoxContract.mintAndOpenBoxes(
            request.user,
            request.count,
            buffLevel > 0
        );

        // handle unlockable rebate as inviter
        Rebate storage playerRebate = inviterRebateMap[request.user];
        uint256 addUnlockedAmount = _calculateUnlockRebateAmountByMint(
            request.count,
            playerRebate.mintedBoxCount
        );
        playerRebate.mintedBoxCount += request.count;
        playerRebate.unlockedRebateToClaim += addUnlockedAmount;
        emit UnlockableRebateIncreased(
            request.user,
            request.count,
            addUnlockedAmount
        );
        emit MintMysteryBox(
            request.user,
            mintRequestId,
            request.count,
            tokenIdList
        );

        delete requestMintMap[mintRequestId];
    }

    // endregion ---- Server ----

    // if player is first in game, then generate a invite code for player.
    function _handleFirstInGame(address sender) private {
        bool isFirstInGame = inviterToInviteCodeMap[sender] == 0;
        if (isFirstInGame) {
            uint32 newInviteCode = _generateInviteCode();
            inviteCodeToInviterMap[newInviteCode] = sender;
            inviterToInviteCodeMap[sender] = newInviteCode;
            emit JoinGame(sender, newInviteCode);
        }
    }

    function _handleInviter(
        address sender,
        uint32 inviteCode
    ) private returns (address) {
        // check inviter
        address inviter = inviteeToInviterMap[sender];
        if (inviter == address(0)) {
            if (inviteCode > 0) {
                inviter = inviteCodeToInviterMap[inviteCode];
                require(
                    inviter != address(0),
                    "failed to find this invite code"
                );
                inviterToInviteesMap[inviter].push(sender);
                inviteeToInviterMap[sender] = inviter;
            }
        }
        return inviter;
    }

    function _generateInviteCode() private view returns (uint32) {
        uint32 newCode;
        bool isUnique;
        uint256 rand = block.number;
        do {
            newCode = 0;
            rand = Utils.generateRandomNumber();
            for (uint i = 0; i < 5; i++) {
                uint32 charValue = uint32(rand % 26);
                newCode |= (charValue << uint32(i * 5));
                rand = rand / 26;
            }

            // check duplicated
            isUnique = inviteCodeToInviterMap[newCode] == address(0);
        } while (!isUnique);

        return newCode;
    }

    // Calculate the bonus proportion allocated to the current Fruit.
    // When globalFruitRewardsProportion is not zero, it will be used preferentially.
    function _calculateFruitRewardsProportion(
        uint256 rewardFruitCount
    ) private view returns (uint256) {
        if (globalFruitRewardsProportion > 0) {
            return globalFruitRewardsProportion;
        }

        uint256 fruitRewardsProportion;
        if (rewardFruitCount <= 1) {
            fruitRewardsProportion = 100; // 0.0001000% * 100_000_000
        } else if (rewardFruitCount <= 50) {
            fruitRewardsProportion = 5000; // 0.0050000% * 100_000_000
        } else if (rewardFruitCount <= 100) {
            fruitRewardsProportion = 10000; // 0.0100000% * 100_000_000
        } else if (rewardFruitCount <= 200) {
            fruitRewardsProportion = 21000; // 0.0210000% * 100_000_000
        } else if (rewardFruitCount <= 400) {
            fruitRewardsProportion = 43040; // 0.0430400% * 100_000_000
        } else if (rewardFruitCount <= 800) {
            fruitRewardsProportion = 88230; // 0.0882300% * 100_000_000
        } else if (rewardFruitCount <= 1200) {
            fruitRewardsProportion = 136760; // 0.1367600% * 100_000_000
        } else if (rewardFruitCount <= 1800) {
            fruitRewardsProportion = 211980; // 0.2119800% * 100_000_000
        } else if (rewardFruitCount <= 2300) {
            fruitRewardsProportion = 276750; // 0.2767500% * 100_000_000
        } else if (rewardFruitCount <= 2800) {
            fruitRewardsProportion = 342930; // 0.3429300% * 100_000_000
        } else if (rewardFruitCount <= 3300) {
            fruitRewardsProportion = 416410; // 0.4164100% * 100_000_000
        } else if (rewardFruitCount <= 4000) {
            fruitRewardsProportion = 522400; // 0.5224000% * 100_000_000
        } else if (rewardFruitCount <= 5000) {
            fruitRewardsProportion = 679130; // 0.6791300% * 100_000_000
        } else if (rewardFruitCount <= 6000) {
            fruitRewardsProportion = 842120; // 0.8421200% * 100_000_000
        } else if (rewardFruitCount <= 7500) {
            fruitRewardsProportion = 1190480; // 1.1904800% * 100_000_000
        } else {
            fruitRewardsProportion = 1349210; // 1.3492100% * 100_000_000
        }

        return fruitRewardsProportion;
    }

    // Each mint of a mystery box unlocks a specific rebate reward k based on the following conditions:
    // - When n <= 10, the k is 30,000 for each n.
    // - When 10 < n <= 50, the k is 55,000 for each n.
    // - When 50 < n <= 100, the k is 70,000 for each n.
    // - When n > 100, the k is 85,000 for each n.
    function _calculateUnlockRebateAmountByMint(
        uint256 mintBoxCount,
        uint256 prevMintedBoxCount
    ) internal pure returns (uint256 amount) {
        uint256 totalMintCount = mintBoxCount + prevMintedBoxCount;
        if (totalMintCount <= 10) {
            amount = 30000 * 10 ** 18 * mintBoxCount;
        } else if (totalMintCount <= 50) {
            amount = 55000 * 10 ** 18 * mintBoxCount;
        } else if (totalMintCount <= 100) {
            amount = 70000 * 10 ** 18 * mintBoxCount;
        } else {
            amount = 85000 * 10 ** 18 * mintBoxCount;
        }

        return amount;
    }
}
