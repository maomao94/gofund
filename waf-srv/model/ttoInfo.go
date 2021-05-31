// 自动生成模板TtoInfo
package model

import (
	"github.com/hehanpeng/gofund/common/global"
	"time"
)

// 如果含有time.Time 请自行import time包
type TtoInfo struct {
	global.GVA_MODEL
	Reference    int       `json:"reference" form:"reference" gorm:"column:reference;comment:引用;type:bigint;size:19;"`
	RegisterTime time.Time `json:"registerTime" form:"registerTime" gorm:"column:register_time;comment:注册时间;type:timestamp;"`
	TtoType      string    `json:"ttoType" form:"ttoType" gorm:"column:tto_type;comment:类型（全局/局部）0:全局 1：局部;type:varchar(1);size:1;"`
	TtoStatus    string    `json:"ttoStatus" form:"ttoStatus" gorm:"column:tto_status;comment:STATUS状态 0：未执行 1：已执行;type:varchar(1);size:1;"`
	BizType      string    `json:"bizType" form:"bizType" gorm:"column:biz_type;comment:业务类型;type:varchar(64);size:64;"`
	CallIp       string    `json:"callIp" form:"callIp" gorm:"column:call_ip;comment:回调ip;type:varchar(64);size:64;"`
	CallMethod   string    `json:"callMethod" form:"callMethod" gorm:"column:call_method;comment:回调方法;type:varchar(128);size:128;"`
	ExecuteTime  time.Time `json:"executeTime" form:"executeTime" gorm:"column:execute_time;comment:执行时间;type:timestamp;"`
	ExpiredTime  int       `json:"expiredTime" form:"expiredTime" gorm:"column:expired_time;comment:过期时间,单位毫秒;type:bigint;size:19;"`
	Ext1         string    `json:"ext1" form:"ext1" gorm:"column:ext1;comment:备用1;type:varchar(128);size:128;"`
}

func (TtoInfo) TableName() string {
	return "tto_info"
}
