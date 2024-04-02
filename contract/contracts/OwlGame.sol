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

    // 用户B Mint 盲盒消耗的Owl，10%实时返给用户A。（每mint一个盲盒，返利10000 Owl，定义为每份奖励）
    // 返利的Owl，积累在用户A的账户中，但是是locked的
    // 用户A需要mint盲盒，才能解锁返利的奖励，每mint 1个盲盒，可以解锁K/10份奖励，mint n个盲盒，可以解锁 $n*k/10$ 份奖励
    // - n ≤ 10 => K=30
    // - 10 < n ≤ 50 => K=55
    // - 50 < n ≤ 100 => K=70
    // - 100 < n => K=85
    struct Rebate {
        uint256 totalRebateEarned;
        uint256 rebatePendingWithdrawal;
        uint256 unlockedRebateToClaim;
        uint256 mintedBoxCount;
    }

    // event MintBox(address indexed user, uint256 tokenId);
    // event OpenBox(address indexed user, uint256 tokenId, BoxType boxType);
    event JoinGame(address indexed user);
    event BindInvitation(address indexed invitee, address inviter);
    event MintBox(address indexed user, uint256 count, uint256[] tokenIdList);
    event StakeOwldinalNft(address indexed user, uint256[] tokenId);
    event StakeMysteryBox(address indexed user, uint256[] tokenId);
    event UnstakeOwldinalNft(address indexed user, uint256[] tokenId);
    // This event cannot accept batch parameters, otherwise it would be impossible to
    // distinguish the rewards for each specific token.
    event UnstakeMysteryBox(
        address indexed user,
        uint256 tokenId,
        BoxType boxType,
        uint256 rewards
    );

    bytes32 public constant SERVER_ROLE = keccak256("SERVER_ROLE");

    ERC20Burnable public owlToken;
    Owldinal public owldinalNftContract;
    OwldinalGenOneBox public mysteryBoxContract;

    // The price that needs to be spent for each mint.
    uint256 public constant MINT_PRICE = 100000;
    uint256 public constant MINT_REBATE = 10000;
    uint256 public constant FRUIT_REWARD_INTERVAL = 3600;

    // The reward proportion for Fruit, scaled by a factor of 1,000,000.
    // This value starts at 0 and is calculated by the _calculateFruitRewardsProportion
    // method based on the current staked amount of fruit. When the value is 0, it indicates
    // that the rewards proportion has not been set yet.
    uint256 private globalFruitRewardsProportion = 0;

    // Owl Token Id -> OwlStakingInfo
    mapping(uint256 => OwlStakingInfo) owlInfoMap;

    // owner address -> owl token id list
    mapping(address => uint256[]) stakedOwldinalsByOwner;
    mapping(uint256 => TokenStakingInfo) tokenInfoMap;

    uint256[] fruitIdList;
    uint256[] elfIdList;

    // mystery box token id -> buffing owldinals id
    mapping(uint256 => uint256[]) buffingOwlsByMysteryBox;

    // inviter address -> invitees
    mapping(address => address[]) inviterToInviteesMap;
    // invitees address -> inviter
    mapping(address => address) inviteeToInviterMap;

    // invite code -> inviter (find invter from code)
    mapping(uint32 => address) inviteCodeToInviterMap;
    // inviter -> invite code (check if first join game)
    mapping(address => uint32) inviterToInviteCodeMap;

    // inviter -> unclaimed rebate
    mapping(address => Rebate) inviterRebateMap;

    // prize pool
    uint256 public prizePool;

    bool public isMoonBoostEnable;
    address[] private moonBoostWhiteList = [address(0xAABB)];

    constructor() {
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
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

    function addPrize(
        uint256 prizeAmount
    ) external onlyRole(DEFAULT_ADMIN_ROLE) {
        require(
            owlToken.transferFrom(msg.sender, address(this), prizeAmount),
            "Transfer failed"
        );
        prizePool = prizePool + prizeAmount;
    }

    function setGlobalFruitRewardsProportion(
        uint256 proportion
    ) external onlyRole(DEFAULT_ADMIN_ROLE) {
        globalFruitRewardsProportion = proportion;
    }

    function setMoonBoost(bool isEnable) external onlyRole(DEFAULT_ADMIN_ROLE) {
        isMoonBoostEnable = isEnable;
    }

    // endregion ---- Admin ----

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
        require(
            isFirstInGame,
            "Players who have already participated in the game cannot be invited."
        );

        // check invite code
        inviter = inviteCodeToInviterMap[inviteCode];
        require(inviter != address(0), "failed to find this invite code");
        inviterToInviteesMap[inviter].push(sender);
        inviteeToInviterMap[sender] = inviter;

        // generate invite code
        uint32 newInviteCode = _generateInviteCode();
        inviteCodeToInviterMap[newInviteCode] = sender;
        inviterToInviteCodeMap[sender] = newInviteCode;

        emit JoinGame(sender);
        emit BindInvitation(sender, inviter);
    }

    function mintMysteryBox(
        uint256 count
    ) external nonReentrant returns (uint256[] memory tokenIdList) {
        require(count > 0, "Count should larger than 0");

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
        }

        // get owldinal buff level
        uint256[] storage stakingOwlIds = stakedOwldinalsByOwner[msg.sender];
        uint256 buffLevel = stakingOwlIds.length;

        // burn cost and do mint, mystery box will open directly (maybe burned)
        owlToken.burn(burnAmount);
        tokenIdList = mysteryBoxContract.mintAndOpenBoxes(
            msg.sender,
            count,
            buffLevel > 0
        );
        // tokenIdList = new uint256[](count);
        // for (uint256 i = 0; i < count; i++) {
        //     tokenIdList[i] = mysteryBoxContract.mintBox(msg.sender);
        // }

        // handle unlockable rebate as inviter
        Rebate storage playerRebate = inviterRebateMap[msg.sender];
        uint256 addUnlockedAmount = _calculateUnlockRebateAmountByMint(
            count,
            playerRebate.mintedBoxCount
        );
        playerRebate.mintedBoxCount += count;
        playerRebate.unlockedRebateToClaim += addUnlockedAmount;

        return tokenIdList;
    }

    function stakeOwldinalNft(uint256[] calldata tokenIdList) external {
        require(tokenIdList.length > 0, "Param is empty");
        // check a maximum of 3 Owldinal NFTs can be staked.
        require(
            stakedOwldinalsByOwner[msg.sender].length + tokenIdList.length <= 3,
            "Owldinal can stake a maximum of three."
        );

        _handleFirstInGame(msg.sender);

        // check all the token is owner's
        for (uint256 i = 0; i < tokenIdList.length; i++) {
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
        for (uint256 i = 0; i < tokenIdList.length; i++) {
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
            uint16 buffLevel = uint16(stakingOwlIds.length);
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

        emit StakeMysteryBox(msg.sender, tokenIdList);
    }

    function unstakeOwldinalNft(uint256[] calldata tokenIdList) external {
        require(tokenIdList.length > 0, "Param is empty");
        for (uint256 i = 0; i < tokenIdList.length; i++) {
            uint256 tokenId = tokenIdList[i];
            require(owlInfoMap[tokenId].owner == msg.sender, "Not owner");
            require(
                owlInfoMap[tokenId].buffedTargetIds.length == 0,
                "this owl token still in use"
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
            for (uint256 i = 0; i < moonBoostWhiteList.length; i++) {
                if (msg.sender == moonBoostWhiteList[i]) {
                    hasMoonBoost = true;
                    break;
                }
            }
        }

        for (uint256 i = 0; i < tokenIdList.length; i++) {
            uint256 tokenId = tokenIdList[i];
            require(tokenInfoMap[tokenId].owner == msg.sender, "Not owner");

            TokenStakingInfo storage tokenInfo = tokenInfoMap[tokenId];

            // calculate rewards
            if (tokenInfo.reward > 0) {
                if (tokenInfo.boxType == BoxType.FRUIT) {
                    uint256 rewardProportion;
                    if (hasMoonBoost) {
                        rewardProportion = 88;
                    } else if (tokenInfo.buffLevel >= 3) {
                        rewardProportion = 85;
                    } else {
                        rewardProportion = 75;
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
                    if (hasMoonBoost) {
                        rewardProportion = 93;
                    } else if (tokenInfo.buffLevel >= 2) {
                        rewardProportion = 90;
                    } else {
                        rewardProportion = 85;
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
        }
        if (totalRewardsForElf > 0) {
            uint256 eachElfRewards = totalRewardsForElf / elfIdList.length;
            for (uint256 i = 0; i < elfIdList.length; i++) {
                tokenInfoMap[elfIdList[i]].reward += eachElfRewards;
            }
        }
    }

    function claimInviterReward() external {
        Rebate storage rebate = inviterRebateMap[msg.sender];
        require(rebate.rebatePendingWithdrawal > 0, "no reward can be claimed");
        require(
            rebate.unlockedRebateToClaim > 0,
            "need more minted box for claim inviter reward"
        );

        // It should be noted that the unlock provided by each mint is calculated as a percentage (for example,
        // the maximum reward is that one mint will unlock 850% of the rebate), so be careful to divide by 100
        // when using it.
        uint256 withdrawAmount;
        if (rebate.rebatePendingWithdrawal > rebate.unlockedRebateToClaim) {
            withdrawAmount = rebate.unlockedRebateToClaim;
            owlToken.transfer(msg.sender, withdrawAmount);
            rebate.unlockedRebateToClaim = 0;
            rebate.rebatePendingWithdrawal =
                rebate.rebatePendingWithdrawal -
                withdrawAmount;
        } else {
            withdrawAmount = rebate.rebatePendingWithdrawal;
            owlToken.transfer(msg.sender, withdrawAmount);
            rebate.rebatePendingWithdrawal = 0;
            rebate.unlockedRebateToClaim =
                rebate.unlockedRebateToClaim -
                rebate.rebatePendingWithdrawal;
        }
    }

    // endregion ---- Player ----

    // region ---- Server ----

    // After staking, Owl rewards are distributed every 4 hours.
    function updateAllFruitRewards() external onlyRole(SERVER_ROLE) {
        require(fruitIdList.length > 0, "No fruit need update rewards");

        // calculate how many fruit can get rewards. Loop twice to reduce the gas cost
        // brought by the length of rewardFruitIdList.
        uint256 rewardFruitCount = 0;
        for (uint256 i = 0; i < fruitIdList.length; i++) {
            uint256 fruitId = fruitIdList[i];
            TokenStakingInfo storage fruit = tokenInfoMap[fruitId];
            if ((block.timestamp - fruit.stakingTime) > FRUIT_REWARD_INTERVAL) {
                rewardFruitCount++;
            }
        }
        uint256[] memory rewardFruitIdList = new uint256[](rewardFruitCount);
        uint256 index = 0;
        for (uint256 i = 0; i < fruitIdList.length; i++) {
            uint256 fruitId = fruitIdList[i];
            TokenStakingInfo storage fruit = tokenInfoMap[fruitId];
            if ((block.timestamp - fruit.stakingTime) > FRUIT_REWARD_INTERVAL) {
                rewardFruitIdList[index] = fruitId;
                index++;
            }
        }

        // fruitRewardsProportion should divied by 1_000_000
        uint256 fruitRewardsProportion = _calculateFruitRewardsProportion(
            rewardFruitCount
        );
        uint256 totalRewards = (prizePool * fruitRewardsProportion) / 1_000_000;
        uint256 eachFruitRewards = totalRewards / rewardFruitCount;
        require(eachFruitRewards > 0, "rewards should not be zero");

        for (uint256 i = 0; i < rewardFruitIdList.length; i++) {
            uint256 fruitId = rewardFruitIdList[i];
            TokenStakingInfo storage fruit = tokenInfoMap[fruitId];
            fruit.reward += eachFruitRewards;
            fruit.stakingTime = uint64(block.timestamp);
        }

        prizePool -= totalRewards;
    }

    // endregion ---- Server ----

    // if player is first in game, then generate a invite code for player.
    function _handleFirstInGame(address sender) private {
        bool isFirstInGame = inviterToInviteCodeMap[sender] == 0;
        if (isFirstInGame) {
            uint32 newInviteCode = _generateInviteCode();
            inviteCodeToInviterMap[newInviteCode] = sender;
            inviterToInviteCodeMap[sender] = newInviteCode;
        }

        emit JoinGame(sender);
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
        do {
            newCode = 0;
            uint256 rand = Utils.generateRandomNumber();
            for (uint i = 0; i < 5; i++) {
                uint8 charValue = uint8(rand % 26);
                newCode |= (charValue << uint8(i * 5));
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
            fruitRewardsProportion = 100; // 0.0001000% * 1_000_000
        } else if (rewardFruitCount <= 50) {
            fruitRewardsProportion = 5000; // 0.0050000% * 1_000_000
        } else if (rewardFruitCount <= 100) {
            fruitRewardsProportion = 10000; // 0.0100000% * 1_000_000
        } else if (rewardFruitCount <= 200) {
            fruitRewardsProportion = 21000; // 0.0210000% * 1_000_000
        } else if (rewardFruitCount <= 400) {
            fruitRewardsProportion = 43040; // 0.0430400% * 1_000_000
        } else if (rewardFruitCount <= 800) {
            fruitRewardsProportion = 88230; // 0.0882300% * 1_000_000
        } else if (rewardFruitCount <= 1200) {
            fruitRewardsProportion = 136760; // 0.1367600% * 1_000_000
        } else if (rewardFruitCount <= 1800) {
            fruitRewardsProportion = 211980; // 0.2119800% * 1_000_000
        } else if (rewardFruitCount <= 2300) {
            fruitRewardsProportion = 276750; // 0.2767500% * 1_000_000
        } else if (rewardFruitCount <= 2800) {
            fruitRewardsProportion = 342930; // 0.3429300% * 1_000_000
        } else if (rewardFruitCount <= 3300) {
            fruitRewardsProportion = 416410; // 0.4164100% * 1_000_000
        } else if (rewardFruitCount <= 4000) {
            fruitRewardsProportion = 522400; // 0.5224000% * 1_000_000
        } else if (rewardFruitCount <= 5000) {
            fruitRewardsProportion = 679130; // 0.6791300% * 1_000_000
        } else if (rewardFruitCount <= 6000) {
            fruitRewardsProportion = 842120; // 0.8421200% * 1_000_000
        } else if (rewardFruitCount <= 7500) {
            fruitRewardsProportion = 1190480; // 1.1904800% * 1_000_000
        } else {
            fruitRewardsProportion = 1349210; // 1.3492100% * 1_000_000
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
            amount = 30000 * mintBoxCount;
        } else if (totalMintCount <= 50) {
            amount = 55000 * mintBoxCount;
        } else if (totalMintCount <= 100) {
            amount = 70000 * mintBoxCount;
        } else {
            amount = 85000 * mintBoxCount;
        }

        return amount;
    }

    function _encodeInviteCode(
        string memory inviteCode
    ) internal pure returns (uint32 encoded) {
        require(bytes(inviteCode).length == 5, "Invalid invite code length");
        for (uint i = 0; i < 5; i++) {
            bytes1 char = bytes(inviteCode)[i];
            uint8 charValue = uint8(char) - 0x41;
            encoded |= (charValue << uint8(i * 5));
        }
    }

    function _decodeInviteCode(
        uint32 encoded
    ) internal pure returns (string memory) {
        bytes memory inviteCode = new bytes(5);
        for (uint i = 0; i < 5; i++) {
            uint8 charValue = uint8((encoded >> (i * 5)) & 0x1F);
            inviteCode[i] = bytes1(charValue + 0x41);
        }

        return string(inviteCode);
    }
}
