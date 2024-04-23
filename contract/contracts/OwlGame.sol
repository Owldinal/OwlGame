// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/access/AccessControl.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
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

    ERC20 public owlToken;
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

    address deadAddress = address(0x00000000000000000000000000000000dEaD);

    struct MintRequest {
        address user;
        uint256 count;
    }
    uint256 mintRequestCounter;
    // request id => mint request
    mapping(uint256 => MintRequest) public requestMintMap;

    bool public isMoonBoostEnable;
    address[] private moonBoostWhiteList = [
        address(0xd0014F8bEc86c7B236fFEF140271E159Dc5D36f9),
        address(0xb4009C1096ca903966212a3096aFC88F2bA9088B),
        address(0xA51145D23fBA581fE45C9202a1A1BfDCC04b248c),
        address(0xD5F773FA2E57B7f838210a738cfb27C30507fbce),
        address(0x49088cF5e14Ce09b3ae312412dF7a8FE42649ABb),
        address(0x37E4E1b56bC1a9cdc5773AA0c6C05793ffF4f931),
        address(0xB2B3fD51b923bdF9d405fd0B192c868784299382),
        address(0xDC2743CF9aFD3B7F3c0a49594099749271925767),
        address(0x180F81ceAd28209E53Cdb6AA7F8F2F29C3f32B04),
        address(0xbc0ee223287656b4cD53DA52d0ac41Bd677Fa851),
        address(0x9EE41Ea35da2FD1e01D136EC06F8951E50156175),
        address(0xFE677893436962D0D827aCa15C114DEFa8717B9D),
        address(0xd5C38625d05b58C67baeA227f29dfe23EBDE7bCE),
        address(0xFe1Ca5312796f7BEF15eEb0Aa3aA7b59098957A7),
        address(0xF2FF9961C6Cf369FB796B4086E53ED04D0D65443),
        address(0x4a44f56ae5706Dd0abB211fa5C704127b7CA8617),
        address(0x6cf2f3358f69B481a3cBc08D45137a4028530Dbe),
        address(0x24056D87e1cC406293808fff79058ea050D276d8),
        address(0xCcE3344a3438c4FF048e07c73B1860AD25d3c226),
        address(0xC134f280C35dFD5ccD47Ea1584e987aB64389fC6),
        address(0x1B1D24A623371034B4c4a42aaf67Aa863075f5d8),
        address(0x058295C1cf87CBb5cb2090dda780B55f227cd5f9),
        address(0xf032DA00C6164343F1886C01D706deB909A40A54),
        address(0xB2D82092e61d342a22F372078C62b81F8c3E86C9),
        address(0x2fa41a2B6447924591244e5ac120283aB8914967),
        address(0x4948c87b7c1330c61cA41FAbE1642d38E1bcA248),
        address(0x5e7236E6D99171c7fF9d76572b5008aa6C58E131),
        address(0xC6aEc21Ab251b83204F5B337C79933a164e2beC0),
        address(0xc547aC99A29526d8615AeE8C929aa997f3DF18D3),
        address(0x24308E3FAF2c1c589022d0f8310895Db8F2D0f6b),
        address(0x9f7Fe8aEf1d6ebB87e25E5B68668b679696eaFbB),
        address(0x9F9349c172601D88C620E06e0155bE68698C562A),
        address(0x051FD6582C45BF7e8Fd6d8648458480861D5c84f),
        address(0x7be2e48d2C814735AaD2dd329a355FaB9c209295),
        address(0x40eBe9858a9b5BE3d3Ef67Bb44611Df8DCef1F3c),
        address(0xB8689c766c876fd18c7E016A334A9166f64d949A),
        address(0x40eF7b510385C0e24672bFFb01A552CBD393167A),
        address(0xD216B099167b8a12346eF5b6155A74aEC4D296dA),
        address(0xFe6F936B26F1f269f27EFB5Dc64b49B4f915BAe7),
        address(0x13e3B73F64b3adE38b1730645041a01F71CF873D),
        address(0xAf914e9DC09644e020A732986F361ecdD4e3c6DF),
        address(0x650879056d32Ea053b7F3F456466B0E11BEAF2E8),
        address(0x0C7c975069e94D8F068bd83d29C79D97a6FFC6D3),
        address(0x1bbDe1E811Ca33E0aA5834A40bA4FC0B066EC05E),
        address(0xd3e1815369A59Cb5fB6dDA64bCB6a8F4885011fd),
        address(0x05378Ad38aD1dd399082e2a04F11980758aaF06F),
        address(0xD1BdaeA3fbe4e2194c7CB78Fd1EC0Af7B9ff77DC),
        address(0xF16904455727e0513405062afbAF6501e104Ee9b),
        address(0x969f1d2112389FE1A7ad9feCeE45715e037EAD7d),
        address(0xbA0db8625201a84E773ABdEc7dD0711B790F1e78),
        address(0x90DCa53C1a6d176147FaB3E30a7e6436214513F5),
        address(0x52a7cD08ec1157b4Af4bEB247Bc95ABb53023998),
        address(0xcBf440Fcb5DA3a767B286a746a286fDFBb3bD98c),
        address(0xf1C79ADCBe2EC98825ec1eeD7Bf72D4668ef64f0),
        address(0x6636b62d9c676ca5a005A55dB918a5B2Ff443711),
        address(0xb931ee787ADF24E2e9446d96f92529ef8942aC15),
        address(0xdd29E1F550231Ba3F247e2f0567F653e59BD25d1),
        address(0xA2fb1a3B295980725529a764A422FB7b7Ad4a8d7),
        address(0x4dC23a54B7DAA68B842B45cc6aDe5CC3f141d7C2),
        address(0xb3683389045b31160454D30EbECe9552ceB6c002),
        address(0x90B6a36513b554caFDD2208EFa69Cc7b0B79d203),
        address(0x21A784156F45C034faA085F3da500D7028241c43),
        address(0x98E074D1cDab84F04B29d36ba20f91c274C8E1b1),
        address(0x6396eD9CEb07FB369d48257bE036fd5f059fB8FA),
        address(0xA1050156F6C9e092f21f26ABE02F44CfF2C91465),
        address(0x69096890bb79E605f5B7E281915a812e629fEAf2),
        address(0x04FB15CD7963b38D1C53DdD0e2Fd1388Bf75bDC0),
        address(0x0F0443c5c315B2E878a9e71b58E3BDBEBb400870),
        address(0x77699D2C22CaAbAd35dba993b859432650E81e07),
        address(0x995f1e0F842aCD5B01639FB2E6Cd0bcaf182dE57),
        address(0xE4212BA9186D01F275648CABEdc9a4EdC29a48B9),
        address(0x1Da9144298FA79cbaAb55b2ff521f0B011D909F7),
        address(0x0e867aF6E3Aa2A0789bc250D5855aDDF320Cf5C6),
        address(0xeaDc3d91F18fb34fEBFbCFb5203d05b2D96e52Ef),
        address(0xAf5C65DFe532d0C79fe8B1980BA4CecA8D588393),
        address(0xa743eea6B3fd8f4bB1861A86C9D99c6AB76E6cd1),
        address(0xC12E7861B464eF044E7BddE43894B0D388E9BeaE),
        address(0x01800797Bc9303B7c80Caa8E5E881FD8EdaC1D61),
        address(0xa4F408CD75F1cce8bc4cc2Fd9689646e69a6FdeD),
        address(0x3cC49318CC5304c8CDB158e1D414B452aa7B72Ad),
        address(0xF5b3b0bc1FCDB437d4546D40B766521614207b09),
        address(0x3F64001FD810Ae8165EA5727980EA733E1d0c378),
        address(0x0f1D9687b779e120fDEf5507b20ebECf5335C377),
        address(0xCeB6e43bB274286b552f9ACCE7Fb20C82b6b8872),
        address(0xa1f81A6Ca3a6E2a942Dc3371b2AcF2BD58834c01),
        address(0x2b5a56D7F19e2a9A641E6bE2509aEC89489703C6),
        address(0x04FB15CD7963b38D1C53DdD0e2Fd1388Bf75bDC0),
        address(0xFCb0208a092F0a106bDdf62f52223E867B87b401),
        address(0xCeB6e43bB274286b552f9ACCE7Fb20C82b6b8872),
        address(0x6ac3Dc95d4D4Feb9DefD47A0fcb8e472b60a098c),
        address(0xaCE87dde9D8bBA8fBf024A0973C17E0654ee3493),
        address(0x043f99Af377BC8fB42a8C3431cbdd672d9782b0F),
        address(0x3c8653496E3A9eb2FFbAB931Eb5067f126b6a9cd),
        address(0xcb8Bc674f11d391d1Bad2F53257016b02BfCC580),
        address(0xeabb6C0ec2c8b2C62681e83Bb465397eEA8fA780),
        address(0x4B59Bc41FEA567dC7A389e07D2DCAFEc1A589FD5),
        address(0x3dd744E3B1123142aA31bC1AEed3059d8A5Be661),
        address(0x64D709dCD8793682CEfEBEC58cfD97e3DC7DBe47),
        address(0x70367fc3115F5f42CB22A7633c05f830E19fD5f6),
        address(0x63EACecF428615eC2Cadc90381e71dE82eA77777),
        address(0xf90b156E8c094A2348D66B669B6e79A407F9A067),
        address(0xC7791a20d6Ca4b35DD23Ef9c5c97f939FA504642),
        address(0xdF4379735Ea99D9fD5b6a4CDEA517a5E6e20E44E),
        address(0xc300baBc3cd259bb7fdce31779f32c53F3131B9c),
        address(0x1A3b8ad2a44c472AFa6Ef3B32FAe35518a92B7fC),
        address(0x49779B21B659BD334b24a2aA94D1902ad028Dce5),
        address(0xBB30AF7F53114A35a33b0C5dE91890f4AAbfD117),
        address(0x834a4240b6106dB238d90D57aBe88772AC2b80fA),
        address(0xba03B1bbdbdAD85922ADBDfBAa870909226d6C34),
        address(0xeB4bA061B0daD79e0c17Db259FE81279A84Fb22B),
        address(0xDB4dcEBbB5B31341e9Bd52172ad27d675FeE99D1),
        address(0x83495590CBEf1Bb7e835f035143C81F1D69DAAD3),
        address(0xDf1ce9Bc2E935750E4bc4D6a2FA15C1c13A7032C),
        address(0xC5450E7E7c52574fe1D6bA9F554fD0faA0dcA313),
        address(0xc9E558f16D219BBE445312D373b36Cce73A03939),
        address(0x3F64001FD810Ae8165EA5727980EA733E1d0c378),
        address(0x656dD06dfF9C1F04F5dF473F0bE210473Ab8C51c),
        address(0xE4212BA9186D01F275648CABEdc9a4EdC29a48B9),
        address(0x742B7c268C5378E0cbF1153B81609580c9069F25),
        address(0x0329FD90E32dA292eE9F663C17A283Ca11Ca4039),
        address(0x30317c1D95bb92cc4D5984Eb695056f883449831),
        address(0xCeB6e43bB274286b552f9ACCE7Fb20C82b6b8872),
        address(0xdDB4Bf4BeeAa68672A83db1789fAb53c24D84247),
        address(0xA3F98dcDdB8f7d14BC733bC202B1bcbD9Bce84B5),
        address(0x236dCf7295eAC62e0373f1af8D469ADDA305b878),
        address(0x1C23Ba0Bd17123D911002135e9c48beA7A5857E3),
        address(0xa62663E5Fb5681882277e0501a6931f067E9AA42),
        address(0x80327002b6c35657B8367290B22c2DC69A399b84),
        address(0xF66a84D5b8df529eaa96c775f0fB6297D913C594),
        address(0x6CE580B0E9081bd60Ba8D11F37F4Eed1796556b8),
        address(0x91CDE8AF2ceb804B9e17D6A4FCc45dd2B6976173),
        address(0xA9D01fe9C3734543203999b3e4D1FcA08CE9Dff3),
        address(0x4695f9e89c26cde1B2c6321d19b5374f3d1b9aa8),
        address(0x12d244711A7b3Af55058549f6599a4603C8DEeDA),
        address(0xbf8b8784DA646b3122292442411954bF0f8f538b),
        address(0x03a508224Dd9e7b1a9660dBfc09c1534B8D9a22A),
        address(0xb910189d34B6beF7d1015CF677EA5101b1d6C389),
        address(0x47106128f0cE8D7092043e0b564F4D6086209E20),
        address(0xA3F98dcDdB8f7d14BC733bC202B1bcbD9Bce84B5),
        address(0x29348Fb00EBCdC2a2110efff40CEd988492bA5c0),
        address(0x560F75c77eef60542EC8f85F71993a941Cd48d78),
        address(0x4155cB3aDbA70DA55A24C002C22242bD25c4F878),
        address(0x94dE3bc43f70bA4a1b5dB91248997551Ec88B213),
        address(0xe628c97177061168ec9b7f43140f9a84710d0B2d),
        address(0x9f98932881E67C003B3a27b1622092d35149c867),
        address(0xA9D01fe9C3734543203999b3e4D1FcA08CE9Dff3),
        address(0x8d13366291A7B5fA25cDF7A43f94f864e25f8482),
        address(0x639fc9A9E4bC3F0BB44f0e79bd0E18DAb22592Fe),
        address(0x356e1Aae345b0b8B1dedC5fFD8F9de33f7c22Cd5),
        address(0x9008338926a799F3A028cEaB4B0bcaF74aEF0725),
        address(0x88509ccD66a55850cF0189e5E9d713c097Ad1f6a),
        address(0x7AD34c497814C3660455c336af870B5c32EE71a3),
        address(0xBa30fa89Cbb088dB75e37F77c230C944da26738A),
        address(0xa84a24E9264068f4adBc1467da250e3672b4c7bF),
        address(0xc0C4e04B435185634066C282E3d9b46CE843FA03),
        address(0xBE8E58E212BCf934F7297E71989f3D0a1C0BA81A),
        address(0x4BaD4C67A9704f6df012ff750d80f9C3832102a3),
        address(0xDdce2e9F2427Bf2182Bd6B409Fc5b5B46878aC5E),
        address(0x69Dbdd6Df8d96bAaCd0920dD2887071403920b11),
        address(0x695ad4733c3A1f7fa860653A81AFb84ef77C5670),
        address(0xB9eE36B20d9Fc0c7B61a01Dc499FFcfdCA244da7),
        address(0x3f6171A71514A9eB63bFf826F89F8e2F58fc1FAd),
        address(0x6071AF68Ce36435B6Ce7BA290E039bf2b42791F9),
        address(0x3db83a0912070c4D9eA0ADbE8564eC82283b3f30),
        address(0xB3A136831AEb9D327DdfB4b62482cf3d8686b6f8),
        address(0x1105b324d8b6fB0aE3960E043980bBbE692d244C),
        address(0x285d03Bc3848580BBF712529547749906043C554),
        address(0xd38B6F6cC4649029CFAFD6DEe510de7240aa86aA),
        address(0x4CF7096c0d9C5420Ab6e5EEdD127BBA7f4FA007E),
        address(0x17bB636aCcCADE156C67d88808108387e1f8F1c0),
        address(0x409D2022eD5E0c77Be6987d1817512C3BB218888),
        address(0x4A8553A16835AfFaB99D4B6172765Ae1171C309A),
        address(0x27361dbEB913A810D9A3070043658f694Afb2164)
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
        owlToken = ERC20(owlTokenAddr);
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
    ) external nonReentrant returns (address inviter) {
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
        owlToken.transfer(deadAddress, burnAmount);
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
    function stakeOwldinalNft(
        uint256[] calldata tokenIdList
    ) external nonReentrant {
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

    function stakeMysteryBox(
        uint256[] calldata tokenIdList
    ) external nonReentrant {
        require(tokenIdList.length > 0, "Param is empty");
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

    function unstakeOwldinalNft(
        uint256[] calldata tokenIdList
    ) external nonReentrant {
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
    ) external nonReentrant {
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
            owlToken.transfer(deadAddress, totalRewardsToBurn);
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

    function claimInviterReward() external nonReentrant {
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
    function updateAllFruitRewards()
        external
        nonReentrant
        onlyRole(SERVER_ROLE)
    {
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
