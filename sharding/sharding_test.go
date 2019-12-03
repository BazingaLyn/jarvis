package sharding

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"testing"
)

func TestExec(t *testing.T) {

	db, err := Open("mysql", "root:1qaz2wsx!@tcp(49.235.67.172:3306)/jarvis")
	if err == nil {
		result, err := db.ShardingExec("insert into t_user (id,name,age) values (?,?,?)", 2, "bazinga", 23)
		if err != nil {
			log.Panicln("movie insert error", err.Error())
		}
		id, err := result.LastInsertId()
		if err != nil {
			log.Panicln("movie insert id error", err.Error())
		}
		log.Println(id)

	}

}
