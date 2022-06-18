package api

import (
	//_ "github.com/baxromumarov/work/api-gateway/api/docs"
	v1 "github.com/baxromumarov/work/api-gateway/api/handlers/v1"
	"github.com/baxromumarov/work/api-gateway/config"
	"github.com/baxromumarov/work/api-gateway/pkg/logger"
	"github.com/baxromumarov/work/api-gateway/services"
	"github.com/gin-gonic/gin"
)

// Option ...
type Option struct {
	Conf           config.Config
	Logger         logger.Logger
	ServiceManager services.IServiceManager
}

// New ...
func New(option Option) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	handlerV1 := v1.New(&v1.HandlerV1Config{
		Logger:         option.Logger,
		ServiceManager: option.ServiceManager,
		Cfg:            option.Conf,
	})

	// https://gorest.co.in/public/v1
	api := router.Group("/v1")
	api.POST("/data/insert", handlerV1.CreateData)
	api.GET("/data/list", handlerV1.GetDataList)
	api.GET("/data/:id", handlerV1.GetDataById)
	api.PUT("/data/:id", handlerV1.UpdateData)
	api.DELETE("/data/:id", handlerV1.DeleteData)

	return router
}
