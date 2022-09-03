package web

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"jwk/test/pkg/config"
	"jwk/test/pkg/handler"
	"jwk/test/utils"
	"strconv"
)

func WebServer() {
	// 强制日志颜色化
	serverConfig := config.Config
	server := serverConfig.Server
	//router := gin.Default()
	router := gin.New()
	gin.ForceConsoleColor()
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		utils.Logger.Info("Endpoint:", zap.String("Method", httpMethod), zap.String("Path", absolutePath), zap.String("Handler", handlerName), zap.Int("NuHandlers", nuHandlers))
	}

	router.Use(handler.GinLogger(utils.Logger), handler.GinRecovery(utils.Logger, true), handler.Cors())
	router.GET("/test", handler.GetSysUserHandler())
	router.GET("/test1", handler.UpdateSysUserHandler())

	serverAddr := server.Host + ":" + strconv.Itoa(server.Port)
	utils.Logger.Info("WebServer Start, Addr: ", zap.Any("serverAddr", serverAddr))
	router.Run(serverAddr)
}
