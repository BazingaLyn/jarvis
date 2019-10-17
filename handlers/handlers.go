package handlers

import (
	"github.com/BazingaLyn/jarvis/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

var movieMap = make(map[string]model.Movie)

// @Summary 获取所有的电影
// @Tags 电影
// @Accept application/x-json-stream
// @Success 200 object model.Result 成功后返回值
// @Router /movie [get]
func GetAllMovie(c *gin.Context) {

	movies := make([]model.Movie, 0)

	for _, v := range movieMap {
		movies = append(movies, v)
	}

	result := model.Result{
		Code:    http.StatusOK,
		Message: "success",
		Data:    movies,
	}

	c.JSON(http.StatusOK, result)

}

// @Summary 根据电影id获取电影信息
// @Tags 电影
// @Success 200 object model.Result 成功后返回值
// @Router /movie/:id [get]
func GetDefaultMovieById(c *gin.Context) {
	movieId := c.Param("id")
	movie := movieMap[movieId]

	result := model.Result{
		Code:    http.StatusOK,
		Message: "success",
		Data:    movie,
	}

	c.JSON(http.StatusOK, result)
}

// @Summary 批量保存电影信息
// @Tags 电影
// @Success 200 object model.Result 成功后返回值
// @Router /batchSaveMovie [post]
func BatchSaveMovie(c *gin.Context) {
	var movies []model.Movie

	c.Bind(&movies)

	for _, v := range movies {
		movieMap[v.Id] = v
	}

	result := model.Result{
		Code:    http.StatusOK,
		Message: "success",
		Data:    nil,
	}

	c.JSON(http.StatusOK, result)

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
// @Summary 保存电影信息
// @Tags 电影
// @Success 200 object model.Result 成功后返回值
// @Router /save/movie [post]
func SaveMovie(c *gin.Context) {
	var movie model.Movie

	c.BindJSON(&movie)

	movieMap[movie.Id] = movie

	result := model.Result{
		Code:    http.StatusOK,
		Message: "success",
		Data:    nil,
	}

	c.JSON(http.StatusOK, result)
}

// @Summary 修改电影信息
// @Tags 电影
// @Success 200 object model.Result 成功后返回值
// @Router /saveDirector/:movieId/:directorName [post]
func AddMovieDirector(c *gin.Context) {
	movieId := c.Query("movieId")
	directorName := c.Query("directorName")

	if movieId == "" || directorName == "" {
		result := model.Result{
			Code:    http.StatusBadRequest,
			Message: "failed",
			Data:    nil,
		}
		c.JSON(http.StatusBadRequest, result)
		return
	}

	movie := movieMap[movieId]
	movie.Directors = append(movie.Directors, directorName)
	movieMap[movieId] = movie

	result := model.Result{
		Code:    http.StatusOK,
		Message: "success",
		Data:    movie,
	}

	c.JSON(http.StatusOK, result)

}
