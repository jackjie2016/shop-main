package models

import (
	"shop/pkg/global"
	"time"
)

type UserExtract struct {
	Uid          int64     `json:"uid"`
	RealName     string    `json:"real_name"`
	ExtractType  string    `json:"extract_type"`
	BankCode     string    `json:"bank_code"`
	BankAddress  string    `json:"bank_address"`
	AlipayCode   string    `json:"alipay_code"`
	ExtractPrice float64   `json:"extract_price"`
	Mark         string    `json:"mark"`
	Balance      float64   `json:"balance"`
	FailMsg      string    `json:"fail_msg"`
	FailTime     time.Time `json:"fail_time"`
	Status       int8      `json:"status"`
	Wechat       string    `json:"wechat"`
	BaseModel
}

func (UserExtract) TableName() string {
	return "user_extract"
}

func AddUserExtract(m *UserExtract) error {
	var err error
	if err = global.Db.Create(m).Error; err != nil {
		return err
	}

	return err
}

func UpdateByUserExtract(m *UserExtract) error {
	var err error
	err = global.Db.Save(m).Error
	if err != nil {
		return err
	}

	return err
}

func DelByUserExtract(ids []int64) error {
	var err error
	err = global.Db.Where("id in (?)", ids).Delete(&UserExtract{}).Error
	if err != nil {
		return err
	}

	return err
}
