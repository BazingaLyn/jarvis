package sharding

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/xwb1989/sqlparser"
	"log"
	"testing"
)

type TableRule struct {
	logicTable                     string
	actualDataNodes                string
	databaseShardingStrategyColumn string
	databaseShardingExpression     string
	tableShardingStrategyColumn    string
	tableShardingExpression        string
}

type ShardingDB interface {
	Exec(sql string, args ...interface{}) (sql.Result, error)
}

type GoShardingDB struct {
	tableRules map[string]TableRule
	dbMaps     map[string]*sql.DB
}

func (g GoShardingDB) Exec(sql string, args ...interface{}) (sql.Result, error) {

	stmt, err := sqlparser.Parse(sql)
	if err != nil {
		return nil, err
	}
	switch stmt.(type) {
	case *sqlparser.Select:
		sel := stmt.(*sqlparser.Select)
		tableName := sqlparser.String(sel.From)
		rule, exist := g.tableRules[tableName]
		if !exist {
			log.Fatalf("table %s no config route please check", tableName)
		}
		return handlerRouteColumn(&sel.Where.Expr, rule, tableName)

	case *sqlparser.Update:
		//return handleUpdate(stmt.(*sqlparser.Update))
	case *sqlparser.Insert:
		//return handleInsert(stmt.(*sqlparser.Insert))
	case *sqlparser.Delete:
		//return handleDelete(stmt.(*sqlparser.Delete))
	}
	return nil, nil
}

func handlerRouteColumn(expr *sqlparser.Expr, rule TableRule, tableName string) (sql.Result, error) {

	switch (*expr).(type) {
	case *sqlparser.AndExpr:

	case *sqlparser.OrExpr:
	case *sqlparser.ComparisonExpr:

		comparisonExpr := (*expr).(*sqlparser.ComparisonExpr)
		colName, ok := comparisonExpr.Left.(*sqlparser.ColName)

		if !ok {
			return nil, errors.New("invalid sql expression, the left must be a column name")
		}
		sqlparser.String(colName)

	case *sqlparser.IsExpr:
	case *sqlparser.RangeCond:

	case *sqlparser.ParenExpr:
	case *sqlparser.NotExpr:
	}

	return nil, nil
}

func TestBase0(t *testing.T) {

	tableRule := TableRule{
		logicTable:                     "t_order",
		actualDataNodes:                "ds${0..1}.t_order_${0..2}",
		databaseShardingStrategyColumn: "user_id",
		databaseShardingExpression:     "ds${user_id % 2}",
		tableShardingStrategyColumn:    "user_id",
		tableShardingExpression:        "t_order_${user_id % 3}",
	}

	tableRuleMap := make(map[string]TableRule)
	tableRuleMap[tableRule.logicTable] = tableRule

	goShardingDB := GoShardingDB{
		tableRules: tableRuleMap,
		dbMaps:     nil,
	}

	//fmt.Println("goShardingDB type:", reflect.TypeOf(goShardingDB))
	//goShardingDB1 := new(GoShardingDB)
	//
	//fmt.Println("goShardingDB1 type:", reflect.TypeOf(goShardingDB1))
	goShardingDB.Exec("select * from t_order where user_id = ?", 3)

}

func TestBase1(t *testing.T) {
	dbMap := make(map[string]*sql.DB)

	db0, err0 := sql.Open("mysql", "root:1qaz2wsx!@tcp(49.235.67.172:3306)/ds0")
	if err0 != nil {
		log.Panicf("connect mysql failed %s", err0)
	}

	db1, err1 := sql.Open("mysql", "root:1qaz2wsx!@tcp(49.235.67.172:3306)/ds1")
	if err1 != nil {
		log.Panicf("connect mysql failed %s", err1)
	}

	dbMap["db0"] = db0
	dbMap["db1"] = db1

}

//func TestExec(t *testing.T) {
//
//	db, err := Open("mysql", "root:1qaz2wsx!@tcp(49.235.67.172:3306)/jarvis")
//	if err == nil {
//		result, err := db.ShardingExec("insert into t_user (id,name,age) values (?,?,?)", 2, "bazinga", 23)
//		if err != nil {
//			log.Panicln("movie insert error", err.Error())
//		}
//		id, err := result.LastInsertId()
//		if err != nil {
//			log.Panicln("movie insert id error", err.Error())
//		}
//		log.Println(id)
//
//	}
//
//}
