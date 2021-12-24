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
	"github.com/spf13/viper"
)
var client *elastic.Client

//InitES ES初始化
func InitES()  {
	//host := viper.GetString("elasticsearch.host")
	//这个地方有个小坑 不加上elastic.SetSniff(false) 会连接不上
	//Client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(host))
	var err error
	url := fmt.Sprintf("%s",viper.GetString("elasticsearch.host"))
	client,err  = elastic.NewClient(elastic.SetURL(url))
	if err != nil {
		fmt.Println(111,err)
		return
	}
	_,_,err = client.Ping(url).Do(context.Background())
	if err != nil {
		fmt.Println(222,err)
		return
	}
	_,err = client.ElasticsearchVersion(url)
	if err != nil {
		fmt.Println(333,err)
		return
	}
	return
}


//CreatBlockData 插入一个区块的数据
func CreatBlockData (data interface{},db string) error{
	_,err := client.Index().
		Index(db). //数据库
		BodyJson(data).
 		Do(context.Background())
	if err != nil{
		fmt.Println(err)
		return err
	}
	return nil
}






