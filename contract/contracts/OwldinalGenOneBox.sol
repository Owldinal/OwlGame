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

    event MintBox(address indexed user, uint256 tokenId);
    event OpenBox(address indexed user, uint256 tokenId, BoxType boxType);

    constructor(
        ERC20Burnable _owlToken,
        OwlGame _gameContract
    ) ERC721("Owldinal Gen1 Box", "Gen1") {
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _grantRole(GAME_CONTRACT_ROLE, address(_gameContract));

        owlToken = _owlToken;
        gameContract = _gameContract;
    }

    function mintBox(
        address owner
    ) external onlyRole(GAME_CONTRACT_ROLE) returns (uint256 tokenId) {
        ++_tokenIdCounter;
        tokenId = _tokenIdCounter;
        _safeMint(owner, tokenId);

        emit MintBox(owner, tokenId);
        return tokenId;
    }

    function openBox(
        uint256 tokenId,
        bool hasBuff
    ) external onlyRole(GAME_CONTRACT_ROLE) returns (BoxType) {
        require(ownerOf(tokenId) == msg.sender, "You are not the owner");
        require(!isBoxOpened(tokenId), "Token has been opened");
        uint256 randomResult = Utils.generateRandomNumber() % 100;
        uint256 elfProbabilityThreshold = 10;
        uint256 fruitProbabilityThreshold = hasBuff ? 89 : 80;

        if (randomResult < elfProbabilityThreshold) {
            tokenBoxTypes[tokenId] = BoxType.ELF;
        } else if (
            randomResult < (elfProbabilityThreshold + fruitProbabilityThreshold)
        ) {
            tokenBoxTypes[tokenId] = BoxType.FRUIT;
        } else {
            _burn(tokenId);
            return BoxType.BURNED;
        }
        emit OpenBox(msg.sender, tokenId, tokenBoxTypes[tokenId]);

        return tokenBoxTypes[tokenId];
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
        for (uint i = 0; i < count; i++) {
            ++_tokenIdCounter;
            uint256 tokenId = _tokenIdCounter;
            _safeMint(owner, tokenId);

            tokenIdList[i] = tokenId;
            emit MintBox(owner, tokenId);

            uint256 randomResult = Utils.generateRandomNumber() % 100;
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

            emit OpenBox(msg.sender, tokenId, tokenBoxTypes[tokenId]);
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
}
