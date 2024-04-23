package service

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"owl-backend/abigen"
	"owl-backend/internal/config"
	"owl-backend/internal/constant"
	"owl-backend/internal/database"
	"owl-backend/internal/eth"
	"owl-backend/internal/model"
	"owl-backend/pkg/log"
)

var (
	owlTokenContract *abigen.OwlToken
	owlGameContract  *abigen.OwlGame
	moonBoostAddress map[string]bool
)

func init() {
	var err error
	owlTokenContract, err = abigen.NewOwlToken(common.HexToAddress(config.C.OwlTokenAddr), eth.Client)
	if err != nil {
		log.Fatalf("Failed init user service: %v", err)
	}

	owlGameContract, err = abigen.NewOwlGame(common.HexToAddress(config.C.OwlGameAddr), eth.Client)

	// init moon boost list.
	moonBoostAddress = make(map[string]bool)
	addressList := []string{
		"0xd0014F8bEc86c7B236fFEF140271E159Dc5D36f9",
		"0xb4009C1096ca903966212a3096aFC88F2bA9088B",
		"0xA51145D23fBA581fE45C9202a1A1BfDCC04b248c",
		"0xD5F773FA2E57B7f838210a738cfb27C30507fbce",
		"0x49088cF5e14Ce09b3ae312412dF7a8FE42649ABb",
		"0x37E4E1b56bC1a9cdc5773AA0c6C05793ffF4f931",
		"0xB2B3fD51b923bdF9d405fd0B192c868784299382",
		"0xDC2743CF9aFD3B7F3c0a49594099749271925767",
		"0x180F81ceAd28209E53Cdb6AA7F8F2F29C3f32B04",
		"0xbc0ee223287656b4cD53DA52d0ac41Bd677Fa851",
		"0x9EE41Ea35da2FD1e01D136EC06F8951E50156175",
		"0xFE677893436962D0D827aCa15C114DEFa8717B9D",
		"0xd5C38625d05b58C67baeA227f29dfe23EBDE7bCE",
		"0xFe1Ca5312796f7BEF15eEb0Aa3aA7b59098957A7",
		"0xF2FF9961C6Cf369FB796B4086E53ED04D0D65443",
		"0x4a44f56ae5706Dd0abB211fa5C704127b7CA8617",
		"0x6cf2f3358f69B481a3cBc08D45137a4028530Dbe",
		"0x24056D87e1cC406293808fff79058ea050D276d8",
		"0xCcE3344a3438c4FF048e07c73B1860AD25d3c226",
		"0xC134f280C35dFD5ccD47Ea1584e987aB64389fC6",
		"0x1B1D24A623371034B4c4a42aaf67Aa863075f5d8",
		"0x058295C1cf87CBb5cb2090dda780B55f227cd5f9",
		"0xf032DA00C6164343F1886C01D706deB909A40A54",
		"0xB2D82092e61d342a22F372078C62b81F8c3E86C9",
		"0x2fa41a2B6447924591244e5ac120283aB8914967",
		"0x4948c87b7c1330c61cA41FAbE1642d38E1bcA248",
		"0x5e7236E6D99171c7fF9d76572b5008aa6C58E131",
		"0xC6aEc21Ab251b83204F5B337C79933a164e2beC0",
		"0xc547aC99A29526d8615AeE8C929aa997f3DF18D3",
		"0x24308E3FAF2c1c589022d0f8310895Db8F2D0f6b",
		"0x9f7Fe8aEf1d6ebB87e25E5B68668b679696eaFbB",
		"0x9F9349c172601D88C620E06e0155bE68698C562A",
		"0x051FD6582C45BF7e8Fd6d8648458480861D5c84f",
		"0x7be2e48d2C814735AaD2dd329a355FaB9c209295",
		"0x40eBe9858a9b5BE3d3Ef67Bb44611Df8DCef1F3c",
		"0xB8689c766c876fd18c7E016A334A9166f64d949A",
		"0x40eF7b510385C0e24672bFFb01A552CBD393167A",
		"0xD216B099167b8a12346eF5b6155A74aEC4D296dA",
		"0xFe6F936B26F1f269f27EFB5Dc64b49B4f915BAe7",
		"0x13e3B73F64b3adE38b1730645041a01F71CF873D",
		"0xAf914e9DC09644e020A732986F361ecdD4e3c6DF",
		"0x650879056d32Ea053b7F3F456466B0E11BEAF2E8",
		"0x0C7c975069e94D8F068bd83d29C79D97a6FFC6D3",
		"0x1bbDe1E811Ca33E0aA5834A40bA4FC0B066EC05E",
		"0xd3e1815369A59Cb5fB6dDA64bCB6a8F4885011fd",
		"0x05378Ad38aD1dd399082e2a04F11980758aaF06F",
		"0xD1BdaeA3fbe4e2194c7CB78Fd1EC0Af7B9ff77DC",
		"0xF16904455727e0513405062afbAF6501e104Ee9b",
		"0x969f1d2112389FE1A7ad9feCeE45715e037EAD7d",
		"0xbA0db8625201a84E773ABdEc7dD0711B790F1e78",
		"0x90DCa53C1a6d176147FaB3E30a7e6436214513F5",
		"0x52a7cD08ec1157b4Af4bEB247Bc95ABb53023998",
		"0xcBf440Fcb5DA3a767B286a746a286fDFBb3bD98c",
		"0xf1C79ADCBe2EC98825ec1eeD7Bf72D4668ef64f0",
		"0x6636b62d9c676ca5a005A55dB918a5B2Ff443711",
		"0xb931ee787ADF24E2e9446d96f92529ef8942aC15",
		"0xdd29E1F550231Ba3F247e2f0567F653e59BD25d1",
		"0xA2fb1a3B295980725529a764A422FB7b7Ad4a8d7",
		"0x4dC23a54B7DAA68B842B45cc6aDe5CC3f141d7C2",
		"0xb3683389045b31160454D30EbECe9552ceB6c002",
		"0x90B6a36513b554caFDD2208EFa69Cc7b0B79d203",
		"0x21A784156F45C034faA085F3da500D7028241c43",
		"0x98E074D1cDab84F04B29d36ba20f91c274C8E1b1",
		"0x6396eD9CEb07FB369d48257bE036fd5f059fB8FA",
		"0xA1050156F6C9e092f21f26ABE02F44CfF2C91465",
		"0x69096890bb79E605f5B7E281915a812e629fEAf2",
		"0x04FB15CD7963b38D1C53DdD0e2Fd1388Bf75bDC0",
		"0x0F0443c5c315B2E878a9e71b58E3BDBEBb400870",
		"0x77699D2C22CaAbAd35dba993b859432650E81e07",
		"0x995f1e0F842aCD5B01639FB2E6Cd0bcaf182dE57",
		"0xE4212BA9186D01F275648CABEdc9a4EdC29a48B9",
		"0x1Da9144298FA79cbaAb55b2ff521f0B011D909F7",
		"0x0e867aF6E3Aa2A0789bc250D5855aDDF320Cf5C6",
		"0xeaDc3d91F18fb34fEBFbCFb5203d05b2D96e52Ef",
		"0xAf5C65DFe532d0C79fe8B1980BA4CecA8D588393",
		"0xa743eea6B3fd8f4bB1861A86C9D99c6AB76E6cd1",
		"0xC12E7861B464eF044E7BddE43894B0D388E9BeaE",
		"0x01800797Bc9303B7c80Caa8E5E881FD8EdaC1D61",
		"0xa4F408CD75F1cce8bc4cc2Fd9689646e69a6FdeD",
		"0x3cC49318CC5304c8CDB158e1D414B452aa7B72Ad",
		"0xF5b3b0bc1FCDB437d4546D40B766521614207b09",
		"0x3F64001FD810Ae8165EA5727980EA733E1d0c378",
		"0x0f1D9687b779e120fDEf5507b20ebECf5335C377",
		"0xCeB6e43bB274286b552f9ACCE7Fb20C82b6b8872",
		"0xa1f81A6Ca3a6E2a942Dc3371b2AcF2BD58834c01",
		"0x2b5a56D7F19e2a9A641E6bE2509aEC89489703C6",
		"0x04FB15CD7963b38D1C53DdD0e2Fd1388Bf75bDC0",
		"0xFCb0208a092F0a106bDdf62f52223E867B87b401",
		"0xCeB6e43bB274286b552f9ACCE7Fb20C82b6b8872",
		"0x6ac3Dc95d4D4Feb9DefD47A0fcb8e472b60a098c",
		"0xaCE87dde9D8bBA8fBf024A0973C17E0654ee3493",
		"0x043f99Af377BC8fB42a8C3431cbdd672d9782b0F",
		"0x3c8653496E3A9eb2FFbAB931Eb5067f126b6a9cd",
		"0xcb8Bc674f11d391d1Bad2F53257016b02BfCC580",
		"0xeabb6C0ec2c8b2C62681e83Bb465397eEA8fA780",
		"0x4B59Bc41FEA567dC7A389e07D2DCAFEc1A589FD5",
		"0x3dd744E3B1123142aA31bC1AEed3059d8A5Be661",
		"0x64D709dCD8793682CEfEBEC58cfD97e3DC7DBe47",
		"0x70367fc3115F5f42CB22A7633c05f830E19fD5f6",
		"0x63EACecF428615eC2Cadc90381e71dE82eA77777",
		"0xf90b156E8c094A2348D66B669B6e79A407F9A067",
		"0xC7791a20d6Ca4b35DD23Ef9c5c97f939FA504642",
		"0xdF4379735Ea99D9fD5b6a4CDEA517a5E6e20E44E",
		"0xc300baBc3cd259bb7fdce31779f32c53F3131B9c",
		"0x1A3b8ad2a44c472AFa6Ef3B32FAe35518a92B7fC",
		"0x49779B21B659BD334b24a2aA94D1902ad028Dce5",
		"0xBB30AF7F53114A35a33b0C5dE91890f4AAbfD117",
		"0x834a4240b6106dB238d90D57aBe88772AC2b80fA",
		"0xba03B1bbdbdAD85922ADBDfBAa870909226d6C34",
		"0xeB4bA061B0daD79e0c17Db259FE81279A84Fb22B",
		"0xDB4dcEBbB5B31341e9Bd52172ad27d675FeE99D1",
		"0x83495590CBEf1Bb7e835f035143C81F1D69DAAD3",
		"0xDf1ce9Bc2E935750E4bc4D6a2FA15C1c13A7032C",
		"0xC5450E7E7c52574fe1D6bA9F554fD0faA0dcA313",
		"0xc9E558f16D219BBE445312D373b36Cce73A03939",
		"0x3F64001FD810Ae8165EA5727980EA733E1d0c378",
		"0x656dD06dfF9C1F04F5dF473F0bE210473Ab8C51c",
		"0xE4212BA9186D01F275648CABEdc9a4EdC29a48B9",
		"0x742B7c268C5378E0cbF1153B81609580c9069F25",
		"0x0329FD90E32dA292eE9F663C17A283Ca11Ca4039",
		"0x30317c1D95bb92cc4D5984Eb695056f883449831",
		"0xCeB6e43bB274286b552f9ACCE7Fb20C82b6b8872",
		"0xdDB4Bf4BeeAa68672A83db1789fAb53c24D84247",
		"0xA3F98dcDdB8f7d14BC733bC202B1bcbD9Bce84B5",
		"0x236dCf7295eAC62e0373f1af8D469ADDA305b878",
		"0x1C23Ba0Bd17123D911002135e9c48beA7A5857E3",
		"0xa62663E5Fb5681882277e0501a6931f067E9AA42",
		"0x80327002b6c35657B8367290B22c2DC69A399b84",
		"0xF66a84D5b8df529eaa96c775f0fB6297D913C594",
		"0x6CE580B0E9081bd60Ba8D11F37F4Eed1796556b8",
		"0x91CDE8AF2ceb804B9e17D6A4FCc45dd2B6976173",
		"0xA9D01fe9C3734543203999b3e4D1FcA08CE9Dff3",
		"0x4695f9e89c26cde1B2c6321d19b5374f3d1b9aa8",
		"0x12d244711A7b3Af55058549f6599a4603C8DEeDA",
		"0xbf8b8784DA646b3122292442411954bF0f8f538b",
		"0x03a508224Dd9e7b1a9660dBfc09c1534B8D9a22A",
		"0xb910189d34B6beF7d1015CF677EA5101b1d6C389",
		"0x47106128f0cE8D7092043e0b564F4D6086209E20",
		"0xA3F98dcDdB8f7d14BC733bC202B1bcbD9Bce84B5",
		"0x29348Fb00EBCdC2a2110efff40CEd988492bA5c0",
		"0x560F75c77eef60542EC8f85F71993a941Cd48d78",
		"0x4155cB3aDbA70DA55A24C002C22242bD25c4F878",
		"0x94dE3bc43f70bA4a1b5dB91248997551Ec88B213",
		"0xe628c97177061168ec9b7f43140f9a84710d0B2d",
		"0x9f98932881E67C003B3a27b1622092d35149c867",
		"0xA9D01fe9C3734543203999b3e4D1FcA08CE9Dff3",
		"0x8d13366291A7B5fA25cDF7A43f94f864e25f8482",
		"0x639fc9A9E4bC3F0BB44f0e79bd0E18DAb22592Fe",
		"0x356e1Aae345b0b8B1dedC5fFD8F9de33f7c22Cd5",
		"0x9008338926a799F3A028cEaB4B0bcaF74aEF0725",
		"0x88509ccD66a55850cF0189e5E9d713c097Ad1f6a",
		"0x7AD34c497814C3660455c336af870B5c32EE71a3",
		"0xBa30fa89Cbb088dB75e37F77c230C944da26738A",
		"0xa84a24E9264068f4adBc1467da250e3672b4c7bF",
		"0xc0C4e04B435185634066C282E3d9b46CE843FA03",
		"0xBE8E58E212BCf934F7297E71989f3D0a1C0BA81A",
		"0x4BaD4C67A9704f6df012ff750d80f9C3832102a3",
		"0xDdce2e9F2427Bf2182Bd6B409Fc5b5B46878aC5E",
		"0x69Dbdd6Df8d96bAaCd0920dD2887071403920b11",
		"0x695ad4733c3A1f7fa860653A81AFb84ef77C5670",
		"0xB9eE36B20d9Fc0c7B61a01Dc499FFcfdCA244da7",
		"0x3f6171A71514A9eB63bFf826F89F8e2F58fc1FAd",
		"0x6071AF68Ce36435B6Ce7BA290E039bf2b42791F9",
		"0x3db83a0912070c4D9eA0ADbE8564eC82283b3f30",
		"0xB3A136831AEb9D327DdfB4b62482cf3d8686b6f8",
		"0x1105b324d8b6fB0aE3960E043980bBbE692d244C",
		"0x285d03Bc3848580BBF712529547749906043C554",
		"0xd38B6F6cC4649029CFAFD6DEe510de7240aa86aA",
		"0x4CF7096c0d9C5420Ab6e5EEdD127BBA7f4FA007E",
		"0x17bB636aCcCADE156C67d88808108387e1f8F1c0",
		"0x409D2022eD5E0c77Be6987d1817512C3BB218888",
		"0x4A8553A16835AfFaB99D4B6172765Ae1171C309A",
		"0x27361dbEB913A810D9A3070043658f694Afb2164",
	}
	for _, address := range addressList {
		moonBoostAddress[address] = true
	}
}

