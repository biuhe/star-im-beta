package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"star-im/src/main/docs"
	"star-im/src/main/service"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.POST("/login", service.Login)
	r.GET("/register", service.ToRegister)
	r.POST("/register", service.CreateUser)

	r.GET("/user/list", service.GetUserList)
	r.POST("/user/create", service.CreateUser)
	r.POST("/user/update", service.UpdateUser)
	r.GET("/user/delete", service.DeleteUser)

	// 发送消息
	r.GET("/ws/send", service.SendMsg)
	r.GET("/ws/chat", service.Chat)

	// 静态资源
	r.Static("/asset", "src/resource/asset/")
	r.StaticFile("/favicon.ico", "src/resource/asset/images/favicon.ico")
	r.LoadHTMLGlob("src/resource/views/**/*")
	r.GET("/", service.GetIndex)
	r.GET("/index", service.GetIndex)

	// swagger info
	docs.SwaggerInfo.Title = "Star-Im"
	docs.SwaggerInfo.Description = "即时通讯接口文档"
	docs.SwaggerInfo.Version = "1.0"
	//docs.SwaggerInfo.Host = "petstore.swagger.io"
	//docs.SwaggerInfo.BasePath = "/v2"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
