package conf

import (
	"context"
	"gin-wire/internal/controller"
	"gin-wire/internal/utils/httpMiddle"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)
import "github.com/gin-gonic/gin"

type HttpServer struct {
	Logger *zap.Logger
	Engine *gin.Engine

	//我们写了什么controller,需要在这里进行加入
	//并且编写对应的provider
	UserCtl *controller.UserController
}

//kratos layout grpc
//func NewHttpServer(logger *zap.Logger, engine *gin.Engine,
//	UserCtl *controller.UserController) *HttpServer {
//
//

//	//UserCtl.initRouter(engine)
//}

func NewGin(logger *zap.Logger) *gin.Engine {
	engine := gin.New()

	//加载中间件
	httpMiddle.LoadMiddleware(logger, engine)

	return engine
}

type HttpServerConfig struct {
	Port string `json:"port"`
}

func (receiver *HttpServer) RunServer() {

	var config HttpServerConfig

	//viper读取配置等操作

	//注册全局验证器等操作

	//测试，所以就手动赋值
	config.Port = "9001"

	//启动gin
	receiver.ginServer(config)
}

func (receiver *HttpServer) ginServer(config HttpServerConfig) {
	httpServer := &http.Server{
		Addr:         ":" + config.Port,
		Handler:      receiver.Engine,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	//协程启动服务器
	go func() {
		err := httpServer.ListenAndServe()
		if err != nil {
			receiver.Logger.Fatal("服务结束：", zap.Error(err))
		}
	}()

	//平滑重启
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := httpServer.Shutdown(ctx); err != nil {
		receiver.Logger.Error("关闭服务错误", zap.Error(err))
	}
	receiver.Logger.Info("释放资源中....")
	<-ctx.Done()
	receiver.Logger.Info("退出了....")
}
