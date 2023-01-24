package models

import "shop/pkg/global"

type SysLog struct {
	Description string `json:"description"`
	Method      string `json:"method"`
	RequestIp   string `json:"requestIp"`
	Username    string `json:"username"`
	Address     string `json:"address"`
	Browser     string `json:"browser"`
	Type        int8   `json:"type"`
	Uid         int64  `json:"uid"`
	Time        int64  `json:"time"`
	BaseModel
}

func (SysLog) TableName() string {
	return "sys_log"
}

// get all
func GetAllLog(pageNUm int, pageSize int, maps interface{}) (int64, []SysLog) {
	var (
		total int64
		lists []SysLog
	)

	global.Db.Model(&SysLog{}).Where(maps).Count(&total)
	global.Db.Where(maps).Offset(pageNUm).Limit(pageSize).Find(&lists)

	return total, lists
}

// last inserted Id on success.
func AddLog(m *SysLog) error {
	if err := global.Db.Create(m).Error; err != nil {
		return err
	}
	return nil
}

func DelBylog(ids []int64) error {
	var err error
	err = global.Db.Where("id in (?)", ids).Delete(&SysLog{}).Error
	if err != nil {
		return err
	}

	return err
}
