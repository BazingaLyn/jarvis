package handlers

import (
	"github.com/BazingaLyn/jarvis/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var movieMap = make(map[string]model.Movie)

func GetAllMovie(c *gin.Context) {

	movies := make([]model.Movie, 0)

	for _, v := range movieMap {
		movies = append(movies, v)
	}
	c.JSON(http.StatusOK, movies)

}

func GetDefaultMovieById(c *gin.Context) {
	movieId := c.Param("id")
	movie := movieMap[movieId]
	c.JSON(http.StatusOK, movie)
}

func BatchSaveMovie(c *gin.Context) {
	var movies []model.Movie

	c.Bind(&movies)

	for _, v := range movies {
		movieMap[v.Id] = v
	}

	c.JSON(http.StatusOK, "ok")

}

/**
POST
{
    "Id": "1",
    "Name": "god like",
    "Score": 45.6,
    "Actors": [
        "成龙"
    ],
    "Directors": [
        "吕克贝松",
        "张艺谋"
    ],
    "MovieTime": 127
}
*/
func SaveMovie(c *gin.Context) {
	var movie model.Movie

	c.BindJSON(&movie)

	log.Println(movie.Name, movie.Directors, movie.Actors)
	movieMap[movie.Id] = movie
	c.JSON(http.StatusOK, "ok")
}

func AddMovieDirector(c *gin.Context) {
	movieId := c.Param("movieId")
	directorName := c.Param("directorName")

	log.Println(movieId + directorName)

}
