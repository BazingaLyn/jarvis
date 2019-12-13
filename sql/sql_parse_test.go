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
		expr := &sel.Where.Expr
		handlerSelect(expr)
	}

}

func handlerSelect(expr *sqlparser.Expr) {

	switch (*expr).(type) {
	case *sqlparser.AndExpr:
		andExpr := (*expr).(*sqlparser.AndExpr)
		leftExpr := andExpr.Left
		rightExpr := andExpr.Right

		fmt.Println(sqlparser.String(leftExpr))
		fmt.Println(sqlparser.String(rightExpr))
	case *sqlparser.ComparisonExpr:
		comparisonExpr := (*expr).(*sqlparser.AndExpr)
		Left := comparisonExpr.Left
		Right := comparisonExpr.Right
		fmt.Println(sqlparser.String(Left))
		fmt.Println(sqlparser.String(Right))
	}
}
