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
        2,
        3,
        4,
        5,
        6,
        7,
        8,
        9,
        10,
        11,
        12,
        13,
        14,
        15,
        16,
        17,
        18,
        19,
        20,
        21,
        22,
        23,
        24,
        25,
        26,
        27,
        28,
        29,
        30,
        31,
        32,
        33,
        34,
        35,
        36,
        37,
        38,
        39,
        40,
        41,
        42,
        43,
        44,
        45,
        46,
        47,
        48,
        49,
        50,
        51,
        52,
        53,
        54,
        55,
        56,
        57,
        58,
        59,
        60,
        61,
        62,
        63,
        64,
        65,
        66,
        67,
        68,
        69,
        70,
        71,
        72,
        73,
        74,
        75,
        76,
        77,
        78,
        79,
        80,
        81,
        82,
        83,
        84,
        85,
        86,
        87,
        88,
        89,
        90,
        91,
        92,
        93,
        94,
        95,
        96,
        97,
        98,
        99,
        100,
        101,
        102,
        103,
        104,
        105,
        106,
        107,
        108,
        109,
        110,
        111,
        112,
        113,
        114,
        115,
        116,
        117,
        118,
        119,
        120,
        121,
        122,
        123,
        124,
        125,
        126,
        127,
        128,
        129,
        130,
        131,
        132,
        133,
        134,
        135,
        136,
        137,
        138,
        139,
        140,
        141,
        142,
        143,
        144,
        145,
        146,
        147,
        148,
        149,
        150,
        151,
        152,
        153,
        154,
        155,
        156,
        157,
        158,
        159,
        160,
        161,
        162,
        163,
        164,
        165,
        166,
        167,
        168,
        169,
        170,
        171,
        172,
        173,
        174,
        175,
        176,
        177,
        178,
        179,
        180,
        181,
        182,
        183,
        184,
        185,
        186,
        187,
        188,
        189,
        190,
        191,
        192,
        193,
        194,
        195,
        196,
        197,
        198,
        199,
        200,
        201,
        202,
        203,
        204,
        205,
        206,
        207,
        208,
        209,
        210,
        211,
        212,
        213,
        214,
        215,
        216,
        217,
        218,
        219,
        220,
        221,
        222,
        223,
        224,
        225,
        226,
        227,
        228,
        229,
        230,
        231,
        232,
        233,
        234,
        235,
        236,
        237,
        238,
        239,
        240,
        241,
        242,
        243,
        244,
        245,
        246,
        247,
        248,
        249,
        250,
        251,
        252,
        253,
        254,
        255,
        256,
        257,
        258,
        259,
        260,
        261,
        262,
        263,
        264,
        265,
        266,
        267,
        268,
        269,
        270,
        271,
        272,
        273,
        274,
        275,
        276,
        277,
        278,
        279,
        280,
        281,
        282,
        283,
        284,
        285,
        286,
        287,
        288,
        289,
        290,
        291,
        292,
        293,
        294,
        295,
        296,
        297,
        298,
        299,
        300,
        301,
        302,
        303,
        304,
        305,
        306,
        307,
        308,
        309,
        310,
        311,
        312,
        313,
        314,
        315,
        316,
        317,
        318,
        319,
        320,
        321,
        322,
        323,
        324,
        325,
        326,
        327,
        328,
        329,
        330,
        331,
        332,
        333,
        334,
        335,
        336,
        337,
        338,
        339,
        340,
        341,
        342,
        343,
        344,
        345,
        346,
        347,
        348,
        349,
        350,
        351,
        352,
        353,
        354,
        355,
        356,
        357,
        358,
        359,
        360,
        361,
        362,
        363,
        364,
        365,
        366,
        367,
        368,
        369,
        370,
        371,
        372,
        373,
        374,
        375,
        376,
        377,
        378,
        379,
        380,
        381,
        382,
        383,
        384,
        385,
        386,
        387,
        388,
        389,
        390,
        391,
        392,
        393,
        394,
        395,
        396,
        397,
        398,
        399,
        400,
        401,
        402,
        403,
        404,
        405,
        406,
        407,
        408,
        409,
        410,
        411,
        412,
        413,
        414,
        415,
        416,
        417,
        418,
        419,
        420,
        421,
        422,
        423,
        424,
        425,
        426,
        427,
        428,
        429,
        430,
        431,
        432,
        433,
        434,
        435,
        436,
        437,
        438,
        439,
        440,
        441,
        442,
        443,
        444,
        445,
        446,
        447,
        448,
        449,
        450,
        451,
        452,
        453,
        454,
        455,
        456,
        457,
        458,
        459,
        460,
        461,
        462,
        463,
        464,
        465,
        466,
        467,
        468,
        469,
        470,
        471,
        472,
        473,
        474,
        475,
        476,
        477,
        478,
        479,
        480,
        481,
        482,
        483,
        484,
        485,
        486,
        487,
        488,
        489,
        490,
        491,
        492,
        493,
        494,
        495,
        496,
        497,
        498,
        499,
        500,
        501,
        502,
        503,
        504,
        505,
        506,
        507,
        508,
        509,
        510,
        511,
        512,
        513,
        514,
        515,
        516,
        517,
        518,
        519,
        520,
        521,
        522,
        523,
        524,
        525,
        526,
        527,
        528,
        529,
        530,
        531,
        532,
        533,
        534,
        535,
        536,
        537,
        538,
        539,
        540,
        541,
        542,
        543,
        544,
        545,
        546,
        547,
        548,
        549,
        550,
        551,
        552,
        553,
        554,
        555,
        556,
        557,
        558,
        559,
        560,
        561,
        562,
        563,
        564,
        565,
        566,
        567,
        568,
        569,
        570,
        571,
        572,
        573,
        574,
        575,
        576,
        577,
        578,
        579,
        580,
        581,
        582,
        583,
        584,
        585,
        586,
        587,
        588,
        589,
        590,
        591,
        592,
        593,
        594,
        595,
        596,
        597,
        598,
        599,
        600,
        601,
        602,
        603,
        604,
        605,
        606,
        607,
        608,
        609,
        610,
        611,
        612,
        613,
        614,
        615,
        616,
        617,
        618,
        619,
        620,
        621,
        622,
        623,
        624,
        625,
        626,
        627,
        628,
        629,
        630,
        631,
        632,
        633,
        634,
        635,
        636,
        637,
        638,
        639,
        640,
        641,
        642,
        643,
        644,
        645,
        646,
        647,
        648,
        649,
        650,
        651,
        652,
        653,
        654,
        655,
        656,
        657,
        658,
        659,
        660,
        661,
        662,
        663,
        664,
        665,
        666,
        667,
        668,
        669,
        670,
        671,
        672,
        673,
        674,
        675,
        676,
        677,
        678,
        679,
        680,
        681,
        682,
        683,
        684,
        685,
        686,
        687,
        688,
        689,
        690,
        691,
        692,
        693,
        694,
        695,
        696,
        697,
        698,
        699,
        700,
        701,
        702,
        703,
        704,
        705,
        706,
        707,
        708,
        709,
        710,
        711,
        712,
        713,
        714,
        715,
        716,
        717,
        718,
        719,
        720,
        721,
        722,
        723,
        724,
        725,
        726,
        727,
        728,
        729,
        730,
        731,
        732,
        733,
        734,
        735,
        736,
        737,
        738,
        739,
        740,
        741,
        742,
        743,
        744,
        745,
        746,
        747,
        748,
        749,
        750,
        751,
        752,
        753,
        754,
        755,
        756,
        757,
        758,
        759,
        760,
        761,
        762,
        763,
        764,
        765,
        766,
        767,
        768,
        769,
        770,
        771,
        772,
        773,
        774,
        775,
        776,
        777,
        778,
        779,
        780,
        781,
        782,
        783,
        784,
        785,
        786,
        787,
        788,
        789,
        790,
        791,
        792,
        793,
        794,
        795,
        796,
        797,
        798,
        799,
        800,
        801,
        802,
        803,
        804,
        805,
        806,
        807,
        808,
        809,
        810,
        811,
        812,
        813,
        814,
        815,
        816,
        817,
        818,
        819,
        820,
        821,
        822,
        823,
        824,
        825,
        826,
        827,
        828,
        829,
        830,
        831,
        832,
        833,
        834,
        835,
        836,
        837,
        838,
        839,
        840,
        841,
        842,
        843,
        844,
        845,
        846,
        847,
        848,
        849,
        850,
        851,
        852,
        853,
        854,
        855,
        856,
        857,
        858,
        859,
        860,
        861,
        862,
        863,
        864,
        865,
        866,
        867,
        868,
        869,
        870,
        871,
        872,
        873,
        874,
        875,
        876,
        877,
        878,
        879,
        880,
        881,
        882,
        883,
        884,
        885,
        886,
        887,
        888,
        889,
        890,
        891,
        892,
        893,
        894,
        895,
        896,
        897,
        898,
        899,
        900,
        901,
        902,
        903,
        904,
        905,
        906,
        907,
        908,
        909,
        910,
        911,
        912,
        913,
        914,
        915,
        916,
        917,
        918,
        919,
        920,
        921,
        922,
        923,
        924,
        925,
        926,
        927,
        928,
        929,
        930,
        931,
        932,
        933,
        934,
        935,
        936,
        937,
        938,
        939,
        940,
        941,
        942,
        943,
        944,
        945,
        946,
        947,
        948,
        949,
        950,
        951,
        952,
        953,
        954,
        955,
        956,
        957,
        958,
        959,
        960,
        961,
        962,
        963,
        964,
        965,
        966,
        967,
        968,
        969,
        970,
        971,
        972,
        973,
        974,
        975,
        976,
        977,
        978,
        979,
        980,
        981,
        982,
        983,
        984,
        985,
        986,
        987,
        988,
        989,
        990,
        991,
        992,
        993,
        994,
        995,
        996,
        997,
        998,
        999
    ];

    address[] private freeMintList = [
        address(0x0fe610C68eB24a7F3E625fE75eb6a4488A00bEc3),
        address(0x7Ce8BF832d0D6c62D8210204Ad7b6bfe30B8Cf48),
        address(0x9E42cAC4fEEAea193fB9F9D11A1dEca69bDE35Fc),
        address(0xFfB69f56dFDc9CE3Fa58ee7cC103530D75bF0eDA),
        address(0x34b5994F98144dd06E3C8044c7AcCfce87Bc63ED),
        address(0xFAac275596BC7FA7ee4a569Db4D5864Bb1EcFf0c),
        address(0xc86Bd7b0F0B3Dd3f4C96468Dc8C6eD928c67c3ba),
        address(0x488D4CAE745cF975EcaE86813AB3B77CF3636c22),
        address(0x6e0f6c5c5F1bfA1b3831613a40B03e7aF5c7E502),
        address(0xF93c81e26D96d91c5DfAB149b1D6ebBE04719C2C),
        address(0xA9BD48d3d5B96bf2FAFd1C530208ed2Cf360BEB0),
        address(0x7095C862a8db9AcbDDdE2f88Bfd1AFAc8E31375B),
        address(0x8F8509cd9F7b0Cf2F7F530Bc50Be3161553c7f22),
        address(0x40BAbd7c0D8296b5db9d6fd20Ab39E138a893e6B),
        address(0x924cCE40635EE704468Eddb13E261A20930a6C45),
        address(0x623c5d3A661Fad49d84FDDC90f0b7D34c2D0C457),
        address(0x8Df93733E2c8ba0202093BF20f8Ac5C1cC94819E),
        address(0xf019991714141dc1b192C2BEa0C4B0da4827C5C1),
        address(0x4ACD8A5A6B2572b26a51D046c880458eB8463B29),
        address(0x5aA9b7c23BE84EAd087E0Cab17CF0F748fd97155),
        address(0x3DFDB2b806750d6CeC9d1c8D6Ea8e44d31833880),
        address(0xe8ea4a00c0977DA96B08318d4562fe209E8b0E1F),
        address(0xed9265b8215cc9437726241Ab899496A90F7f50A),
        address(0x99cD34f26344BF3d6deC0029223D2CDF19637261),
        address(0x239A2dC13D9A5E2C722e33B58B2b5fD5f710f7b0),
        address(0x45BFab7Ca83795Ca65a679f260C358C6c6083020),
        address(0xd7FCc38aad07Cf341249f67e133fD8539DF3c99F),
        address(0x5f48a1c54b2a827d990e76B963Fa37CcF05c8e35),
        address(0x40006d4A2f3F4eABe8bBB0aebC8c158835dB1084),
        address(0x60E483cdCf8Ca76Cb83594675caD74fe3D6F97eC),
        address(0x7e5861C3D70a736D921b62E996F2f60537eF17F1),
        address(0x9bC3607175f5D7811cc3366a4521a95fB9f858Ea),
        address(0x1E12DD62Ed11E803B7c50D910E0f719e497Ce0aE),
        address(0x9a24C8dd696eE650124eA801fD84CdfdC802280C),
        address(0xf9d362cFF11cd6B30f8E3A4CFcDe2871224f0643),
        address(0x79f71E9464F257F2a4c232B0c1D6f8131DcC23D3),
        address(0x40C76e11A996A71E1697ed207Db45E1c337e0472),
        address(0xF88Ba6c35C802F77390Ce0A3d9c0A2c67aF3E342),
        address(0x224C111C0715dc18dB69692ad8264751fcda8390),
        address(0x8299810B34e41F78039863d05ad51D9aBa7391AB),
        address(0xB562822A6dB7676D67C31147ededdf58C711907f),
        address(0x32c8c294aFf1103209ae3F59acB741F270E020AC),
        address(0x61fd0D043d519F5A2bD05785000f30Db96809429),
        address(0x04EE22568B4ABBfF87a6827BC4f801b81D99146B),
        address(0x49380A694D57696Ac6f2E5181c8Ef5233b8bFF8F),
        address(0x87d63B96ca7695775fddE18ADE27480143F9dfA0),
        address(0xc2fF0E93503ff5a9C6D7F2f26be9e9804F5f89c2),
        address(0x7Ac389A8a387861FcBC0D171B89d7e978680a1f0),
        address(0xC373157CB678B1A079dDea3870D31D1F7283F273),
        address(0x6ac8141DAbDe726E9EE85d09B28A3E441230a505),
        address(0x17022f3D81a8D73AFC489E880017ED3595294c67),
        address(0x220D269f0b2e94739505f367c3FBB4f6b2f824fb),
        address(0x49B3735bB4C0E01846996702b167e0E6b64f1ac7),
        address(0x8c00cbA2694Bcc5bd5C3928adF490b43f19D99a9),
        address(0x7f3551826e26aDbaD6f3907abcB7247277e391FF),
        address(0x7cB1943A25ee94b5dc34EFBbC0EaBc23082bA870),
        address(0x9BaEd1D5f42a5Bfe2E8293eD0B0c9b028E2F2aDd),
        address(0x2DfF3122eE9CFD45EC6954aA21A39C5B0a743da4),
        address(0xf4BAB82bB7bfabe831A7e0e0166d832BCFD71265),
        address(0x035535b9Cb724E88d647d5D42fF0aD2cFe9d75AF),
        address(0x2D1c9e7DB8E757a2fD6973E685f08D8D8416A422),
        address(0x4A06F047FEBbD4A3c97809492c9f5C10B9B9150e),
        address(0x34Bc905d74bc04Ea0C41A79D64bf2a8c153766dD),
        address(0x22A85880abc847f07bdC00d3434Ea3032f53dB02),
        address(0xa264E3F7c2D1d632A8743d7FEB0C6FD0F0E534B7),
        address(0xf61426a7FE335907710d0d2a81aA7e5d461930Ad),
        address(0x87D9D46C002a3b52b55FF8c7AbB48ebB72737342),
        address(0xeD903FB3C360aC7BA0640A2F0754f8b362E250f8),
        address(0xe76B4F8C8fc7A81A3869A2cb766aC20d41ABb1F8),
        address(0xF743b9F77ceF50549AE16B10571d76446ea5BA37),
        address(0x93383277928B75510eEB06Eda43334B87dA3b4B7),
        address(0x2C7a8b681F6A1F7d418706a7f5F8d63e183D5F9c),
        address(0xe7Aba96e21109eE0A99d6CfC4a852214Bc9F9D8E),
        address(0xDAC5352649e7A03a42D7D73C8C31A6bECEed100A),
        address(0x00794b7B900165F452e2b02cC191e6b8B6F8F869),
        address(0x804500dD2259d456d56924e94A2A33b670fd10A4),
        address(0x8fb4968F75D5DE9032CB8e66d10D7Ca05eA098ac),
        address(0x4269365f001aB881b1d298acc8088d2Bcc39eb94),
        address(0xDb7A4c6C10D341f97736C1bF2ecf1173A61A26b3),
        address(0x9dd6C2cdd86B507e7EE1AE605AD034eBC1c726aF),
        address(0xe4C2ccf40d0041613a64a65b3846aA590cea0e6e),
        address(0x8fBEEEC4F3f6090D318497ee87c11316C4efDBDC),
        address(0x39fB547D00C72D5735eBB08120482CA258cD5fAE),
        address(0x8A4BFA2e4dC5eD72c22a357F653549b02689c18f),
        address(0x1aDcBCc69D99E4e89876b06FAb865043Fc6B8F1e),
        address(0x6007b3fd31AFf7775B2DC2a10EB6ac82a96FE462),
        address(0x6299975EAa0a20a21C7a01Ca87D97db118c7E28c),
        address(0xc8BC1951b93E7ede6Fcb31FEd7Da94A6D134e962),
        address(0xAB4258C5CDA2DF96BE5Fb897a0B5a3d1c9e06FEc),
        address(0x412728cfb6210eE674Bd54b747d0b861E26268C6),
        address(0x91244d6D3A27a74A37Aa8D5541157ecf1e2eA1C0),
        address(0x3385Bc434D049d9C1D2c6DBd369407ce06ff0649),
        address(0xcDb8c5185A341eF78f27619c3A739f2a14A6253f),
        address(0x2cD85Bf58259Eeb9C0d1D741Fa4a5Bf0dec24b1c),
        address(0x4A808188C0572F23bdDf1dD6aE1cA745719804f3),
        address(0x7174924BE343aA46B7C82EBB2b98272725Eeb6dF),
        address(0xf792B5c9e0CdA70F0240A6C4D9AFe7e24767e545),
        address(0x4179f93Bf5f6a582B93eDdE5C7186d49beC99db7),
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

    function mintByAdmin(
        bytes memory signature
    )
        external
        nonReentrant
        onlyRole(DEFAULT_ADMIN_ROLE)
        returns (uint256 tokenId)
    {
        _validMint(signature);
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
