package router

import (
	"CODE/controller"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(tagsController *controller.TagsController) *gin.Engine {
	routes := gin.Default()
	//add swagger
	routes.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	routes.GET("", func(ctx *gin.Context) {
		ctx.JSON(200, "Welcome Home")
	})
	baseRouter := routes.Group("/api")
	tagsRouter := baseRouter.Group("/tags")
	tagsRouter.GET("", tagsController.FindAll)
	tagsRouter.GET("/:tagId", tagsController.FindById)
	tagsRouter.POST("", tagsController.FindAll)
	tagsRouter.PATCH("/:tagId", tagsController.Update)
	tagsRouter.DELETE("/:tagId", tagsController.Delete)
	return routes
}
