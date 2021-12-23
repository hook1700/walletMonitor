package routes

import (
	"github.com/gin-gonic/gin"
	"redisData/controller"
)

// RegisterWebRoutes 注册路由 baby 游戏路由

var walletController = new(controller.Controller)

func RegisterWebRoutes(router *gin.RouterGroup) {
	//通过地址获取hash列表
	router.Any("/getTxHashListByAddress", walletController.GetTxHashListByAddress)
	//通过钱包地址和交易地址 获取hash列表MonitorWallet
	router.Any("/monitorWallet", walletController.MonitorWallet)
}
