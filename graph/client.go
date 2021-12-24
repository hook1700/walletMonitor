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


var QueryBlockNum struct {
	Block struct {
		Number graphql.Int
	}

}
var QueryBlockNumData struct {
	Block struct {
		Number graphql.Int
		//Hash   graphql.String
		//Parent struct {
		//	Number           graphql.Int
		//	TransactionCount graphql.Int
		//}
		//TransactionsRoot graphql.String
		//TransactionCount graphql.Int
		//Miner            struct {
		//	Address          graphql.String
		//	Balance          graphql.String
		//	TransactionCount graphql.String
		//	Code             graphql.String
		//}
		//Timestamp graphql.String
		//Transactions []struct {
		//	Hash graphql.String
		//	From struct {
		//		Address          graphql.String
		//		TransactionCount graphql.String
		//	}
		//	To struct {
		//		Address          graphql.String
		//		TransactionCount graphql.String
		//	}
		//	Value    graphql.String
		//	Status          graphql.Int
		//	CreatedContract struct {
		//		Address string
		//	}
		//	Logs []struct {
		//		Account struct {
		//			Address graphql.String
		//		}
		//		Topics []graphql.String
		//		Data   graphql.String
		//	}
		//}
	} `graphql:"block(number:$number)"`
}




//0x14655169c60948b969b1a19ec8113b134a91280fb58c0f35c083f85821032eb1
func main() {
	//q := "query: \"{\\n  block(number: 13524585) {\\n    number\\n    hash\\n    parent {\\n      number\\n      transactionCount\\n    }\\n    transactionsRoot\\n    transactionCount\\n    miner {\\n      address\\n      balance\\n      transactionCount\\n      code\\n    }\\n    timestamp\\n    transactions {\\n      hash\\n      from {\\n        address\\n        transactionCount\\n      }\\n      to {\\n        address\\n        transactionCount\\n      }\\n      value\\n      status\\n      createdContract {\\n        address\\n      }\\n      logs {\\n        account {\\n          address\\n        }\\n        topics\\n        data\\n      }\\n    }\\n  }\\n}\\n\"\n"
	//链接GQL
	client := graphql.NewClient("https://wb.xfack.com/graphql", nil)
	//hash := "0x14655169c60948b969b1a19ec8113b134a91280fb58c0f35c083f85821032eb1"
	//type Long uint64
	var number uint64
	number = 13530647
	variables := map[string]interface{}{
		"number":graphql.Int(number) ,
	}
	err := client.Query(context.Background(), &QueryBlockNumData,variables)
	if err != nil {
		//fmt.Println(1)
		fmt.Println(err)
		//logger.Error(err)
		return
	}
	//fmt.Printf("%T",QueryBlockNum.Block.Number)
	dataByte, err1 := json.Marshal(&QueryBlockNumData)
	if err1 != nil {
		fmt.Println(2)
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
