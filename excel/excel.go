package excel

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/BazingaLyn/jarvis/model"
	"log"
	"strconv"
	"strings"
)

func ReadExcel(dir string, movieDataChan chan<- []model.ElasticMovie) {

	f, err := excelize.OpenFile(dir)
	if err != nil {
		log.Println(err)
		return
	}

	rows := f.GetRows("Sheet1")
	var movies []model.ElasticMovie
	for line, row := range rows {
		if line == 0 {
			continue
		}
		var movie model.ElasticMovie
		movie.Id = line
		for index, colCell := range row {
			if index == 0 {
				movie.MovieName = colCell
			}
			if index == 1 {
				score, err := strconv.ParseFloat(colCell, 32)
				if err == nil {
					movie.Score = score
				}
			}
			if index == 10 {
				types := strings.Split(colCell, "/")
				for _, v := range types {
					movie.Type = append(movie.Type, v)
				}
			}

			if index == 11 {
				directors := strings.Split(colCell, "/")
				for _, v := range directors {
					movie.Directors = append(movie.Directors, v)
				}
			}

			if index == 12 {
				actors := strings.Split(colCell, "/")
				for _, v := range actors {
					if v == "null" {
						movie.Actors = append(movie.Actors, "未知")
					} else {

						movie.Actors = append(movie.Actors, v)
					}
				}
			}

			if index == 14 {
				nationStr := strings.Split(colCell, "/")
				for _, v := range nationStr {
					movie.Nations = append(movie.Nations, v)
				}
			}

			if index == 15 {
				languageStr := strings.Split(colCell, "/")
				for _, v := range languageStr {
					movie.Languages = append(movie.Languages, v)
				}
			}

			if index == 17 {
				i, _ := strconv.Atoi(colCell)
				if i == 0 {
					movie.FileLength = 120
				} else {
					movie.FileLength = i
				}
			}

			if index == 22 {
				if colCell == "null" {
					movie.Describe = "无"
				} else {
					movie.Describe = colCell
				}
			}
		}
		if len(movies) == 5 {
			movieDataChan <- movies
			movies = movies[0:0]
		} else {
			movies = append(movies, movie)
		}
	}
	if len(movies) != 0 {
		movieDataChan <- movies
	}

	close(movieDataChan)
}
