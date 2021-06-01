package service

import (
	"context"
	"errors"
	"waf-srv/model"
	"waf-srv/pkg/invoker"
	"waf-srv/request"

	"github.com/hehanpeng/gofund/common/resp"

	"github.com/gotomicro/ego/core/etrace"

	"github.com/hehanpeng/gofund/common/req"
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

// 执行超时转发逻辑
func DealCronTtoInfo(ctx context.Context, ttoInfo model.TtoInfo) error {
	//TtoInf record = new TtoInf();
	//record.setTtoid(input.getTtoid());
	//record.setTtoStatus(ForwardParamConstant.TTOINF_STATUS_DEALED);
	//Map<String, String> reqMap = new HashMap<>();
	//reqMap.put(TtoInf.TTOID, StringUtilF.parseObjectToString(input.getTtoid()));
	//reqMap.put(TtoInf.REFERENCE, StringUtilF.parseObjectToString(input.getReference()));
	//reqMap.put(TtoInf.BIZ_TYPE, StringUtilF.parseObjectToString(input.getBizType()));
	//reqMap.put(TtoInf.TTO_TYPE, StringUtilF.parseObjectToString(input.getTtoType()));
	//reqMap.put(TtoInf.CALL_CLASS, StringUtilF.parseObjectToString(input.getCallClass()));
	//reqMap.put(TtoInf.CALL_METHOD, StringUtilF.parseObjectToString(input.getCallMethod()));
	//reqMap.put(TtoInf.EXT1, StringUtilF.parseObjectToString(input.getExt1()));
	//reqMap.put(TtoInf.EXT2, StringUtilF.parseObjectToString(input.getExt2()));
	//reqMap.put(TtoInf.EXT3, StringUtilF.parseObjectToString(input.getExt3()));
	//reqMap.put(TtoInf.EXT4, StringUtilF.parseObjectToString(input.getExt4()));
	//reqMap.put(TtoInf.EXT5, StringUtilF.parseObjectToString(input.getExt5()));
	//try {
	//	//grpc 调用
	//	Class<?> cl = Class.forName(input.getCallClass());
	//	SpringUtil.springInvokeMethod(cl, input.getCallMethod(), new Object[]{reqMap});
	//	ttoInfMapper.updateByPrimaryKeySelective(record);
	//} catch (Throwable e) {
	//	log.error("tto rpc error {}调用失败", input.getCallClass());
	//}
	// http 调用
	callSrvHttpComp := invoker.CallSrvHttpComps[ttoInfo.CallSrvName]
	if callSrvHttpComp == nil {
		invoker.Logger.Errorf("callSrvHttpComps[v%] is not register", ttoInfo.CallSrvName)
		return errors.New("callSrvHttpComp is nil")
	}
	req := callSrvHttpComp.R()
	// Inject traceId Into Header
	c1 := etrace.HeaderInjector(ctx, req.Header)
	info, err := req.SetContext(c1).
		SetBody(ttoInfo).
		SetResult(resp.Resp{}).
		Post(ttoInfo.CallMethod)
	if err != nil {
		invoker.Logger.Error("callSrvHttpComp post error")
		return errors.New("callSrvHttpComp post error")
	}
	invoker.Logger.Infof("result Success: %v", info.Result().(resp.Resp))
	return nil
}