func GetUserInfo(wallet string) (response *model.GetUserInfoResponse, code model.ResponseCode, msg string) {
	user, notFound, err := getCurrentUser(wallet)
	if notFound {
		//return nil, model.NotFound, fmt.Sprintf("No user with address %v", wallet)
		// if not found ,still return all the struct
		user = &model.UserInfo{}
	} else if err != nil {
		return nil, model.ServerInternalError, "Error fetching user info"
	}

	// load owl balance
	amount, err := loadOwlBalanceByWallet(common.HexToAddress(wallet))
	if err != nil {
		//return nil, model.ServerInternalError, "Error fetching owl balance"
		amountDecimal := decimal.New(0, 0)
		amount = &amountDecimal
	}

	// Get Owldinal Nft
	var nftTokens []model.OwldinalNftToken
	result := database.DB.Where("owner = ?", wallet).Find(&nftTokens)
	if result.Error != nil {
		return nil, model.ServerInternalError, fmt.Sprintf("Error fetching NFT tokens (%v)", result.Error)
	}
	owlInfo := model.UserBoxInfo{
		Total: len(nftTokens),
	}
	for _, owldinal := range nftTokens {
		if owldinal.IsStaking {
			owlInfo.Staked += 1
			owlInfo.StakedIdList = append(owlInfo.StakedIdList, owldinal.TokenId)
		} else {
			owlInfo.UnstakedIdList = append(owlInfo.UnstakedIdList, owldinal.TokenId)
		}
	}

	// check moon boost
	isMoonBoost := false
	if config.C.NeedCheckMoonBoost && !notFound {
		isMoonBoost = moonBoostAddress[wallet] || owlInfo.Total > 0
	}

	// get mystery box info & total_earned
	var boxTokens []model.MysteryBoxToken
	if err := database.DB.Where("owner = ?", wallet).Find(&boxTokens).Error; err != nil {
		return nil, model.ServerInternalError, "Error fetching mystery boxes"
	}
	fruitInfo := model.UserBoxInfo{}
	elfInfo := model.UserBoxInfo{}
	for _, token := range boxTokens {
		var boxInfo *model.UserBoxInfo
		if token.BoxType == constant.BoxTypeFruit {
			boxInfo = &fruitInfo
		} else {
			boxInfo = &elfInfo
		}
		boxInfo.Total += 1

		if token.IsStaking {
			boxInfo.Staked += 1
			boxInfo.StakedIdList = append(boxInfo.StakedIdList, token.TokenId)
		} else {
			boxInfo.UnstakedIdList = append(boxInfo.UnstakedIdList, token.TokenId)
		}
	}
	var lastApr model.AprSnapshot
	if err := database.DB.Order("id DESC").First(&lastApr).Error; err != nil {
		return nil, model.ServerInternalError, fmt.Sprintf("failed to load apr: %v", err)
	}
	fruitInfo.Apr = lastApr.FruitApr
	fruitInfo.Apy = lastApr.FruitApy
	elfInfo.Apr = lastApr.ElfApr
	elfInfo.Apy = lastApr.ElfApy

	// calculate referral
	totalReferral := user.ClaimedReferral.Add(user.UnclaimedReferral)
	lockedReferral := user.UnclaimedReferral.Sub(user.UnlockableReferral)
	if lockedReferral.IsNegative() {
		lockedReferral = decimal.Zero
	}
	availableReferral := user.UnclaimedReferral.Sub(lockedReferral)

	response = &model.GetUserInfoResponse{
		Wallet:         wallet,
		HasJoined:      !notFound,
		OwlBalance:     amount.IntPart(),
		TotalEarned:    user.TotalEarned,
		InvitationCode: user.InviteCode,
		InviteCount:    user.InviteCount,
		BuffLevel:      user.BuffLevel,
		IsMoonBoost:    isMoonBoost,
		OwlInfo:        owlInfo,
		ElfInfo:        elfInfo,
		FruitInfo:      fruitInfo,
		ReferralRewards: model.UserReferralRewards{
			Total:     totalReferral.IntPart(),
			Claimed:   user.ClaimedReferral.IntPart(),
			Available: availableReferral.IntPart(),
			Locked:    lockedReferral.IntPart(),
		},
	}

	return response, model.Success, ""
}

