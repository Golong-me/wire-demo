package httpMiddle

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func LoadMiddleware(logger *zap.Logger, engine *gin.Engine) {
	//跨域中间件
	engine.Use(cors.New(cors.Config{
		AllowAllOrigins:  true, // 允许所有来源
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
	}))

	//引入ginzap中间件，gin的日志打印由zap来完成

	//统一错误处理中间件,所有的错误向外抛出，由gin处理

	//token中间件
}
