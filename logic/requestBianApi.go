/**
 @author:way
 @date:2021/12/22
 @note
**/

package logic

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"redisData/dao/elasticsearch"
	"redisData/model"
	"redisData/pkg/logger"
	"strings"
)

//Do 封装https 请求
func Do(method string, url string, payload io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return nil, err
	}

	// Set the auth for the request.
	req.SetBasicAuth("admin", "Admin@123")

	return http.DefaultClient.Do(req)
}

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

// GetDataByGraphQL 通过js访问前端服务https://wb.xfack.com/graphql
func GetDataByGraphQL(number int,db string)  {
	//逻辑
	url := "https://wb.xfack.com/graphql"
	jsonData := fmt.Sprintf("{\"query\":\"{\\n  block(number: %d) {\\n    number\\n    hash\\n    parent {\\n      number\\n      transactionCount\\n    }\\n    transactionsRoot\\n    transactionCount\\n    miner {\\n      address\\n      balance\\n      transactionCount\\n      code\\n    }\\n    timestamp\\n    transactions {\\n      hash\\n      from {\\n        address\\n        transactionCount\\n      }\\n      to {\\n        address\\n        transactionCount\\n      }\\n      value\\n      status\\n      createdContract {\\n        address\\n      }\\n      logs {\\n        account {\\n          address\\n        }\\n        topics\\n        data\\n      }\\n    }\\n  }\\n}\\n\",\"variables\":null}",number)
	payload := strings.NewReader(jsonData)
	resp, err := Do("POST",url,payload)
	if err != nil {
		logger.Error(err)
		return
	}
	body,_:=io.ReadAll(resp.Body)
	//反序列化
	var data model.RespGraphData
	err3 := json.Unmarshal(body, &data)
	if err3 != nil {
		fmt.Println(err3)
		return
	}
	//插入ES
	err1 := elasticsearch.CreatBlockData(data.Data,db)
	if err1 != nil {
		fmt.Println(err1)
		return 
	}
	fmt.Println("插入Es成功")

	//client, err := elastic.NewClient(elastic.SetURL("http://10.10.10.8:9200"))
	//if err != nil {
	//	// Handle error
	//	panic(err)
	//}
	//
	//fmt.Println("connect to es success")
	//put1, err := client.Index().
	//	Index("walletMonitor").
	//	BodyJson(data.Data).
	//	Do(context.Background())
	//if err != nil {
	//	// Handle error
	//	panic(err)
	//}
	//fmt.Println(put1.Id,put1.Index,put1.Type)

}

// GetDataByGraphServer 通过js访问前端服务https://wb.xfack.com/graphql
func GetDataByGraphServer(query string) []byte{
	//逻辑
	url := "https://wb.xfack.com/graphql"
	queryStr := query
	fmt.Println(query)
	payload := strings.NewReader(queryStr)

	resp, err := Do("POST",url,payload)

	if err != nil {
		logger.Error(err)
		return nil
	}
	body,_:=io.ReadAll(resp.Body)
	//返回字符串
	return  body
}

