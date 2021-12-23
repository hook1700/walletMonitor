package routes

import (
	"github.com/gin-gonic/gin"
	"redisData/middleware"
)

func SetUp() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors()) //跨域
	r.Use(middleware.TraceLogger())
	//r.Use(middleware.TlsHandler())  // 支持wss
	 // 日志上下文进行绑定追踪
	//查询，查询redis上的数据，返回给前端
	//websocket
	wallet := r.Group("/api/wallet")
	RegisterWebRoutes(wallet)
	return r
}
