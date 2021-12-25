/**
 @author:way
 @date:2021/12/25
 @note
**/

package main

import (
	"encoding/json"
	"fmt"
	"redisData/dao/redis"
	"redisData/logic"
	"redisData/model"
	"redisData/utils"
)

func main() {


	//获取最新的区块
	query := "{\"query\":\"{\\n  block{number}\\n}\",\"variables\":null}"
	respData := logic.GetDataByGraphServer(query)
	//反序列化
	var data model.RespBlockNumber
	err := json.Unmarshal(respData, &data)
	if err != nil {
		fmt.Println(err)
		return
	}


	//获取redis里面的区块号
	value, err := redis.GetDataByKey("newestBlockNumber")
	if err != nil {
		fmt.Println(err)
		return
	}
	blockNum := utils.StringToInt(value)
	if blockNum == data.Data.Block.Number{
		return
	}


}