func GetUserOwldinalList(wallet string, pagination model.PaginationRequest) (response *model.PaginationResponse[model.UserOwldinal], code model.ResponseCode, msg string) {
	result := &model.PaginationResponse[model.UserOwldinal]{
		Page:    pagination.Page,
		PerPage: pagination.PerPage,
	}

	db := database.DB.Model(&model.OwldinalNftToken{}).Where("owner = ?", wallet)

	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return nil, model.ServerInternalError, fmt.Sprintf("failed to count OwldinalNftToken records: %v", err)
	}

	result.Total = int(total)
	result.PageCount = int((total + int64(pagination.PerPage) - 1) / int64(pagination.PerPage))

	var list []model.OwldinalNftToken
	err = db.Offset((pagination.Page - 1) * pagination.PerPage).
		Limit(pagination.PerPage).
		Find(&list).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			result.List = []model.UserOwldinal{}
		} else {
			return nil, model.ServerInternalError, fmt.Sprintf("failed to find OwldinalNftToken records: %v", err)
		}
	} else {
		result.List = make([]model.UserOwldinal, 0, len(list))

		for _, token := range list {
			data := &model.UserOwldinal{
				TokenId:   token.TokenId,
				TokenUri:  token.TokenUri,
				IsStaking: token.IsStaking,
			}
			result.List = append(result.List, *data)
		}
	}

	return result, model.Success, ""
}

