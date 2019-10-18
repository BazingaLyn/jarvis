package routers

import (
	"github.com/BazingaLyn/jarvis/handlers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const APIBase = "api/v1/test"

func Routers() *gin.Engine {

	router := gin.New()
	router.Use(gin.Logger())
	v1 := router.Group(APIBase)
	{
		v1.GET("/movie", handlers.GetAllMovie)
		v1.GET("/movie/:id", handlers.GetDefaultMovieById)
		v1.POST("/save/movie", handlers.SaveMovie)
		v1.POST("/save/elastic/movie", handlers.SaveElasticMovie)
		v1.POST("/batchSaveMovie", handlers.BatchSaveMovie)
		v1.POST("/saveDirector", handlers.AddMovieDirector)
	}

	url := ginSwagger.URL("http://localhost:10086/swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router

}
