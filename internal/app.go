package internal

import (
	_ "golang-template-module/docs"
	handler "golang-template-module/internal/delivery/handlers"
	httpv1 "golang-template-module/internal/delivery/http/v1"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func RunListening() {
	r := gin.New()
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())
	v1 := r.Group("/api/v1")
	{
		v1.POST("/auth", handler.ErrorHandler(httpv1.RouterAuth))
		v1.POST("/create", handler.ErrorHandler(httpv1.RouterRegistration))
		user := v1.Group("/user")
		{
			user.GET("/:id", handler.ErrorHandler(httpv1.RouterUserGetId))
			user.GET("/list", handler.ErrorHandler(httpv1.RouterUserGetList))
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run()
}
