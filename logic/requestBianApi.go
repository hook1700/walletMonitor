/**
 @author:way
 @date:2021/12/22
 @note
**/

package logic

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"redisData/pkg/logger"
)

//GetTxListByAddress 通过合约或者钱包地址，返回交易列表
func GetTxListByAddress(address string, apikey string, sort string, offset string, page string) []byte {
	//校验参数
	if len(address) <= 0 {
		logger.Error(errors.New("输入地址为空"))
		return nil
	}
	if len(apikey) <= 0 {
		apikey = "DUKNN1QZMITSSZC61YINTD1CWQ92FWEKHM"
	}
	if len(sort) <= 0 {
		sort = "desc"
	}
	if len(offset) <= 0 {
		offset = "1000"
	}
	if len(page) <= 0 {
		page = "1"
	}
	//逻辑
	url := fmt.Sprintf("https://api.bscscan.com/api?module=account&action=txlist&address=%s&startblock=0&endblock=99999999&page=%s&offset=%s&sort=%s&apikey=%s",address,page,offset,sort,apikey)
	logger.Info(url)
	response, GErr := http.Get(url) //这步访问可能会慢
	if GErr != nil {
		logger.Error(GErr)
		return nil
	}
	//反序列化
	body, ReadAllErr := ioutil.ReadAll(response.Body)
	if ReadAllErr != nil{
		logger.Error(ReadAllErr)
		return nil
	}
	return body
}

//MonitorAddressAndContract 币安监控钱包接口,通过地址和合约地址监控钱包下某合约的交易
func MonitorAddressAndContract(contractAddress string,address string, apikey string, sort string, offset string, page string) []byte  {
	//校验参数
	if len(address) <= 0 {
		logger.Error(errors.New("输入合约地址为空"))
		return nil
	}
	if len(address) <= 0 {
		logger.Error(errors.New("输入钱包地址为空"))
		return nil
	}
	if len(apikey) <= 0 {
		apikey = "DUKNN1QZMITSSZC61YINTD1CWQ92FWEKHM"
	}
	if len(sort) <= 0 {
		sort = "desc"
	}
	if len(offset) <= 0 {
		offset = "1000"
	}
	if len(page) <= 0 {
		page = "1"
	}
	//逻辑
	url := fmt.Sprintf("https://api.bscscan.com/api?module=account&action=tokentx&contractaddress=%s&address=%s&startblock=0&endblock=99999999&page=%s&offset=%s&sort=%s&apikey=%s",contractAddress,address,page,offset,sort,apikey)
	logger.Info(url)
	response, GErr := http.Get(url) //这步访问可能会慢
	if GErr != nil {
		logger.Error(GErr)
		return nil
	}
	//反序列化
	body, ReadAllErr := ioutil.ReadAll(response.Body)
	if ReadAllErr != nil{
		logger.Error(ReadAllErr)
		return nil
	}
	return body
}