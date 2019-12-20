package movie

import (
	"github.com/BazingaLyn/jarvis/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary 获取某个电影的信息
// @Tags 电影模块
// @version 1.0
// @Accept application/x-json-stream
// @Param id path int true "电影id"
// @Success 200 object model.Response 成功后返回值
// @Router /movie/get/{id} [get]
func GetMovieById(context *gin.Context) {

	id := context.Param("id")

	context.JSON(http.StatusOK, gin.H{
		"result": model.Response{
			Code: http.StatusOK,
			Msg:  "success",
			Data: model.Movie{
				Id:   id,
				Name: "大话西游",
			},
		},
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

	id := context.Param("id")

	context.JSON(http.StatusOK, gin.H{
		"result": model.Response{
			Code: http.StatusOK,
			Msg:  "success",
			Data: model.Movie{
				Id:   id,
				Name: "大话西游",
			},
		},
	})
}

// @Summary 保存电影信息
// @Tags 电影模块
// @version 1.0
// @Accept application/x-json-stream
// @Param movie body model.Movie true "保存的电影基本信息"
// @Success 200 object model.Response 成功后返回值
// @Router /movie/delete/{id} [get]
func DeleteMovie(context *gin.Context) {

	id := context.Param("id")

	context.JSON(http.StatusOK, gin.H{
		"result": model.Response{
			Code: http.StatusOK,
			Msg:  "success",
			Data: model.Movie{
				Id:   id,
				Name: "大话西游",
			},
		},
	})
}
