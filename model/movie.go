package model

import (
	"database/sql"
	"github.com/BazingaLyn/jarvis/db"
	"log"
)

type Movie struct {
	Id         int     `json:"id"`
	MovieName  string  `json:"movieName"`
	Score      float64 `json:"score"`
	Type       string  `json:"type"`
	Directors  string  `json:"directors"`
	Actors     string  `json:"actors"`
	Nations    string  `json:"nations"`
	Languages  string  `json:"languages"`
	FileLength int     `json:"fileLength"`
	Describe   string  `json:"describe"`
}

type Movies struct {
	Movies []Movie `json:"movies"`
}

func (movies *Movies) BatchSaveMovie() bool {

	stmt, err := db.Db.Prepare("insert into jarvis.t_movie (movie_name, score, movie_type, directors, actors, nations, languages, file_length, description) values (?,?,?,?,?,?,?,?,?);")

	if err != nil {
		log.Println("prepare failed ", err.Error())
		return false
	}

	defer stmt.Close()

	for _, movie := range movies.Movies {
		if _, err := stmt.Exec(movie.MovieName,
			movie.Score,
			movie.Type,
			movie.Directors,
			movie.Actors,
			movie.Nations,
			movie.Languages,
			movie.FileLength,
			movie.Describe); err != nil {
			log.Println("exec failed ", err.Error())
		}
	}
	return true

}

func (movie *Movie) DeleteMovieById() bool {
	result, e := db.Db.Exec("delete from t_movie where id = ?", movie.Id)

	if e != nil {
		log.Println("delete movie failed", e.Error())
		return false
	}

	i, e := result.RowsAffected()
	if e != nil {
		log.Println("delete movie failed", e.Error())
		return false
	}

	if i < 1 {
		log.Printf("not found id %d movie,delete failed", movie.Id)
		return false
	}

	return true
}

func (movie *Movie) GetMovieById() Movie {
	result := db.Db.QueryRow("select * from t_movie where id = ?", movie.Id)

	movieResults := Movie{}

	err := result.Scan(&movieResults.Id,
		&movieResults.MovieName,
		&movieResults.Score,
		&movieResults.Type,
		&movieResults.Directors,
		&movieResults.Actors,
		&movieResults.Nations,
		&movieResults.Languages,
		&movieResults.FileLength,
		&movieResults.Describe)

	if err != nil {
		if err == sql.ErrNoRows {
			return movieResults
		} else {
			log.Panicln("select movie failed", err.Error())
		}
	}

	return movieResults
}

//{
//"actors": "甄子丹/吴樾/吴建豪",
//"describe": "因故来到美国唐人街的叶问，意外卷入一场当地军方势力与华人武馆的纠纷，面对日益猖狂的民族歧视与压迫，叶问挺身而出，在美国海军陆战队军营拼死一战，以正宗咏春，向世界证明了中国功夫",
//"directors": "叶伟信",
//"fileLength": 107,
//"languages": "粤语/英语",
//"movieName": "叶问4",
//"nations": "中国香港",
//"score": 7.3,
//"type": "剧情/动作"
//}
func (movie *Movie) Save() int64 {
	result, e := db.Db.Exec("insert into jarvis.t_movie (movie_name, score, movie_type, directors, actors, nations, languages, file_length, description) values (?,?,?,?,?,?,?,?,?);",
		movie.MovieName,
		movie.Score,
		movie.Type,
		movie.Directors,
		movie.Actors,
		movie.Nations,
		movie.Languages,
		movie.FileLength,
		movie.Describe,
	)
	if e != nil {
		log.Println("movie insert err", e.Error())
		return 0
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Println("movie insert id error", err.Error())
		return 0
	}
	return id
}
