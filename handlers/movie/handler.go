package movie

import (
	"github.com/BazingaLyn/jarvis/constants"
	"github.com/BazingaLyn/jarvis/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary 批量保存电影信息
// @Tags 电影模块
// @version 1.0
// @Accept application/x-json-stream
// @Param movie body model.Movie true "批量保存电影基本信息"
// @Success 200 object model.Response 成功后返回值
// @Router /movie/batch/save [post]
func BatchSaveMovie(context *gin.Context) {
	movies := model.Movies{}

	e := context.ShouldBindJSON(&movies)
	code := constants.FAIL
	if e == nil {
		flag := movies.BatchSaveMovie()
		if flag {
			code = constants.SUCCESS
		}
	}

	context.JSON(http.StatusOK, model.Response{
		Code: code,
		Msg:  "",
		Data: nil,
	})
}

// @Summary 获取某个电影的信息
// @Tags 电影模块
// @version 1.0
// @Accept application/x-json-stream
// @Param id path int true "电影id"
// @Success 200 object model.Response 成功后返回值
// @Router /movie/query/{id} [get]
func GetMovieById(context *gin.Context) {

	idStr := context.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		context.JSON(http.StatusOK, model.Response{
			Code: constants.FAIL,
			Msg:  "movie id should be int type",
			Data: nil,
		})
		return
	}

	movie := model.Movie{
		Id: id,
	}

	result := movie.GetMovieById()

	context.JSON(http.StatusOK, model.Response{
		Code: constants.SUCCESS,
		Msg:  "success",
		Data: result,
	})

}

// @Summary 保存电影信息
// @Tags 电影模块
// @version 1.0
// @Accept application/x-json-stream
// @Param movie body model.Movie true "保存的电影基本信息"
// @Success 200 object model.Response 成功后返回值
// @Router /movie/save [post]
func SaveMovie(context *gin.Context) {
	movie := model.Movie{}

	e := context.ShouldBindJSON(&movie)
	code := constants.FAIL
	if e == nil {
		id := movie.Save()
		if id > 0 {
			code = constants.SUCCESS
		}
	}

	context.JSON(http.StatusOK, model.Response{
		Code: code,
		Msg:  "",
		Data: nil,
	})
}

// @Summary 根据id删除电影信息
// @Tags 电影模块
// @version 1.0
// @Accept application/x-json-stream
// @Param id path int true "要删除的电影id"
// @Success 200 object model.Response 成功后返回值
// @Router /movie/delete/{id} [get]
func DeleteMovie(context *gin.Context) {

	idStr := context.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		context.JSON(http.StatusOK, model.Response{
			Code: constants.FAIL,
			Msg:  "movie id should be int type",
			Data: nil,
		})
		return
	}

	movie := model.Movie{
		Id: id,
	}

	code := constants.FAIL

	flag := movie.DeleteMovieById()

	if flag {
		code = constants.SUCCESS
	}

	context.JSON(http.StatusOK, model.Response{
		Code: code,
		Msg:  "",
		Data: nil,
	})
}
