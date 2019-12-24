package routers

import (
	"github.com/BazingaLyn/jarvis/handlers/movie"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Routers() *gin.Engine {

	routers := gin.Default()
	movieRouter := routers.Group("/movie")
	{
		movieRouter.GET("/query/:id", movie.GetMovieById)
		movieRouter.POST("/save", movie.SaveMovie)
		movieRouter.POST("/batch/save", movie.BatchSaveMovie)
		movieRouter.GET("/delete/:id", movie.DeleteMovie)
	}
	url := ginSwagger.URL("http://47.98.164.130:8080/swagger/doc.json")
	routers.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	return routers
}
