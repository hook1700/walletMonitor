/**
 @author:way
 @date:2021/12/23
 @note
**/

package elasticsearch


import (
	"context"
	"github.com/olivere/elastic/v7"
	"github.com/spf13/viper"
	"redisData/pkg/logger"
)
var Client *elastic.Client
//InitES ES初始化
func init()  {
	host := viper.GetString("elasticsearch.host")
	//这个地方有个小坑 不加上elastic.SetSniff(false) 会连接不上
	Client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(host))
	if err != nil {
		logger.Error(err)
		return
	}
	_,_,err = Client.Ping(host).Do(context.Background())
	if err != nil {
		logger.Error(err)
		return
	}
	_,err = Client.ElasticsearchVersion(host)
	if err != nil {
		logger.Error(err)
		return
	}
	return
}


func CreatBlockData (i interface{}) error{
	_,err := Client.Index().
		Index("walletMonitor").  //数据库
		BodyJson(i).
		Do(context.Background())
	if err != nil{
		logger.Error(err)
		return err
	}
	return nil
}

