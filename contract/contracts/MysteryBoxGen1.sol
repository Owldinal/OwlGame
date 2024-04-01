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
    FRUIT
}

contract MysteryBoxGen1 is ERC721, AccessControl {
    bytes32 public constant GAME_CONTRACT_ROLE =
        keccak256("GAME_CONTRACT_ROLE");

    ERC20Burnable public owlToken;
    OwlGame public gameContract;

    uint256 private _tokenIdCounter;

    mapping(uint256 => BoxType) public boxTypes;

    event MintBox(address indexed user, uint256 boxId);
    event OpenBox(address indexed user, uint256 tokenId, BoxType boxType);

    constructor(
        ERC20Burnable _owlToken,
        OwlGame _gameContract
    ) ERC721("MysteryBoxGen1", "MBG1") {
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _grantRole(GAME_CONTRACT_ROLE, address(_gameContract));

        owlToken = _owlToken;
        gameContract = _gameContract;
    }

    function mintBox(address owner) external onlyRole(GAME_CONTRACT_ROLE) {
        _tokenIdCounter++;
        _safeMint(owner, _tokenIdCounter);

        emit MintBox(owner, _tokenIdCounter);
    }

    function openBox(
        uint256 tokenId,
        bool hasBuff
    ) external onlyRole(GAME_CONTRACT_ROLE) returns (BoxType) {
        require(ownerOf(tokenId) == msg.sender, "You are not the owner");
        require(!isBoxOpened(tokenId), "Token has been opened");
        uint256 randomResult = Utils.generateRandomNumber() % 100;
        uint256 isElf = 10;
        uint256 isFruit = hasBuff ? 89 : 80;

        if (randomResult < isElf) {
            boxTypes[tokenId] = BoxType.ELF;
        } else if (randomResult < (isElf + isFruit)) {
            boxTypes[tokenId] = BoxType.FRUIT;
        } else {
            _burn(tokenId);
            return BoxType.UNOPENED;
        }
        // gameContract.stakeNFT(msg.sender, tokenId);
        emit OpenBox(msg.sender, tokenId, boxTypes[tokenId]);

        return boxTypes[tokenId];
    }

    function isBoxOpened(uint256 tokenId) public view returns (bool) {
        return boxTypes[tokenId] != BoxType.UNOPENED;
    }

    function getBoxType(uint256 tokenId) external view returns (BoxType) {
        return boxTypes[tokenId];
    }

    // The following functions are overrides required by Solidity.

    function supportsInterface(
        bytes4 interfaceId
    ) public view override(ERC721, AccessControl) returns (bool) {
        return super.supportsInterface(interfaceId);
    }
}
