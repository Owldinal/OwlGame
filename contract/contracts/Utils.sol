// Utils.sol
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

library Utils {
    function generateRandomNumber() internal view returns (uint256 rand) {
        uint256 blocknumber = block.number;
        uint256 mod = blocknumber > 255 ? 255 : blocknumber - 1;
        uint256 random_gap = uint256(
            keccak256(abi.encodePacked(blockhash(blocknumber - 1), msg.sender))
        ) % mod;
        uint256 random_block = blocknumber - 1 - random_gap;

        bytes32 sha = keccak256(
            abi.encodePacked(
                blockhash(random_block),
                msg.sender,
                block.coinbase,
                block.prevrandao
            )
        );
        return uint256(sha);
    }

    function generateRandomNumberBySeed(
        uint256 seed
    ) internal view returns (uint256 rand) {
        uint256 blocknumber = seed;
        uint256 mod = blocknumber - 1;
        uint256 random_gap = uint256(
            keccak256(abi.encodePacked(blockhash(blocknumber - 1), msg.sender))
        ) % mod;
        uint256 random_block = blocknumber - 1 - random_gap;

        bytes32 sha = keccak256(
            abi.encodePacked(
                blockhash(random_block),
                seed,
                msg.sender,
                block.coinbase,
                block.prevrandao
            )
        );
        return uint256(sha);
    }

    function removeValue(
        uint256[] storage array,
        uint256 valueToRemove
    ) internal {
        uint256 length = array.length;

        for (uint256 i = 0; i < length; i++) {
            if (array[i] == valueToRemove) {
                array[i] = array[length - 1];
                array.pop();
                return;
            }
        }
    }
}
