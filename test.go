/**
 @author:way
 @date:2021/12/24
 @note
**/

package main

import (
	"redisData/dao/elasticsearch"
	"redisData/logic"
)

func main() {
	elasticsearch.InitES()
	for i :=13530647;i > 13530637;i--{
		//db 为ES的db
		logic.GetDataByGraphQL(i,"wallet")
	}
}
