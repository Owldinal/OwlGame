package database

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
		Logger:         logger.Default.LogMode(logger.Silent),
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

//goland:noinspection GoMixedReceiverTypes
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

//goland:noinspection GoMixedReceiverTypes
func (idList IdList) Value() (driver.Value, error) {
	if idList == nil {
		return "", nil
	}

	strValues := make([]string, len(idList))
	for i, intVal := range idList {
		strValues[i] = fmt.Sprintf("%d", intVal)
	}
	return strings.Join(strValues, ","), nil
}
