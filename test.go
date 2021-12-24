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

	logic.GetDataByGraphQL(13530647)
}
