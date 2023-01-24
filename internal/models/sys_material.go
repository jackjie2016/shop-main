package models

import "shop/pkg/global"

type SysMaterial struct {
	Name     string `json:"name" valid:"Required;"`
	Type     string `json:"type"`
	Url      string `json:"url"`
	GroupId  int64  `json:"groupId"`
	CreateId int64  `json:"create_id"`
	BaseModel
}

func (SysMaterial) TableName() string {
	return "sys_material"
}

func GetAllMaterial(pageNUm int, pageSize int, maps interface{}) (int64, []SysMaterial) {
	var (
		total int64
		data  []SysMaterial
	)
	global.Db.Model(&SysMaterial{}).Where(maps).Count(&total)
	global.Db.Where(maps).Offset(pageNUm).Limit(pageSize).Order("id DESC").Find(&data)

	return total, data
}

func AddMaterial(m *SysMaterial) error {
	var err error
	if err = global.Db.Create(m).Error; err != nil {
		return err
	}

	return err
}

func UpdateByMaterial(m *SysMaterial) error {
	var err error
	err = global.Db.Save(m).Error
	if err != nil {
		return err
	}

	return err
}

func DelByMaterial(ids []int64) error {
	var err error
	err = global.Db.Where("id in (?)", ids).Delete(&SysMaterial{}).Error
	if err != nil {
		return err
	}

	return err
}
