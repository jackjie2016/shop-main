package models

import "shop/pkg/global"

type StoreCanvas struct {
	Terminal int    `json:"terminal"`
	Json     string `json:"json"`
	Ttype    int    `json:"type" gorm:"column:type"`
	Name     string `json:"name"`
	BaseModel
}

func (StoreCanvas) TableName() string {
	return "store_canvas"
}

func AddCanvas(m *StoreCanvas) error {
	var err error
	if err = global.Db.Create(m).Error; err != nil {
		return err
	}

	return err
}

func UpdateByCanvas(m *StoreCanvas) error {
	var err error
	err = global.Db.Save(m).Error
	if err != nil {
		return err
	}

	return err
}

func DelByCanvas(ids []int64) error {
	var err error
	err = global.Db.Where("id in (?)", ids).Delete(&StoreCanvas{}).Error
	if err != nil {
		return err
	}

	return err
}
