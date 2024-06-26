// SPDX-License-Identifier: MIT
// Compatible with OpenZeppelin Contracts ^5.0.0
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Burnable.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721Burnable.sol";
import "@openzeppelin/contracts/access/AccessControl.sol";
import "./OwlGame.sol";
import "./Utils.sol";

enum BoxType {
    UNOPENED, // not use.
    ELF,
    FRUIT,
    BURNED
}

contract OwldinalGenOneBox is ERC721, AccessControl {
    bytes32 public constant GAME_CONTRACT_ROLE =
        keccak256("GAME_CONTRACT_ROLE");

    ERC20Burnable public owlToken;
    OwlGame public gameContract;

    uint256 private _tokenIdCounter;

    mapping(uint256 => BoxType) public tokenBoxTypes;

    mapping(address => bool) public transferWhiteListMap;
    bool public isTransferEnable = false;

    event MintBox(address indexed user, uint256 tokenId, BoxType boxType);

    constructor(
        ERC20Burnable _owlToken,
        OwlGame _gameContract
    ) ERC721("Owldinal Gen1 Box", "Gen1") {
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _grantRole(GAME_CONTRACT_ROLE, address(_gameContract));

        owlToken = _owlToken;
        gameContract = _gameContract;
    }

    function setTransferEnable(
        bool isEnable
    ) external onlyRole(DEFAULT_ADMIN_ROLE) {
        isTransferEnable = isEnable;
    }

    function addTransferWhiteList(
        address[] calldata addrs
    ) external onlyRole(DEFAULT_ADMIN_ROLE) {
        for (uint i = 0; i < addrs.length; i++) {
            transferWhiteListMap[addrs[i]] = true;
        }
    }

    function deleteTransferWhiteList(
        address[] calldata addrs
    ) external onlyRole(DEFAULT_ADMIN_ROLE) {
        for (uint i = 0; i < addrs.length; i++) {
            delete transferWhiteListMap[addrs[i]];
        }
    }

    function mintAndOpenBoxes(
        address owner,
        uint256 count,
        bool hasBuff
    )
        external
        onlyRole(GAME_CONTRACT_ROLE)
        returns (uint256[] memory tokenIdList)
    {
        tokenIdList = new uint256[](count);
        uint256 seed = block.number;
        for (uint i = 0; i < count; i++) {
            ++_tokenIdCounter;
            uint256 tokenId = _tokenIdCounter;
            _safeMint(owner, tokenId);

            tokenIdList[i] = tokenId;

            seed = Utils.generateRandomNumberBySeed(seed);
            uint256 randomResult = seed % 100;
            uint256 elfProbabilityThreshold = 10;
            uint256 fruitProbabilityThreshold = hasBuff ? 89 : 80;

            if (randomResult < elfProbabilityThreshold) {
                tokenBoxTypes[tokenId] = BoxType.ELF;
            } else if (
                randomResult <
                (elfProbabilityThreshold + fruitProbabilityThreshold)
            ) {
                tokenBoxTypes[tokenId] = BoxType.FRUIT;
            } else {
                tokenBoxTypes[tokenId] = BoxType.BURNED;
                _burn(tokenId);
            }

            emit MintBox(owner, tokenId, tokenBoxTypes[tokenId]);
        }

        return tokenIdList;
    }

    function isBoxOpened(uint256 tokenId) public view returns (bool) {
        return tokenBoxTypes[tokenId] != BoxType.UNOPENED;
    }

    function getBoxType(uint256 tokenId) external view returns (BoxType) {
        return tokenBoxTypes[tokenId];
    }

    // The following functions are overrides required by Solidity.

    function supportsInterface(
        bytes4 interfaceId
    ) public view override(ERC721, AccessControl) returns (bool) {
        return super.supportsInterface(interfaceId);
    }

    // disable transfer
    function transferFrom(
        address from,
        address to,
        uint256 tokenId
    ) public override {
        if (!isTransferEnable) {
            require(
                transferWhiteListMap[msg.sender] ||
                    hasRole(DEFAULT_ADMIN_ROLE, msg.sender) ||
                    hasRole(GAME_CONTRACT_ROLE, msg.sender),
                "OwldinalGenOneBox: transfer disabled"
            );
        }
        super.transferFrom(from, to, tokenId);
    }

    function safeTransferFrom(
        address from,
        address to,
        uint256 tokenId,
        bytes memory _data
    ) public override {
        if (!isTransferEnable) {
            require(
                transferWhiteListMap[msg.sender] ||
                    hasRole(DEFAULT_ADMIN_ROLE, msg.sender) ||
                    hasRole(GAME_CONTRACT_ROLE, msg.sender),
                "OwldinalGenOneBox: transfer disabled"
            );
        }
        super.safeTransferFrom(from, to, tokenId, _data);
    }
}