func GetUserMysteryBoxList(wallet string, pagination model.PaginationRequest) (response *model.PaginationResponse[model.UserMysteryBox], code model.ResponseCode, msg string) {
	response = &model.PaginationResponse[model.UserMysteryBox]{
		Page:    pagination.Page,
		PerPage: pagination.PerPage,
	}

	db := database.DB.Model(&model.MysteryBoxToken{}).Where("owner = ?", wallet)

	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return nil, model.ServerInternalError, fmt.Sprintf("failed to count MysteryBoxToken records: %v", err)
	}

	response.Total = int(total)
	response.PageCount = int((total + int64(pagination.PerPage) - 1) / int64(pagination.PerPage))

	var list []model.MysteryBoxToken
	err = db.Offset((pagination.Page - 1) * pagination.PerPage).
		Limit(pagination.PerPage).
		Find(&list).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.List = []model.UserMysteryBox{}
		} else {
			return nil, model.ServerInternalError, fmt.Sprintf("failed to find MysteryBoxToken records: %v", err)
		}
	} else {

		var lastApr model.AprSnapshot
		if err := database.DB.Order("id DESC").First(&lastApr).Error; err != nil {
			return nil, model.ServerInternalError, fmt.Sprintf("failed to load apr: %v", err)
		}

		resultList := make([]model.UserMysteryBox, 0, len(list))
		for _, token := range list {
			data := &model.UserMysteryBox{
				TokenId:   token.TokenId,
				BoxType:   token.BoxType,
				Earning:   token.CurrentRewards,
				Claimed:   token.ClaimedRewards,
				IsStaking: token.IsStaking,
			}

			if token.BoxType == constant.BoxTypeFruit {
				data.Apr = lastApr.FruitApr
				data.Apy = lastApr.FruitApy
			} else if token.BoxType == constant.BoxTypeElf {
				data.Apr = lastApr.ElfApr
				data.Apy = lastApr.ElfApy
			}

			//if token.IsStaking {
			//	// 单个ELF的APR= （单个ELF的日平均Earning/100000）*365
			//	stakingRewards := token.CurrentRewards
			//	stakingDays := int64(time.Now().Sub(*token.StakingTime).Hours()) / 24
			//	if stakingDays == 0 {
			//		stakingDays = 1
			//	}
			//	apr := stakingRewards.
			//		Div(decimal.NewFromInt(stakingDays)).
			//		Div(constant.MysteryBoxMintPrice).
			//		Mul(decimal.NewFromInt32(365))
			//	data.Apr = apr.InexactFloat64()
			//	data.Apy = math.Pow(1+(data.Apr/365), 365) - 1
			//}
			resultList = append(resultList, *data)
		}

		response.List = resultList
	}

	return response, model.Success, ""
}

