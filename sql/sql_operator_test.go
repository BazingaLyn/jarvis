package sql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"testing"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:1qaz2wsx!@tcp(49.235.67.172:3306)/jarvis")

	if err != nil {
		log.Panicln("err:", err.Error())
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)
}

// exec方式插入
func TestExecInsert(t *testing.T) {

	result, err := db.Exec("insert t_user (name,age) values (?,?)", "lyn", 30)
	if err != nil {
		log.Panicln("err:", err.Error())
	}
	i, _ := result.LastInsertId()
	log.Printf("id is %d", i)

}

func TestPrepareInsert(t *testing.T) {

	stmt, err := db.Prepare("insert t_user (name,age) values (?,?)")
	if err != nil {
		log.Panicln("err:", err.Error())
	}

	defer stmt.Close()

	result, _ := stmt.Exec("ted", 31)
	i, _ := result.LastInsertId()
	log.Printf("id is %d", i)

}

func TestPrepareBatchInsert(t *testing.T) {

	war3players := []struct {
		name string
		age  int
	}{
		{"happy", 27},
		{"err0", 21},
	}

	stmt, err := db.Prepare("insert t_user (name,age) values (?,?)")
	if err != nil {
		log.Panicln("err:", err.Error())
	}

	defer stmt.Close()

	for _, war3player := range war3players {
		result, _ := stmt.Exec(war3player.name, war3player.age)
		i, _ := result.LastInsertId()
		log.Printf("id is %d", i)
	}

}

func TestExecUpdate(t *testing.T) {
	result, err := db.Exec("update t_user set age = age+1 where name = ?", "ted")
	if err != nil {
		log.Panicln("err:", err.Error())
	}
	i, _ := result.RowsAffected()
	log.Printf("affect row is %d", i)
}

func TestSelectOneRow(t *testing.T) {
	var name string
	var age int
	db.QueryRow("select name,age from t_user where id = ?", 3).Scan(&name, &age)

	fmt.Printf("username is %s age is %d", name, age)
}

func TestSelectMultiRow(t *testing.T) {
	rows, e := db.Query("select * from t_user")
	if e != nil {
		log.Panicln("err:", e.Error())
	}

	for rows.Next() {
		var (
			id   int
			name string
			age  int
		)
		if err := rows.Scan(&id, &name, &age); err != nil {
			log.Fatal(err)
		}

		log.Printf("id %d name is %s age is %d", id, name, age)

	}
}

func TestDbState(t *testing.T) {
	stats := db.Stats()
	log.Println(stats.OpenConnections)
	log.Println(stats.InUse)
	log.Println(stats.Idle)
}
