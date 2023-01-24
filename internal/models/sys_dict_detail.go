package models

import "shop/pkg/global"

type SysDictDetail struct {
	Label    string `json:"label" valid:"Required;"`
	Value    string `json:"value" valid:"Required;"`
	Sort     int    `json:"sort"`
	DictId   int64  `json:"dictId"`
	DictName string `json:"dictName"`
	BaseModel
}

func (SysDictDetail) TableName() string {
	return "sys_dict_detail"
}

// get all
func GetAllDictDetail(pageNUm int, pageSize int, maps interface{}) (int64, []SysDictDetail) {
	var (
		total int64
		lists []SysDictDetail
	)
	global.Db.Model(&SysDictDetail{}).Where(maps).Count(&total)
	global.Db.Where(maps).Offset(pageNUm).Limit(pageSize).Find(&lists)

	return total, lists
}

func AddDictDetail(m *SysDictDetail) error {
	var err error
	if err = global.Db.Create(m).Error; err != nil {
		return err
	}

	return err
}

func UpdateByDictDetail(m *SysDictDetail) error {
	var err error
	err = global.Db.Save(m).Error
	if err != nil {
		return err
	}

	return err
}

func DelByDictDetail(ids []int64) error {
	var err error
	err = global.Db.Model(&SysDictDetail{}).Where("id in (?)", ids).Update("is_del", 1).Error
	if err != nil {
		return err
	}

	return err
}
