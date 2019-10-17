package routers

import (
	"github.com/BazingaLyn/jarvis/handlers"
	"github.com/gin-gonic/gin"
)

const APIBase = "api/v1/test"

//noinspection GoUnresolvedReference
func Routers() *gin.Engine {

	router := gin.New()
	router.Use(gin.Logger())
	v1 := router.Group(APIBase)
	{
		v1.GET("/movie", handlers.GetAllMovie)
		v1.GET("/movie/:id", handlers.GetDefaultMovieById)
		v1.POST("/save/movie", handlers.SaveMovie)
		v1.POST("/batchSaveMovie", handlers.BatchSaveMovie)
		v1.POST("/saveDirector/:movieId/:directorName", handlers.AddMovieDirector)
	}

	return router

}
