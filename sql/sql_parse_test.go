package sql

import (
	"fmt"
	"log"
	"testing"
)
import "github.com/xwb1989/sqlparser"

func TestSQLParse(t *testing.T) {
	sql := "select * from t_user where a = 'abc'"
	stmt, err := sqlparser.Parse(sql)

	if err != nil {
		log.Panicf("err  is %s", err.Error())
	}

	switch stmt.(type) {
	case *sqlparser.Select:
		sel := stmt.(*sqlparser.Select)
		fmt.Println(sqlparser.String(sel.From))
		fmt.Println(sel.Where.Expr)

	}

}
