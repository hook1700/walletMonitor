/**
 @author:way
 @date:2021/12/24
 @note
**/

package main

import (
	"encoding/json"
	"fmt"
	"redisData/dao/elasticsearch"
	"redisData/dao/redis"
	"redisData/logic"
	"redisData/model"
	"redisData/pkg/logger"
	"redisData/setting"
	"redisData/utils"
	"time"
)
func init() {
	// 定义日志目录
	logger.Init("buy")
	// 初始化 viper 配置
	if err := setting.Init(""); err != nil {
		logger.Info("viper init fail")
		logger.Error(err)
		return
	}
	//初始化redis
	if err := redis.InitClient(); err != nil {
		logger.Info("init redis fail err")
		logger.Error(err)
		return
	}

	elasticsearch.InitES()
}

func startSyncBlockData()  {
	fmt.Println("开始同步数据")
	//通过网络请求获取区块号
	query := "{\"query\":\"{\\n  block{number}\\n}\",\"variables\":null}"
	respData := logic.GetDataByGraphServer(query)
	//反序列化
	var data model.RespBlockNumber
	err := json.Unmarshal(respData, &data)
	if err != nil {
		fmt.Println(err)
		return
	}
	//拿到线上访问获取的区块号
	blockNumber := data.Data.Block.Number
	//存一份到redis
	err2 := redis.CreateDurableKey("newestBlockNumber", blockNumber)
	if err2 != nil {
		fmt.Println(err2)
		return 
	}

	for i :=blockNumber;i > (blockNumber-2);i--{
		//db 为ES的db
		logic.GetDataByGraphQL(i,"wallet")
	}
}

func startSyncNewBlock()  {
	fmt.Println("开始监控最新区块，出现最新区块即将同步")
	//先获取最新的区块号
	//获取最新的区块
	for {
		query := "{\"query\":\"{\\n  block{number}\\n}\",\"variables\":null}"
		respData := logic.GetDataByGraphServer(query)
		//反序列化
		var data model.RespBlockNumber
		err := json.Unmarshal(respData, &data)
		if err != nil {
			fmt.Println(err)
			return
		}
		//拿到线上访问获取的区块号
		blockNumber := data.Data.Block.Number

		//获取redis里面的区块号
		value, err := redis.GetDataByKey("newestBlockNumber")
		if err != nil {
			fmt.Println(err)
			return
		}
		blockNum := utils.StringToInt(value)
		if blockNum != blockNumber{
			fmt.Println("发现区块更新，且同步数据")
			//就把区块的数据同步下来
			logic.GetDataByGraphQL(blockNum,"wallet")
			//每次同步完成得更新redis数据
			err2 := redis.CreateDurableKey("newestBlockNumber", blockNumber)
			if err2 != nil {
				fmt.Println(err2)
				return
			}
		}
		time.Sleep(time.Second*3)
	}
}

func main() {
	go startSyncBlockData()
	time.Sleep(5*time.Second)
	go startSyncNewBlock()
	//阻塞线程
	select {}
}
