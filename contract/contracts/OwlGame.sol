// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/access/AccessControl.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Burnable.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721Burnable.sol";

import "./Owldinal.sol";
import "./MysteryBoxGen1.sol";
import "./Utils.sol";

contract OwlGame is AccessControl {
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

    struct Invite {
        address inviter;
        address invitee;
        bool isSet;
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
        // It should be noted that the unlock provided by each mint is calculated as a percentage (for example,
        // the maximum reward is that one mint will unlock 850% of the rebate), so be careful to divide by 100
        // when using it.
        uint32 unlockableRewardCountPercent;
        uint16 mintedBoxCount;
    }

    ERC20Burnable public owlToken;
    Owldinal public boxGen0Contract;
    MysteryBoxGen1 public boxGen1Contract;
    uint256 public constant MINT_PRICE = 100000;
    uint32 public constant MINT_REBATE = 10000;
    uint256 public constant FRUIT_REWARD_INTERVAL = 3600;

    // Owl Token Id -> OwlStakingInfo
    mapping(uint256 => OwlStakingInfo) owlInfoMap;

    // owner address -> owl token id list
    mapping(address => uint256[]) owlStakingMap;

    mapping(uint256 => TokenStakingInfo) tokenInfoMap;
    uint256[] fruitIdList; // for update fruit
    uint256[] elfIdList;
    mapping(uint256 => uint256[]) tokenToBuffOwlMap;

    // inviter address -> invitees
    mapping(address => address[]) inviterToInviteesMap;
    mapping(address => address) inviteeToInviterMap;

    // invite code -> inviter (find invter from code)
    mapping(uint32 => address) inviteCodeToInviterMap;
    // inviter -> invite code (check if first join game)
    mapping(address => uint32) inviterToInviteCodeMap;

    // inviter -> unclaimed rebate
    mapping(address => Rebate) inviterRebateMap;

    // prize pool
    uint256 public prizePool;

    constructor() {
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
    }

    function addPrize(uint256 prizeAmount) external {
        require(
            owlToken.transferFrom(msg.sender, address(this), prizeAmount),
            "Transfer failed"
        );
        prizePool = prizePool + prizeAmount;
    }

    function mintGen1Box(uint32 inviteCode) external {
        require(
            owlToken.transferFrom(msg.sender, address(this), MINT_PRICE),
            "Transfer failed"
        );

        _handleFirstInGame(msg.sender);
        address inviter = _handleInviter(msg.sender, inviteCode);

        uint256 prizeAmount = MINT_PRICE / 2;
        prizePool = prizePool + prizeAmount;

        uint256 burnAmount = MINT_PRICE - prizeAmount;
        if (inviter != address(0)) {
            burnAmount = burnAmount - MINT_REBATE;

            uint256 rebateAmount = MINT_REBATE;
            Rebate storage inviterRebate = inviterRebateMap[inviter];
            inviterRebate.totalRebateEarned =
                inviterRebate.totalRebateEarned +
                rebateAmount;
            inviterRebate.rebatePendingWithdrawal =
                inviterRebate.rebatePendingWithdrawal +
                rebateAmount;
        }

        owlToken.burn(burnAmount);
        boxGen1Contract.mintBox(msg.sender);

        // add unlocakableRewardCount;
        Rebate storage playerRebate = inviterRebateMap[msg.sender];
        playerRebate.mintedBoxCount = playerRebate.mintedBoxCount + 1;
        uint32 addUnlockCount = 0;
        if (playerRebate.mintedBoxCount <= 10) {
            addUnlockCount = 300;
        } else if (playerRebate.mintedBoxCount <= 50) {
            addUnlockCount = 550;
        } else if (playerRebate.mintedBoxCount <= 100) {
            addUnlockCount = 700;
        } else {
            addUnlockCount = 850;
        }
        playerRebate.unlockableRewardCountPercent =
            playerRebate.unlockableRewardCountPercent +
            addUnlockCount;
    }

    // 1. Stake owl token or boxGen1 token. If gen1 is not unboxed, it will be unboxed directly.
    // 2. If player are participating in the game for the first time, player can input the "invite code"
    // field to establish an invitation relationship.
    function stakeBox(
        uint256 tokenId,
        bool isGen0,
        uint32 inviteCode
    ) external {
        if (isGen0) {
            require(
                boxGen0Contract.ownerOf(tokenId) == msg.sender,
                "Not owner"
            );

            require(
                owlStakingMap[msg.sender].length < 3,
                "Owldinal can stake a maximum of three."
            );
        } else {
            require(
                boxGen1Contract.ownerOf(tokenId) == msg.sender,
                "Not owner"
            );
        }
        _handleFirstInGame(msg.sender);
        _handleInviter(msg.sender, inviteCode);

        // start stake
        if (isGen0) {
            owlToken.transferFrom(msg.sender, address(this), tokenId);
            owlStakingMap[msg.sender].push(tokenId);
            owlInfoMap[tokenId] = OwlStakingInfo({
                tokenId: tokenId,
                owner: msg.sender,
                stakingTime: uint64(block.timestamp),
                buffedTargetIds: new uint256[](0)
            });
        } else {
            BoxType boxType = boxGen1Contract.getBoxType(tokenId);
            uint256[] storage stakingOwlIds = owlStakingMap[msg.sender];
            uint16 buffLevel = uint16(stakingOwlIds.length);
            if (boxType == BoxType.UNOPENED) {
                boxType = boxGen1Contract.openBox(tokenId, buffLevel > 0);
                require(
                    boxType != BoxType.UNOPENED,
                    "You mystery box has been destoryed"
                );
            }

            // update owl buff target when staking elf/fruit
            for (uint256 i = 0; i < stakingOwlIds.length; i++) {
                owlInfoMap[stakingOwlIds[i]].buffedTargetIds.push(tokenId);
                tokenToBuffOwlMap[tokenId].push(stakingOwlIds[i]);
            }

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
    }

    function unstakeOwl(uint256 tokenId) external {
        require(owlInfoMap[tokenId].owner == msg.sender, "Not owner");
        require(
            owlInfoMap[tokenId].buffedTargetIds.length == 0,
            "this owl token still in use"
        );

        boxGen0Contract.transferFrom(address(this), msg.sender, tokenId);
        delete owlInfoMap[tokenId];
        uint256[] storage ids = owlStakingMap[msg.sender];
        Utils.removeValue(ids, tokenId);
    }

    function unstakeToken(uint256 tokenId) external {
        require(tokenInfoMap[tokenId].owner == msg.sender, "Not owner");
        TokenStakingInfo storage tokenInfo = tokenInfoMap[tokenId];

        if (tokenInfo.reward > 0) {
            if (tokenInfo.boxType == BoxType.FRUIT) {
                uint256 rewardsCanClaim = (tokenInfo.reward *
                    (tokenInfo.buffLevel >= 3 ? 85 : 75)) / 100;
                uint256 rewardsForElf = tokenInfo.reward - rewardsCanClaim;

                owlToken.transfer(tokenInfo.owner, rewardsCanClaim);

                uint256 eachElfRewards = rewardsForElf / elfIdList.length;
                for (uint256 i = 0; i < elfIdList.length; i++) {
                    tokenInfoMap[elfIdList[i]].reward =
                        tokenInfoMap[elfIdList[i]].reward +
                        eachElfRewards;
                }
            } else {
                uint256 rewardsCanClaim = (tokenInfo.reward *
                    (tokenInfo.buffLevel >= 2 ? 90 : 85)) / 100;
                uint256 rewardsToBurn = tokenInfo.reward - rewardsCanClaim;

                owlToken.transfer(tokenInfo.owner, rewardsCanClaim);
                owlToken.burn(rewardsToBurn);
            }
        }

        boxGen1Contract.transferFrom(address(this), tokenInfo.owner, tokenId);

        // remove from owl's buff list
        if (tokenToBuffOwlMap[tokenId].length > 0) {
            for (uint256 i = 0; i < tokenToBuffOwlMap[tokenId].length; i++) {
                uint256 owlId = tokenToBuffOwlMap[tokenId][i];
                Utils.removeValue(owlInfoMap[owlId].buffedTargetIds, tokenId);
            }
            delete tokenToBuffOwlMap[tokenId];
        }

        // remove from token info
        if (tokenInfo.boxType == BoxType.FRUIT) {
            Utils.removeValue(fruitIdList, tokenId);
        } else {
            Utils.removeValue(elfIdList, tokenId);
        }
        delete tokenInfoMap[tokenId];
    }

    // After staking, Owl rewards are distributed every 4 hours.
    function updateAllFruitRewards() external {
        // TODO: check sender is server.

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
        uint256 totalRewards = (prizePool / 1_000_000) * fruitRewardsProportion;
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

    function claimInviterReward() external {
        Rebate storage rebate = inviterRebateMap[msg.sender];
        require(rebate.rebatePendingWithdrawal > 0, "no reward can be claimed");
        require(
            rebate.unlockableRewardCountPercent > 0,
            "need more minted box for claim inviter reward"
        );

        // It should be noted that the unlock provided by each mint is calculated as a percentage (for example,
        // the maximum reward is that one mint will unlock 850% of the rebate), so be careful to divide by 100
        // when using it.
        uint32 pendingRebatePercent = uint32(
            (rebate.rebatePendingWithdrawal / MINT_REBATE) * 100
        );
        if (pendingRebatePercent > rebate.unlockableRewardCountPercent) {
            uint256 rebateCanClaim = rebate.unlockableRewardCountPercent *
                (MINT_REBATE / 100);
            owlToken.transfer(msg.sender, rebateCanClaim);
            rebate.rebatePendingWithdrawal =
                rebate.rebatePendingWithdrawal -
                rebateCanClaim;
            rebate.unlockableRewardCountPercent = 0;
        } else {
            owlToken.transfer(msg.sender, rebate.rebatePendingWithdrawal);
            rebate.rebatePendingWithdrawal = 0;
            rebate.unlockableRewardCountPercent =
                rebate.unlockableRewardCountPercent -
                pendingRebatePercent;
        }
    }

    // if player is first in game, then generate a invite code for player.
    function _handleFirstInGame(address sender) private {
        bool isFirstInGame = inviterToInviteCodeMap[sender] == 0;
        if (isFirstInGame) {
            uint32 newInviteCode = _generateInviteCode();
            inviteCodeToInviterMap[newInviteCode] = sender;
            inviterToInviteCodeMap[sender] = newInviteCode;
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

    function _calculateFruitRewardsProportion(
        uint256 rewardFruitCount
    ) private pure returns (uint256) {
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
