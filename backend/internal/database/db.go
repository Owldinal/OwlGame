package database

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"math/big"
	"owl-backend/internal/config"
	"strings"
	"time"

	"gorm.io/driver/mysql"
)

var (
	DB *gorm.DB
)

func init() {
	var err error
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=UTC",
		config.C.DBUser,
		config.C.DBPassword,
		config.C.DBHost,
		config.C.DBPort,
		config.C.DBName,
	)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		TranslateError: true,
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v mysql init successfully\n", time.Now())
}

type Model struct {
	ID        uint           `gorm:"primarykey" json:"-"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type IdList []uint64

func (idList *IdList) Scan(value interface{}) error {
	if value == nil {
		*idList = make(IdList, 0)
		return nil
	}

	val, ok := value.(string)
	if !ok {
		return errors.New("buffingIds should be a string")
	}

	parts := strings.Split(val, ",")
	*idList = make(IdList, 0, len(parts))
	for _, part := range parts {
		var intVal uint64
		if part != "" {
			if err := json.Unmarshal([]byte(part), &intVal); err != nil {
				return err
			}
			*idList = append(*idList, intVal)
		}
	}
	return nil
}

func (idList *IdList) Value() (driver.Value, error) {
	if idList == nil {
		return "", nil
	}

	// 将整数切片转换为逗号分隔的字符串
	strValues := make([]string, len(*idList))
	for i, intVal := range *idList {
		strValues[i] = fmt.Sprintf("%d", intVal)
	}
	return strings.Join(strValues, ","), nil
}

type Amount struct {
	*big.Int
}

func NewAmount(value string) Amount {
	number, _ := big.NewInt(0).SetString(value, 10)
	return Amount{number}
}

func (a *Amount) Scan(value interface{}) error {
	switch v := value.(type) {
	case int64:
		a.Int = big.NewInt(v)
	case []byte:
		a.Int = new(big.Int)
		_, success := a.Int.SetString(string(v), 10)
		if !success {
			return errors.New("Amount: Scan: invalid input")
		}
	case string:
		a.Int = new(big.Int)
		_, success := a.Int.SetString(v, 10)
		if !success {
			return errors.New("Amount: Scan: invalid input")
		}
	case nil:
		a.Int = nil
	default:
		return errors.New("Amount: Scan: unsupported data type")
	}
	return nil
}

func (a Amount) Value() (driver.Value, error) {
	if a.Int == nil {
		return nil, nil
	}
	return a.Int.String(), nil
}
