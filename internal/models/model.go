package models

import (
	"fmt"
	"github.com/frankie/go-app/global"
	"github.com/frankie/go-app/pkg/setting"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Model struct {
	ID          int    `gorm:"primary_key" json:"id"`
	CreatedBy   string `json:"created_by"`
	CreatedDate string `json:"created_date"`
	UpdatedDate string `json:"updated_date"`
	UpdatedBy   string `json:"updated_by"`
	DeletedDate string `json:"deleted_date"`
	IsDel       uint8  `json:"is_del"`
}

func NewDBEngine(dataBasesetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	db, err := gorm.Open(dataBasesetting.DBType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		dataBasesetting.UserName,
		dataBasesetting.Password,
		dataBasesetting.Host,
		dataBasesetting.DBName,
		dataBasesetting.Charset,
		dataBasesetting.ParseTime,
	))
	if err != nil {
		return nil, err
	}

	if global.ServerSetting.RunMode == "debug" {
		db.LogMode(true)
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(dataBasesetting.MaxIdleConns)
	db.DB().SetMaxOpenConns(dataBasesetting.MaxOpenConns)

	return db, nil
}
