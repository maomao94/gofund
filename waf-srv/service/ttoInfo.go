package service

import (
	"github.com/hehanpeng/gofund/common/req"
	"waf-srv/model"
	"waf-srv/pkg/invoker"
	"waf-srv/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateTtoInfo
//@description: 创建TtoInfo记录
//@param: ttoInfo model.TtoInfo
//@return: err error

func CreateTtoInfo(ttoInfo model.TtoInfo) (err error) {
	err = invoker.Db.Create(&ttoInfo).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteTtoInfo
//@description: 删除TtoInfo记录
//@param: ttoInfo model.TtoInfo
//@return: err error

func DeleteTtoInfo(ttoInfo model.TtoInfo) (err error) {
	err = invoker.Db.Delete(&ttoInfo).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteTtoInfoByIds
//@description: 批量删除TtoInfo记录
//@param: ids req.IdsReq
//@return: err error

func DeleteTtoInfoByIds(ids req.IdsReq) (err error) {
	err = invoker.Db.Delete(&[]model.TtoInfo{}, "id in ?", ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateTtoInfo
//@description: 更新TtoInfo记录
//@param: ttoInfo *model.TtoInfo
//@return: err error

func UpdateTtoInfo(ttoInfo model.TtoInfo) (err error) {
	err = invoker.Db.Save(&ttoInfo).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetTtoInfo
//@description: 根据id获取TtoInfo记录
//@param: id uint
//@return: err error, ttoInfo model.TtoInfo

func GetTtoInfo(id uint) (err error, ttoInfo model.TtoInfo) {
	err = invoker.Db.Where("id = ?", id).First(&ttoInfo).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetTtoInfoInfoList
//@description: 分页获取TtoInfo记录
//@param: info req.TtoInfoSearch
//@return: err error, list interface{}, total int64

func GetTtoInfoInfoList(info request.TtoInfoSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := invoker.Db.Model(&model.TtoInfo{})
	var ttoInfos []model.TtoInfo
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&ttoInfos).Error
	return err, ttoInfos, total
}
