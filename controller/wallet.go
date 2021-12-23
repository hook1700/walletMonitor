/**
 @author:way
 @date:2021/12/16
 @note
**/

package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"redisData/dao/redis"
	"redisData/logic"
	"redisData/model"
	"redisData/pkg/logger"
)

type Controller struct {
}

//GetTxHashListByAddress  获取交易hash列表
func (h *Controller) GetTxHashListByAddress(c *gin.Context) {
	var(
		p model.ParamAddress
		resp []string
	)
	//接收参数
	err1 := c.Bind(&p)
	if err1 != nil {
		logger.Error(err1)
		ResponseError(c,501)
		return
	}
	//逻辑处理1.访问币安api  2.把交易hash存redis,同时返回
	for {
		i := 1
		byteData := logic.GetTxListByAddress(p.Address, "", "", "", fmt.Sprintf("%d",i))
		if byteData == nil{
			logger.Error(errors.New("请求币安api返回数据为空"))
			ResponseError(c,500)
			return
		}
		//反序列化
		var data model.RespTxList
		err2 := json.Unmarshal(byteData,&data)
		if err2 != nil {
			logger.Error(err2)
			ResponseError(c,500)
			return
		}
		//获取切片中的tx hash
		for _,v := range data.Result{
			resp = append(resp,v.Hash)
		}
		if len(data.Result) == 1000{
			i++
		}else {
			break
		}
	}

	//返回数据
	//resp切片转json
	respByte, err4 := json.Marshal(resp)
	if err4 != nil {
		logger.Error(err4)
		ResponseError(c,500)
		return
	}
	err3 := redis.CreateDurableKey(fmt.Sprintf("txList:%s",p.Address),string(respByte))
	if err3 != nil {
		logger.Error(errors.New("redis添加数据失败"))
		return
	}
	ResponseSuccess(c, resp)
}

//MonitorWallet 监控钱包地址和和合约的交集
func (h *Controller) MonitorWallet(c *gin.Context) {
	var(
		p model.ParamAddressAndContract
		resp []string
	)
	//接收参数
	err1 := c.Bind(&p)
	if err1 != nil {
		logger.Error(err1)
		ResponseError(c,501)
		return
	}
	//逻辑处理1.访问币安api  2.把交易hash存redis,同时返回
	for {
		i := 1
		byteData := logic.MonitorAddressAndContract(p.ContractAddress,p.Address, "", "", "", fmt.Sprintf("%d",i))
		if byteData == nil{
			logger.Error(errors.New("请求币安api返回数据为空"))
			ResponseError(c,500)
			return
		}
		//反序列化
		var data model.RespMonitorList
		err2 := json.Unmarshal(byteData,&data)
		if err2 != nil {
			logger.Error(err2)
			ResponseError(c,500)
			return
		}
		//获取切片中的tx hash
		for _,v := range data.Result{
			resp = append(resp,v.Hash)
		}
		if len(data.Result) == 1000{
			i++
		}else {
			break
		}
	}
	//byteData := logic.MonitorAddressAndContract(p.ContractAddress,p.Address, "", "", "", "1")
	//if byteData == nil{
	//	logger.Error(errors.New("请求币安api返回数据为空"))
	//	ResponseError(c,500)
	//	return
	//}
	////反序列化
	//var data model.RespMonitorList
	//err2 := json.Unmarshal(byteData,&data)
	//if err2 != nil {
	//	logger.Error(err2)
	//	ResponseError(c,500)
	//	return
	//}
	////获取切片中的tx hash
	//
	////判断
	//for _,v := range data.Result{
	//	resp = append(resp,v.Hash)
	//}
	//resp切片转json
	respByte, err4 := json.Marshal(resp)
	if err4 != nil {
		logger.Error(err4)
		ResponseError(c,500)
		return
	}
	err3 := redis.CreateDurableKey(fmt.Sprintf("ERC20Monitor:%s:%s",p.ContractAddress,p.Address),string(respByte))
	if err3 != nil {
		logger.Error(errors.New("redis添加数据失败"))
		return
	}
	//返回数据
	ResponseSuccess(c, resp)
}



