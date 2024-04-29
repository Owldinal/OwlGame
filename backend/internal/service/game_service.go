package service

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"owl-backend/internal/constant"
	"owl-backend/internal/database"
	"owl-backend/internal/model"
	"strconv"
	"time"
)

func GetGameInfo() (response *model.GetGameInfoResponse, code model.ResponseCode, msg string) {
	var snapshots []model.DailyPoolSnapshot
	// Get newest data
	if err := database.DB.Order("id desc").Limit(2).Find(&snapshots).Error; err != nil {
		return nil, model.ServerInternalError, fmt.Sprintf("Internal Error (%v)", err)
	}
	if len(snapshots) == 0 {
		return nil, model.ServerInternalError, fmt.Sprintf("Failed fetch game snpashot")
	}

	currentDate := time.Now().UTC().Truncate(24 * time.Hour)
	var today, yesterday *model.DailyPoolSnapshot
	if snapshots[0].Date != currentDate {
		// The snapshot for today has not been generated;
		// use yesterday's data and all the change rates are zero
		today = &snapshots[0]
		yesterday = &snapshots[0]
	} else {
		today = &snapshots[0]
		yesterday = &snapshots[1]
	}

	response = &model.GetGameInfoResponse{
		TotalRewards:   today.TotalPoolAmount,
		TotalMarketCap: decimal.NewFromInt(2100000000).Sub(today.TotalBurn),
		TotalBurned:    today.TotalBurn,
	}

	yesterdayMarketCap := decimal.NewFromInt(2100000000).Sub(yesterday.TotalBurn)
	if yesterdayMarketCap.IsZero() {
		response.TotalMarketCapChange = decimal.NewFromInt(0)
	} else {
		response.TotalMarketCapChange = response.TotalMarketCap.Sub(yesterdayMarketCap).Div(yesterdayMarketCap)
	}

	if yesterday.TotalBurn.IsZero() {
		response.TotalBurnedChange = decimal.NewFromInt(1)
	} else {
		response.TotalBurnedChange = today.TotalBurn.Sub(yesterday.TotalBurn).Div(yesterday.TotalBurn)
	}

	// Frontend will load this data from https://docs.dexscreener.com/api/reference
	//owlToUsdMultiple := 0.001
	//response.TotalRewardUSD = response.TotalRewards.Mul(decimal.NewFromFloat(owlToUsdMultiple))
	//
	//response.OwlPrice = decimal.NewFromFloat(0.001)
	//response.OwlPriceChange = decimal.NewFromFloat(0)

	type Result struct {
		BoxType constant.BoxType
		Count   uint64
	}
	var results []Result
	if err := database.DB.Model(&model.MysteryBoxToken{}).
		Select("box_type, count(*) as count").
		Where("is_staking = ?", true).
		Group("box_type").
		Find(&results).Error; err != nil {
		return nil, model.ServerInternalError, fmt.Sprintf("Internal Error (%v)", err)
	}
	for _, count := range results {
		if count.BoxType == constant.BoxTypeElf {
			response.StakedElfCount = count.Count
		} else if count.BoxType == constant.BoxTypeFruit {
			response.StakedFruitCount = count.Count
		}
	}

	return response, model.Success, ""
}

func GetRewardsTrending() (response *model.GetGameRewardsTrendingResponse, code model.ResponseCode, msg string) {
	//now := time.Now()

	var snapshots model.DailyPoolSnapshot
	if err := database.DB.Model(&model.DailyPoolSnapshot{}).Order("id desc").First(&snapshots).Error; err != nil {
		return response, model.ServerInternalError, "Failed to load snapshot"
	}
	now := snapshots.Date

	dailyDataPoints, dailyFrom, dailyTo, err1 := generateRewardsTrending(now, 24*time.Hour, 12)
	weeklyDataPoints, weeklyFrom, weeklyTo, err2 := generateRewardsTrending(now, 7*24*time.Hour, 12)
	monthlyDataPoints, monthlyFrom, monthlyTo, err3 := generateRewardsTrending(now, 30*24*time.Hour, 12)
	if err1 != nil || err2 != nil || err3 != nil {
		return response, model.ServerInternalError, "Failed to load snapshot"
	}

	response = &model.GetGameRewardsTrendingResponse{
		Daily: &model.RewardTrendingDetail{
			From:       dailyFrom,
			To:         dailyTo,
			DataPoints: dailyDataPoints,
		},
		Weekly: &model.RewardTrendingDetail{
			From:       weeklyFrom,
			To:         weeklyTo,
			DataPoints: weeklyDataPoints,
		},
		Monthly: &model.RewardTrendingDetail{
			From:       monthlyFrom,
			To:         monthlyTo,
			DataPoints: monthlyDataPoints,
		},
	}

	return response, model.Success, ""
}

