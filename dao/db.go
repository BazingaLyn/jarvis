package dao

import (
	"database/sql"
	"github.com/BazingaLyn/jarvis/model"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strings"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("mysql", "root:1qaz2wsx@tcp(47.98.164.130:3306)/jarvis")
	if err != nil {
		log.Panicln("err:", err.Error())
	}
	Db.SetMaxOpenConns(10)
	Db.SetMaxIdleConns(10)
}

func Save(movie *model.ElasticMovie) int64 {
	result, e := Db.Exec("insert into jarvis.t_movie (movie_name, score, movie_type, directors, actors, nations, languages, file_length, description) values (?,?,?,?,?,?,?,?,?);",
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
