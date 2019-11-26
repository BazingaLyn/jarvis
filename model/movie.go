package model

import (
	"github.com/BazingaLyn/jarvis/dao"
	"log"
	"strings"
)

type ElasticMovie struct {
	Id         int
	MovieName  string
	Score      float64
	Type       []string
	Directors  []string
	Actors     []string
	Nations    []string
	Languages  []string
	FileLength int
	Describe   string
}

//func (movies *[]ElasticMovie) Save(){
//
//}

func (movie *ElasticMovie) Save() int64 {
	result, e := dao.Db.Exec("insert into jarvis.t_movie (movie_name, score, movie_type, directors, actors, nations, languages, file_length, description) values (?,?,?,?,?,?,?,?,?);",
		movie.MovieName,
		movie.Score,
		strings.Join(movie.Type, ","),
		strings.Join(movie.Directors, ","),
		strings.Join(movie.Actors, ","),
		strings.Join(movie.Nations, ","),
		strings.Join(movie.Languages, ","),
		movie.FileLength,
		movie.Describe,
	)
	if e != nil {
		log.Panicln("movie insert error", e.Error())
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Panicln("movie insert id error", err.Error())
	}
	return id
}