func GetUserInviteList(wallet string, pagination model.PaginationRequest) (response interface{}, code model.ResponseCode, msg string) {
	result := &model.PaginationResponse[model.InviteRelation]{
		Page:    pagination.Page,
		PerPage: pagination.PerPage,
	}

	db := database.DB.Model(&model.InviteRelation{}).Where("inviter = ?", wallet)

	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return nil, model.ServerInternalError, fmt.Sprintf("failed to count InviteRelation records: %v", err)
	}

	result.Total = int(total)
	result.PageCount = int((total + int64(pagination.PerPage) - 1) / int64(pagination.PerPage))

	var list []model.InviteRelation
	err = db.Offset((pagination.Page - 1) * pagination.PerPage).
		Limit(pagination.PerPage).
		Find(&list).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			result.List = []model.InviteRelation{}
		} else {
			return nil, model.ServerInternalError, fmt.Sprintf("failed to find InviteRelation records: %v", err)
		}
	} else {
		result.List = list
	}

	return result, model.Success, ""
}

func GetMintTx(requestTx string) (response *model.GetMintTxResponse, code model.ResponseCode, msg string) {
	job := model.RequestMintJob{
		RequestTxHash: requestTx,
	}
	if err := database.DB.Where(&job).First(&job).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, model.NotFound, "Can not find this request tx"
		} else {
			return nil, model.ServerInternalError, fmt.Sprintf("failed to get mint tx: %v", err)
		}
	}

	response = &model.GetMintTxResponse{
		RequestTx: requestTx,
		MintTx:    job.JobTxHash,
		JobMsg:    job.Result,
	}

	if job.Status == constant.MintJobStatusFailed && job.RetryCount < 3 {
		response.JobStatus = constant.MintJobStatusProcessing
	} else {
		response.JobStatus = job.Status
	}

	return response, model.Success, ""
}

func getCurrentUser(wallet string) (userInfo *model.UserInfo, notFound bool, err error) {
	userInfoItem := model.UserInfo{
		Address: wallet,
	}
	if err := database.DB.Where(&userInfoItem).First(&userInfoItem).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, true, err
		} else {
			return nil, false, err
		}
	}

	return &userInfoItem, false, nil
}

func loadOwlBalanceByWallet(wallet common.Address) (*decimal.Decimal, error) {
	balance, err := owlTokenContract.BalanceOf(&bind.CallOpts{}, wallet)
	if err != nil {
		return nil, err
	}

	amount := decimal.NewFromBigInt(balance, -18)
	return &amount, nil
}
