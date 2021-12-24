/**
 @author:way
 @date:2021/12/23
 @note
**/

package elasticsearch

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
)
var client *elastic.Client
//InitES ES初始化
func InitES()  {
	//host := viper.GetString("elasticsearch.host")
	//这个地方有个小坑 不加上elastic.SetSniff(false) 会连接不上
	//Client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(host))
	var err error
	client,err  = elastic.NewClient(elastic.SetURL("http://10.10.10.8:9200"))
	if err != nil {
		fmt.Println(111,err)
		return
	}
	_,_,err = client.Ping("http://10.10.10.8:9200").Do(context.Background())
	if err != nil {
		fmt.Println(222,err)
		return
	}
	_,err = client.ElasticsearchVersion("http://10.10.10.8:9200")
	if err != nil {
		fmt.Println(333,err)
		return
	}
	return
}

func CreatBlockData (data interface{}) error{

	_,err := client.Index().
		Index("wallet"). //数据库
		BodyJson(data).
 		Do(context.Background())
	if err != nil{
		fmt.Println(err)
		return err
	}
	return nil
}

func GetBlockData()  {
	
}



