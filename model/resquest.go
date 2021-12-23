/**
 @author:way
 @date:2021/12/22
 @note
**/

package model

//ParamAddress 参数为一个地址
type ParamAddress struct {
	Address string `json:"address" form:"address"` //钱包地址或者合约地址
}

//ParamAddressAndContract 参数为一个合约地址，和一个钱包地址
type ParamAddressAndContract struct {
	Address         string `json:"address" form:"address"`                   //钱包地址
	ContractAddress string `json:"contract_address" form:"contract_address"` //合约地址
}
