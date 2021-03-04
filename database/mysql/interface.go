package mysql

import (
	"database/sql"
)

type Database interface {
	connect() error
	getDb() *sql.DB
	Option(closure func(Database))
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row

	/*
	ex: mysql.NewDB().Transaction(func(database mysql.Database) {
			var user User
			_,err := database.Model(&user).Insert()

			if err != nil {
				log.Error(err)
				log.Error(database.Rollback())
			} else {
				log.Error(database.Commit())
			}
		 })
	*/
	Transaction(closure func(Database))

	beginTransaction() error
	Commit() error
	Rollback() error

	Model(s interface{}) Table
	getTx() *sql.Tx

	CloseDb()
	CloseRows(rows *sql.Rows)
	CloseStmt(stmt *sql.Stmt)
}

type Join interface {
	/* on("user.id", "=", "coin.user_id") */
	On(column1 string, operator string, column2 string) Table
}

type Table interface {
	/* ex: InnerJoin("wallet", "wallet.user_id", "=", "user.id") */
	InnerJoin(joinTable string, column1 string, operator string, column2 string) Table

	/* ex: LeftJoin("wallet", "wallet.user_id", "=", "user.id") */
	LeftJoin(joinTable string, column1 string, operator string, column2 string) Table

	/*
		ex: InnerJoinWithWhere("coin", func(join mysql.Join) {
				join.On("coin.id", "=", "coin.id").Where("name","=","fffff")
			})

			LeftJoinWithWhere("coin", func(join mysql.Join) {
				join.On("coin.id", "=", "coin.id").Where("name","=","fffff")
			})
	*/
	InnerJoinWithWhere(joinTable string, closure func(Join)) Table
	LeftJoinWithWhere(joinTable string, closure func(Join)) Table

	/*
		ex: Where("user.id", "=", 1) or
			Where("name", "like", "%aaa%")
			Where("user.id", "=", func(table mysql.Table) {
				table.Select([]string{"id"}).Form("coin").Where("name", "like", "USDT")
			})
			Where(func(table Table) {
				table.Where("user_coupon.STATUS", "=", "0").
					WhereOr("coupon_list.end_at", "<", "2019-11-06")
			})
	*/
	Where(values ...interface{}) Table

	/*
		   ex: WhereIn("and", true, "user.id", []interface{}{1,2,3,4})
		       WhereIn("or", true, "name", []interface{}{"a","b"})
			   WhereIn("or", true, "name", func(table mysql.Table) {
					table.Select([]string{"id"}).Form("coin").Where("name", "like", "USDT")
			   })
	*/
	WhereIn(boolean string, column string, values interface{}) Table
	WhereNotIn(boolean string, column string, values interface{}) Table

	/*
		ex: WhereOr("user.id", "=", 1)
		    WhereOr("name", "like", "%aaa%")
		    WhereOr("user.id", "=", func(table mysql.Table) {
				table.Select([]string{"id"}).Form("coin").Where("name", "like", "USDT")
		    })
	*/
	WhereOr(values ...interface{}) Table

	/*
		ex: WhereOr("and user.id = order.user_id")
	*/
	WhereRaw(sql string) Table
	WhereBetween(boolean string, column string, values []interface{}) Table
	WhereNotBetween(boolean string, column string, values []interface{}) Table

	/*
		ex: WhereColumne("and", "a", "||","(cccc = bbb)")
	*/
	WhereColumne(boolean string, column1 string, operator string, column2 string) Table

	WhereNull(column string) Table
	WhereNotNull(column string) Table

	OrderBy(columns []string, value []string) Table
	GroupBy(column string) Table

	/*
		ex: WhereOr("sum(xxx) ", ">", 1)
	*/
	Having(column string, operator string, value interface{}) Table

	Limit(offset int, limit int) Table

	Form(tableName string) Table
	Select(columns []string) Table

	/*
	ex:
	 first := mysql.Model(model).Select([]string{"amount"}).Where("user_id","=", model.UserId)
	 mysql.Model(model)).Where("user_id","=", model.UserId).Union(first).Get(func(rows *sql.Rows){})
	*/
	Union(model interface{}) Table
	RowLock() Table

	Insert() (int64, error)
	Update(columns []string)error
	Delete() error
	Find() *sql.Row
	Get(closure func(*sql.Rows)) Table

	getDml() *sqlDml
	getDql() *sqlDql
	getWhere() *sqlWhere
	exec() (sql.Result, error)
	execSQL() (int64, error)
	SQL() (string, []interface{})
	Count() int
}

type Schema interface {
	Create(tableName string, closure func(Blueprint))
}

type Blueprint interface {
	increments(column string) Blueprint

	tinyint(column string, length int) Blueprint
	int(column string, length int) Blueprint
	varchar(column string, length int) Blueprint
	decimal(column string, length int, decimal int) Blueprint
	text(column string) Blueprint
	timestamp(column string) Blueprint

	comment(text string) Blueprint
}
