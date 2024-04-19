// SPDX-License-Identifier: MIT
// Compatible with OpenZeppelin Contracts ^5.0.0
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721Burnable.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "@openzeppelin/contracts/utils/cryptography/MessageHashUtils.sol";
import "@openzeppelin/contracts/utils/ReentrancyGuard.sol";
import "@openzeppelin/contracts/access/AccessControl.sol";
import "./Utils.sol";

contract Owldinal is ERC721URIStorage, AccessControl, ReentrancyGuard {
    address private immutable _idOneOwner;
    address private immutable _defender;
    IERC20 private immutable _voyaToken;

    uint256 public constant MINT_PRICE = 45 * 10 ** 14; // 0.0045
    uint256 public constant MINT_VOYA_THRESHOLD = 50 * 10 ** 18;

    uint256 public tokenIdCounter = 1;
    mapping(address => bool) private hasMinted;

    uint256 public remainVoyaCount = 900;
    uint256 public whiteListEndedBlock;

    // URI 1 is reserved
    uint256[] private _availableTokenUris = [
        1
    ];

    address[] private freeMintList = [
        address(0x816719e7B7c425d36629d1bac35d53Aeb627aAf9)
    ];

    // Events
    event MintBox(
        address indexed user,
        uint256 boxId,
        uint256 mintType,
        string url
    ); // minttype: 1=whitelist, 2=voya, 3=tokenOne

    constructor(
        uint256 endedBlock,
        address voyaTokenAddr,
        address idOneOwnerAddr,
        address defenderAddr
    ) ERC721("Owldinal", "Ow") {
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        whiteListEndedBlock = endedBlock;
        _voyaToken = IERC20(voyaTokenAddr);
        _idOneOwner = idOneOwnerAddr;
        _defender = defenderAddr;
    }

    function checkCanMint(
        address addr
    ) external view returns (bool canMint, uint256 mintType) {
        if (hasMinted[addr]) {
            return (false, 0);
        }

        // check if can mint token one
        if (addr == _idOneOwner && _ownerOf(1) == address(0)) {
            return (true, 3);
        } else if (block.number < whiteListEndedBlock) {
            // check if can mint by whitelist
            uint256 length = freeMintList.length;
            for (uint256 i = 0; i < length; i++) {
                if (freeMintList[i] == addr) {
                    return (true, 1);
                }
            }
        }

        // check if can mint by voya
        if (
            remainVoyaCount > 0 &&
            _voyaToken.balanceOf(addr) >= MINT_VOYA_THRESHOLD
        ) {
            return (true, 2);
        }

        return (false, 0);
    }

    function validSignature(
        bytes memory signature
    ) external nonReentrant returns (bool) {
        _validMint(signature);
        return true;
    }

    function mintByAdmin()
        external
        // bytes memory signature
        nonReentrant
        onlyRole(DEFAULT_ADMIN_ROLE)
        returns (uint256 tokenId)
    {
        // _validMint(signature);
        ++tokenIdCounter;
        _safeMint(msg.sender, tokenIdCounter);
        emit MintBox(msg.sender, tokenIdCounter, 3, tokenURI(tokenIdCounter));
        return tokenIdCounter;
    }

    function mint(
        bytes memory signature
    )
        external
        payable
        nonReentrant
        returns (uint256 tokenId, uint256 mintType)
    {
        require(msg.value >= MINT_PRICE, "Insufficient BTC sent");
        require(!hasMinted[msg.sender], "You have already minted a box");
        _validMint(signature);

        // check if can mint token one
        if (msg.sender == _idOneOwner) {
            _safeMint(msg.sender, 1);
            hasMinted[msg.sender] = true;

            string memory uri = _generateTokenURI(tokenIdCounter);
            _setTokenURI(tokenIdCounter, uri);

            emit MintBox(msg.sender, 1, 3, tokenURI(tokenIdCounter));
            return (1, 3);
        }

        // check if can mint by whitelist
        if (block.number < whiteListEndedBlock) {
            bool isWhitelisted = false;
            uint256 length = freeMintList.length;
            for (uint256 i = 0; i < length; i++) {
                if (freeMintList[i] == msg.sender) {
                    isWhitelisted = true;
                    freeMintList[i] = freeMintList[length - 1];
                    freeMintList.pop();
                    break;
                }
            }

            if (isWhitelisted) {
                ++tokenIdCounter;
                _safeMint(msg.sender, tokenIdCounter);
                hasMinted[msg.sender] = true;
                string memory uri = _generateTokenURI(tokenIdCounter);
                _setTokenURI(tokenIdCounter, uri);

                emit MintBox(
                    msg.sender,
                    tokenIdCounter,
                    1,
                    tokenURI(tokenIdCounter)
                );
                return (tokenIdCounter, 1);
            }
        }

        // check if can mint by voya
        if (remainVoyaCount > 0) {
            if (_voyaToken.balanceOf(msg.sender) >= MINT_VOYA_THRESHOLD) {
                --remainVoyaCount;
                ++tokenIdCounter;
                hasMinted[msg.sender] = true;
                _safeMint(msg.sender, tokenIdCounter);
                string memory uri = _generateTokenURI(tokenIdCounter);
                _setTokenURI(tokenIdCounter, uri);
                emit MintBox(
                    msg.sender,
                    tokenIdCounter,
                    2,
                    tokenURI(tokenIdCounter)
                );
                return (tokenIdCounter, 2);
            }
        }

        revert("Not eligible to mint");
    }

    // admin
    function withdraw(
        address payable recipient
    ) external onlyRole(DEFAULT_ADMIN_ROLE) returns (uint256) {
        require(recipient != address(0), "Cannot withdraw to the zero address");
        uint256 balance = address(this).balance;
        require(balance > 0, "No funds to withdraw");

        (bool success, ) = recipient.call{value: balance}("");
        require(success, "Transfer failed");
        return balance;
    }

    function getCurrentBalance()
        external
        view
        onlyRole(DEFAULT_ADMIN_ROLE)
        returns (uint256)
    {
        return address(this).balance;
    }

    receive() external payable {}

    fallback() external payable {}

    // private

    function _generateTokenURI(
        uint256 tokenId
    ) internal returns (string memory) {
        require(_availableTokenUris.length > 0, "No available token uris");

        if (tokenId == 1) {
            return "1";
        } else {
            uint256 length = _availableTokenUris.length;
            uint256 index = Utils.generateRandomNumber() % length;
            uint256 result = _availableTokenUris[index];

            _availableTokenUris[index] = _availableTokenUris[length - 1];
            _availableTokenUris.pop();

            return Strings.toString(result);
        }
    }

    function _baseURI() internal pure override returns (string memory) {
        return
            "https://ipfs.io/ipfs/QmUJuZpBdkdjFYzRNNSsAG58Zr2eLKFZMqUKfVu33FF6Wt/";
    }

    function _validMint(bytes memory signature) internal view {
        bytes32 hash = keccak256(abi.encode(msg.sender, address(this)));
        require(
            ECDSA.recover(
                MessageHashUtils.toEthSignedMessageHash(hash),
                signature
            ) == _defender,
            "Invalid signature"
        );
    }

    // The following functions are overrides required by Solidity.

    function supportsInterface(
        bytes4 interfaceId
    ) public view override(ERC721URIStorage, AccessControl) returns (bool) {
        return super.supportsInterface(interfaceId);
    }
}
