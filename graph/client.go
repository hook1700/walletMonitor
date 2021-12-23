/**
 @author:way
 @date:2021/12/23
 @note
**/

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hasura/go-graphql-client"
	"redisData/pkg/logger"
)

//var query struct {
//	Block struct {
//		Name   graphql.String
//		Height graphql.Float `graphql:"height(unit: METER)"`
//	} `graphql:"human(id: \"1000\")"`
//}

var QueryBlockNum struct {
	Block struct {
		Number graphql.Int
	}
}

var QueryBlockNumData struct {
	Block struct {
		Number graphql.Int
		Hash   graphql.String
		Parent struct {
			Number           graphql.Int
			TransactionCount graphql.Int
		}
		TransactionsRoot graphql.String
		TransactionCount graphql.Int
		Miner            struct {
			Address          graphql.String
			Balance          graphql.String
			TransactionCount graphql.String
			Code             graphql.String
		}
		Timestamp graphql.String
		Transactions []struct {
			Hash graphql.String
			From struct {
				Address          graphql.String
				TransactionCount graphql.String
			}
			To struct {
				Address          graphql.String
				TransactionCount graphql.String
			}
			Value    graphql.String
			Status          graphql.Int
			CreatedContract struct {
				Address string
			}
			Logs []struct {
				Account struct {
					Address graphql.String
				}
				Topics []graphql.String
				Data   graphql.String
			}
		}
	} `graphql:"block(number:13529864)"`
}



func main() {
	//链接GQL
	client := graphql.NewClient("https://wb.xfack.com/graphql", nil)



	err := client.Query(context.Background(), &QueryBlockNumData, nil)
	if err != nil {
		logger.Error(err)
		return
	}
	fmt.Println(QueryBlockNumData)
	dataByte, err1 := json.Marshal(&QueryBlockNumData)
	if err1 != nil {
		logger.Error(err1)
		return
	}
	fmt.Println(string(dataByte))
	//err2 := elasticsearch.CreatBlockData(QueryBlockNumData)
	//if err2 != nil {
	//	logger.Error(err2)
	//	return
	//}
}