func generateRewardsTrending(start time.Time, interval time.Duration, count int) ([]model.DataPoint, *time.Time, *time.Time, error) {
	var snapshots []model.DailyPoolSnapshot
	dataStrings := make([]string, count)
	for i := 0; i < count; i++ {
		date := start.Add(-interval * time.Duration(i))
		dataStrings[i] = date.Format("2006-01-02")
	}
	if err := database.DB.Where("date IN ?", dataStrings).Find(&snapshots).Error; err != nil {
		return nil, nil, nil, err
	}

	var dataPoints []model.DataPoint
	for _, snapshot := range snapshots {
		dataPoints = append(dataPoints, model.DataPoint{
			Date:             snapshot.Date,
			TotalPoolAmount:  snapshot.TotalPoolAmount,
			AllocatedRewards: snapshot.AllocatedRewards,
		})
	}

	var actualFrom, actualTo time.Time
	if len(dataPoints) > 0 {
		actualFrom = dataPoints[0].Date
		actualTo = dataPoints[len(dataPoints)-1].Date

		return dataPoints, &actualFrom, &actualTo, nil
	} else {
		return dataPoints, nil, nil, nil
	}
}

func GetTreasuryRevenueHistory(cursor model.RewardsHistoryRequest) (response *model.CursorResponse[model.TreasuryRevenueHistory], code model.ResponseCode, msg string) {
	var records []model.RewardPoolTransactionRecord
	query := database.DB.Model(&model.RewardPoolTransactionRecord{}).Order("id DESC")
	if cursor.Cursor > 0 {
		query = query.Where("id < ?", cursor.Cursor)
	}
	if len(cursor.Address) > 0 {
		query = query.Where("user = ?", cursor.Address)
	}
	var totalCount int64
	query.Count(&totalCount)
	if err := query.Limit(cursor.Limit).Find(&records).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, model.ServerInternalError, fmt.Sprintf("Failed fetch reward transaction. (%v)", err)
		}
	}

	var result []model.TreasuryRevenueHistory
	for _, record := range records {
		history := model.TreasuryRevenueHistory{
			Address:         record.User,
			Operation:       record.Operation,
			Description:     record.Description,
			Count:           int64(record.Count),
			Amount:          record.Amount,
			TransactionHash: record.TxHash,
		}
		result = append(result, history)
	}

	nextCursor := 0
	if len(records) > 0 {
		nextCursor = int(records[len(records)-1].ID)
	}
	response = &model.CursorResponse[model.TreasuryRevenueHistory]{
		Cursor:     cursor.Cursor,
		Limit:      cursor.Limit,
		NextCursor: nextCursor,
		HasMore:    totalCount > int64(cursor.Limit),
		List:       result,
	}

	return response, model.Success, ""
}

func GetRequestJobStatus(requestTxHash string) (response *model.RequestJobResponse, code model.ResponseCode, msg string) {
	job := model.RequestMintJob{
		RequestTxHash: requestTxHash,
	}

	if err := database.DB.Where(&job).First(&job).Error; err != nil {
		return nil, model.NotFound, fmt.Sprintf("failed to find job with tx %v. err: %v", requestTxHash, err)
	}

	response = &model.RequestJobResponse{
		RequestTx:  requestTxHash,
		JobTx:      job.JobTxHash,
		IsSuccess:  job.HasConfirmed,
		LastUpdate: &job.UpdatedAt,
	}
	if response.JobTx != "" {
		response.JobUrl = fmt.Sprintf("https://scan.merlinchain.io/tx/%v", response.JobTx)
	}
	return response, model.Success, ""
}

func CheckSignature(request model.CheckSignatureRequest) (code model.ResponseCode, msg string) {

	//fmt.Printf("request: %+v\n", request)

	hashedMessage := []byte("\x19Ethereum Signed Message:\n" + strconv.Itoa(len(request.Message)) + request.Message)
	hash := crypto.Keccak256Hash(hashedMessage)
	//fmt.Printf("hash: %x\n", hash)

	decodedMessage := hexutil.MustDecode(request.Signature)
	if decodedMessage[64] == 27 || decodedMessage[64] == 28 {
		decodedMessage[64] -= 27
	}

	sigPublicKeyECDSA, err := crypto.SigToPub(hash.Bytes(), decodedMessage)
	if sigPublicKeyECDSA == nil {
		err = errors.New("could not get a public get from the message signature")
	}
	if err != nil {
		return model.ServerInternalError, err.Error()
	}

	publicKey := crypto.PubkeyToAddress(*sigPublicKeyECDSA).String()
	if publicKey != request.Address {
		return model.Success, "false"
	}

	return model.Success, "true"

}
