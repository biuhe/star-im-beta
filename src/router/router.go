package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"star-im/src/docs"
	"star-im/src/service"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.GET("/index", service.GetIndex)
	r.POST("/login", service.Login)
	r.GET("/user/list", service.GetUserList)
	r.POST("/user/create", service.CreateUser)
	r.POST("/user/update", service.UpdateUser)
	r.GET("/user/delete", service.DeleteUser)

	// programatically set swagger info
	docs.SwaggerInfo.Title = "Star-Im"
	docs.SwaggerInfo.Description = "即时通讯接口文档"
	docs.SwaggerInfo.Version = "1.0"
	//docs.SwaggerInfo.Host = "petstore.swagger.io"
	//docs.SwaggerInfo.BasePath = "/v2"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
