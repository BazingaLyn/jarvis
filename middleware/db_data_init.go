package middleware

import (
	"github.com/BazingaLyn/jarvis/excel"
	"github.com/BazingaLyn/jarvis/model"
	"log"
)

func MovieExcelDataToDbInit() {
	ch := make(chan []model.ElasticMovie, 5)
	go excel.ReadExcel("./excel/movies.xlsx", ch)

	go SaveMoviesToDb(ch)

}

func SaveMoviesToDb(ch <-chan []model.ElasticMovie) {
	for {

		movies, ok := <-ch
		if !ok {
			log.Println("Receiver :close channel")
			break
		}

		model.ElasticMovies = movies

		for i, eachMovie := range movies {

			log.Println(i)
			log.Println(eachMovie)
			eachMovie.Save()
		}

	}
}
