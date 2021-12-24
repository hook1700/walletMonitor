/**
 @author:way
 @date:2021/12/22
 @note
**/

package model

//RespTxList 通过合约或者钱包地址，返回交易列表
type RespTxList struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  []struct {
		BlockNumber       string `json:"blockNumber"`
		TimeStamp         string `json:"timeStamp"`
		Hash              string `json:"hash"`
		Nonce             string `json:"nonce"`
		BlockHash         string `json:"blockHash"`
		TransactionIndex  string `json:"transactionIndex"`
		From              string `json:"from"`
		To                string `json:"to"`
		Value             string `json:"value"`
		Gas               string `json:"gas"`
		GasPrice          string `json:"gasPrice"`
		IsError           string `json:"isError"`
		TxreceiptStatus   string `json:"txreceipt_status"`
		Input             string `json:"input"`
		ContractAddress   string `json:"contractAddress"`
		CumulativeGasUsed string `json:"cumulativeGasUsed"`
		GasUsed           string `json:"gasUsed"`
		Confirmations     string `json:"confirmations"`
	} `json:"result"`
}

//RespMonitorList 币安监控钱包接口,通过地址和合约地址监控钱包下某合约的交易,监控ERC20钱包
type RespMonitorList struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  []struct {
		BlockNumber       string `json:"blockNumber"`
		TimeStamp         string `json:"timeStamp"`
		Hash              string `json:"hash"`
		Nonce             string `json:"nonce"`
		BlockHash         string `json:"blockHash"`
		From              string `json:"from"`
		ContractAddress   string `json:"contractAddress"`
		To                string `json:"to"`
		Value             string `json:"value"`
		TokenName         string `json:"tokenName"`
		TokenSymbol       string `json:"tokenSymbol"`
		TokenDecimal      string `json:"tokenDecimal"`
		TransactionIndex  string `json:"transactionIndex"`
		Gas               string `json:"gas"`
		GasPrice          string `json:"gasPrice"`
		GasUsed           string `json:"gasUsed"`
		CumulativeGasUsed string `json:"cumulativeGasUsed"`
		Input             string `json:"input"`
		Confirmations     string `json:"confirmations"`
	} `json:"result"`
}

//RespGraphData 接受GraphServer返回的数据
type RespGraphData struct {
	Data struct {
		Block struct {
			Number int    `json:"number"`
			Hash   string `json:"hash"`
			Parent struct {
				Number           int `json:"number"`
				TransactionCount int `json:"transactionCount"`
			} `json:"parent"`
			TransactionsRoot string `json:"transactionsRoot"`
			TransactionCount int    `json:"transactionCount"`
			Miner            struct {
				Address          string `json:"address"`
				Balance          string `json:"balance"`
				TransactionCount string `json:"transactionCount"`
				Code             string `json:"code"`
			} `json:"miner"`
			Timestamp    string `json:"timestamp"`
			Transactions []struct {
				Hash string `json:"hash"`
				From struct {
					Address          string `json:"address"`
					TransactionCount string `json:"transactionCount"`
				} `json:"from"`
				To struct {
					Address          string `json:"address"`
					TransactionCount string `json:"transactionCount"`
				} `json:"to"`
				Value           string      `json:"value"`
				Status          int         `json:"status"`
				CreatedContract interface{} `json:"createdContract"`
				Logs            []struct {
					Account struct {
						Address string `json:"address"`
					} `json:"account"`
					Topics []string `json:"topics"`
					Data   string   `json:"data"`
				} `json:"logs"`
			} `json:"transactions"`
		} `json:"block"`
	} `json:"data"`
}