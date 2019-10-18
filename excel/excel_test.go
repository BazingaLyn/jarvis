package excel

import (
	"fmt"
	"github.com/BazingaLyn/jarvis/model"
	"strconv"
	"strings"
	"testing"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func TestExcel_read(t *testing.T) {
	f, err := excelize.OpenFile("./movies.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	rows := f.GetRows("Sheet1")
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
	}
}
