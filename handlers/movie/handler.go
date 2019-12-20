package movie

import (
	"github.com/BazingaLyn/jarvis/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary 获取某个电影的信息
// @Tags 电影模块
// @version 1.0
// @Accept application/x-json-stream
// @Param id path int true "电影id"
// @Success 200 object model.Response 成功后返回值
// @Router /movie/get/{id} [get]
func GetMovieById(context *gin.Context) {

	idStr := context.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		context.JSON(http.StatusOK, model.Response{
			Code: 2,
			Msg:  "movie id should be int type",
			Data: nil,
		})
		return
	}

}

// @Summary 保存电影信息
// @Tags 电影模块
// @version 1.0
// @Accept application/x-json-stream
// @Param movie body model.Movie true "保存的电影基本信息"
// @Success 200 object model.Response 成功后返回值
// @Router /movie/save [post]
func SaveMovie(context *gin.Context) {

}

// @Summary 根据id删除电影信息
// @Tags 电影模块
// @version 1.0
// @Accept application/x-json-stream
// @Param id path int true "要删除的电影id"
// @Success 200 object model.Response 成功后返回值
// @Router /movie/delete/{id} [get]
func DeleteMovie(context *gin.Context) {

}
