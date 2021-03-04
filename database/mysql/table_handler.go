package mysql

import (
	"database/sql"
	"errors"
	"telegram_bot_api/util/log"
	"reflect"
	"strings"
)

type TableHandler struct {
	db  	   Database
	sqlDml     *sqlDml
	sqlDql     *sqlDql
	sqlWhere   *sqlWhere
	isSubQuery bool
	values 	   []interface{}
	sql 	   string
}

func (_table *TableHandler) getDml() *sqlDml {
	return _table.sqlDml
}

func (_table *TableHandler) getDql() *sqlDql {
	return _table.sqlDql
}

func (_table *TableHandler) getWhere() *sqlWhere {
	return _table.sqlWhere
}

func (_table *TableHandler) On(column1 string, operator string, column2 string) Table {
	if column1 == "" || operator == ""  || column2 == "" {
		log.Error(errors.New("on column1 and operator and column2 can not empty"))
		return nil
	}
	_table.sqlDql.generateOn(column1, operator, column2)
	return _table
}

func (_table *TableHandler) InnerJoin(joinTable string, column1 string, operator string, column2 string) Table {
	if column1 == "" || operator == "" || column2 == "" {
		log.Error(errors.New("innerJoin column1 and operator and column2 can not empty string"))
		return nil
	}
	_table.sqlDql.generateJoin(joinTable, column1, operator, column2, "inner")
	return _table
}

func (_table *TableHandler)LeftJoin(joinTable string, column1 string, operator string, column2 string) Table {
	if joinTable == "" || column1 == "" || operator == "" || column2 == "" {
		log.Error(errors.New("leftJoin joinTable and column1 and operator and column2 can not empty string"))
		return nil
	}
	_table.sqlDql.generateJoin(joinTable, column1, operator, column2, "left")
	return _table
}

func (_table *TableHandler) InnerJoinWithWhere(joinTable string, closure func(Join)) Table {
	if joinTable == "" || closure == nil {
		log.Error(errors.New("innerJoinWithWhere joinTable can not empty string, join can not nil"))
		return nil
	}
	_table.sqlDql.generateJoinWithWhere(joinTable, closure, "inner")
	return _table
}

func (_table *TableHandler) LeftJoinWithWhere(joinTable string, closure func(Join)) Table {
	if joinTable == "" || closure == nil {
		log.Error(errors.New("leftJoinWithWhere joinTable can not empty string, join can not nil"))
		return nil
	}
	_table.sqlDql.generateJoinWithWhere(joinTable, closure, "left")
	return _table
}

func (_table *TableHandler) Where(values ...interface{}) Table {
	if len(values) != 1 && len(values) != 3 {
		log.Error(errors.New("func where pass paramter count join can 1 and 3"))
		return nil
	}

	if len(values) == 1 {
		if value, ok := values[0].(func(Table)); ok {
			_table.sqlWhere.generateWhere(_table.isSubQuery, "and", "", "", value)
		} else {
			log.Error(errors.New("func where first paramter join can func(Table)"))
			return nil
		}
	}

	if len(values) == 3 {
		if column, ok := values[0].(string); !ok {
			log.Error(errors.New("func where first paramter join can string"))
			return nil
		} else if operator, ok := values[1].(string); !ok {
			log.Error(errors.New("func where second paramter join can string"))
			return nil
		} else if value := values[2]; value == nil {
			log.Error(errors.New("func where third paramter can not nil"))
			return nil
		} else {
			_table.sqlWhere.generateWhere(_table.isSubQuery, "and", column, operator, value)
		}
	}

	return _table
}

func (_table *TableHandler) WhereIn(boolean string, column string, value interface{}) Table {
	boolean = strings.ToLower(boolean)
	if boolean != "or" && boolean != "and" {
		log.Error(errors.New("whereIn boolean can only be 'or' or 'and' "))
		return nil
	}

	if column == "" || value == nil {
		log.Error(errors.New("whereIn column can not empty string, value can not nil"))
		return nil
	}
	_table.sqlWhere.generateWhereIn(_table.isSubQuery, boolean, false, column, value)
	return _table
}

func (_table *TableHandler) WhereNotIn(boolean string, column string, value interface{}) Table {
	boolean = strings.ToLower(boolean)
	if boolean != "or" && boolean != "and" {
		log.Error(errors.New("whereNotIn boolean can only be 'or' or 'and' "))
		return nil
	}

	if column == "" || value == nil {
		log.Error(errors.New("whereNotIn column can not empty string, value can not nil"))
		return nil
	}
	_table.sqlWhere.generateWhereIn(_table.isSubQuery, boolean, true, column, value)
	return _table
}

func (_table *TableHandler) WhereOr(values ...interface{}) Table {
	if len(values) != 1 && len(values) != 3 {
		log.Error(errors.New("func where pass paramter count join can 1 and 3"))
		return nil
	}

	if len(values) == 1 {
		if value, ok := values[0].(func(Table)); ok {
			_table.sqlWhere.generateWhere(_table.isSubQuery, "or", "", "", value)
		} else {
			log.Error(errors.New("func where first paramter join can func(Table)"))
			return nil
		}
	}

	if len(values) == 3 {
		if column, ok := values[0].(string); !ok {
			log.Error(errors.New("func where first paramter join can string"))
			return nil
		} else if operator, ok := values[1].(string); !ok {
			log.Error(errors.New("func where second paramter join can string"))
			return nil
		} else if value := values[2]; value == nil {
			log.Error(errors.New("func where third paramter can not nil"))
			return nil
		} else {
			_table.sqlWhere.generateWhere(_table.isSubQuery, "or", column, operator, value)
		}
	}

	return _table
}

func (_table *TableHandler) WhereRaw(sql string) Table {
	if sql == "" {
		log.Error(errors.New("whereRaw sql can not empty string"))
		return nil
	}
	_table.sqlWhere.generateWhereRaw(sql)
	return _table
}

func (_table *TableHandler) WhereBetween(boolean string, column string, values []interface{}) Table {
	boolean = strings.ToLower(boolean)
	if boolean != "or" && boolean != "and" {
		log.Error(errors.New("whereBetween boolean can only be 'or' or 'and' "))
		return nil
	}

	if column == "" || values == nil || len(values) != 2 {
		log.Error(errors.New("whereBetween column can not empty string, value can not nil"))
		return nil
	}
	_table.sqlWhere.generateWhereBetween(false, boolean, column, values)
	return _table
}

func (_table *TableHandler) WhereNotBetween(boolean string, column string, values []interface{}) Table {
	boolean = strings.ToLower(boolean)
	if boolean != "or" && boolean != "and" {
		log.Error(errors.New("whereNotBetween boolean can only be 'or' or 'and' "))
		return nil
	}

	if column == "" || values == nil || len(values) != 2 {
		log.Error(errors.New("whereNotBetween column can not empty string, value can not nil"))
		return nil
	}
	_table.sqlWhere.generateWhereBetween(true, boolean, column, values)
	return _table
}

func (_table *TableHandler) WhereColumne(boolean string, column1 string, operator string, column2 string) Table {
	if column1 == "" || operator == "" || column2 == "" {
		log.Error(errors.New("whereColumne column1 can and operator and column2 not empty string"))
		return nil
	}
	_table.sqlWhere.generateWhereColumne(boolean, column1, operator, column2)
	return _table
}

func (_table *TableHandler) WhereNull(column string) Table {
	if column == "" {
		log.Error(errors.New("whereNull column can not empty string"))
		return nil
	}
	_table.sqlWhere.generateWhereNull(false, column)
	return _table
}

func (_table *TableHandler) WhereNotNull(column string) Table {
	if column == "" {
		log.Error(errors.New("whereNotNull column can not empty string"))
		return nil
	}
	_table.sqlWhere.generateWhereNull(true, column)
	return _table
}

func (_table *TableHandler) Union(model interface{}) Table {
	if reflect.TypeOf(model).Kind() != reflect.Ptr {
		log.Error(errors.New("union model can not nil"))
		return nil
	}
	_table.sqlDql.generateUnion(model)
	return _table
}

func (_table *TableHandler) Select(columns []string) Table {
	if columns == nil {
		log.Error(errors.New("select columns can not nil"))
		return nil
	}

	_table.sqlDql.setSelectColumn(columns)
	return _table
}

func (_table *TableHandler) Form(tableName string) Table {
	_table.sqlDml.tableName = tableName
	_table.sqlDql.tableName = tableName
	_table.sqlWhere.tableName = tableName
	return _table
}

func (_table *TableHandler) OrderBy(columns []string, value []string) Table {
	_table.sqlDql.generateOrderBy(columns, value)
	return _table
}

func (_table *TableHandler) GroupBy(column string) Table {
	if column == "" {
		log.Error(errors.New("groupBy column can not empty string"))
		return nil
	}
	_table.sqlDql.generateGroupBy(column)
	return _table
}

func (_table *TableHandler) Having(column string, operator string, value interface{}) Table {
	if column == "" || operator == "" || value == nil {
		log.Error(errors.New("having column and operator can not empty string, value can not nil"))
		return nil
	}

	_table.sqlDql.generateHaving(column, operator, value)

	return _table
}

func (_table *TableHandler) Limit(offset int, limit int) Table {
	_table.sqlDql.generateLimit(offset, limit)
	return _table
}

func (_table *TableHandler) Offset(index int) Table {
	_table.sqlDql.generateOffset(index)
	return _table
}

func (_table *TableHandler) RowLock() Table {
	_table.sqlDql.generateRowLock()
	return _table
}

func (_table *TableHandler) exec() (sql.Result, error) {
	db := _table.db

	stmt, err := db.getTx().Prepare(_table.sqlDml.sql)
	defer _table.db.CloseStmt(stmt)
	if err != nil {
		log.Error(db.Rollback())
		return nil, err
	} else if res, err := stmt.Exec(_table.sqlDml.values...); err != nil {
		log.Error(db.Rollback())
		return nil, err
	} else {
		return res, err
	}
}

func (_table *TableHandler) execSQL() (int64, error) {
	defer _table.db.CloseDb()
	if err := _table.db.connect(); err != nil {
		return 0, err
	}

	isTransaction := _table.db.getTx() == nil
	if isTransaction {
		if err := _table.db.beginTransaction(); err != nil {
			return 0, err
		}
	}

	if res, err := _table.exec(); err != nil {
		return 0, err
	} else if insertId, err := res.LastInsertId(); err != nil {
		log.Error(_table.db.Rollback())
		return 0, err
	} else if !isTransaction {
		return insertId, nil
	} else if err := _table.db.Commit(); err != nil {
		log.Error(_table.db.Rollback())
		return 0, err
	} else {
		return insertId, nil
	}
}

func (_table *TableHandler) Insert() (int64, error) {
	_table.sql, _table.values = _table.sqlDml.generateInsertSQL()
	return _table.execSQL()
}

func (_table *TableHandler)Update(columns []string) error {
	if columns == nil {
		return errors.New("update columns can not nil")
	}

	_table.sql, _table.values = _table.sqlDml.generateUpdateSQL(columns, _table.sqlWhere)

	_, err := _table.execSQL()
	return err
}

func (_table *TableHandler) Delete() error {
	_table.sql, _table.values = _table.sqlDml.generateDeleteSQL(_table.sqlWhere)
	_, err := _table.execSQL()
	return err
}

func (_table *TableHandler) Get(closure func(*sql.Rows)) Table {
	if closure == nil {
		log.Error(errors.New("get closure can not nil"))
		return nil
	} else {
		_table.sql, _table.values = _table.sqlDql.generateSQL(_table.sqlWhere)

		defer _table.db.CloseDb()
		if err := _table.db.connect(); err != nil {
			log.Error(err)
			return nil
		}

		stmt, err := _table.db.getDb().Prepare(_table.sqlDql.sql)
		defer _table.db.CloseStmt(stmt)
		if err != nil {
			log.Error(err)
			return nil
		}

		rows, err := stmt.Query(_table.sqlDql.values...)
		defer _table.db.CloseRows(rows)
		if err != nil {
			log.Error(err)
			return nil
		}

		for rows.Next() {
			closure(rows)
		}

		return _table
	}
}

func (_table *TableHandler) Find() *sql.Row {
	_table.sql, _table.values = _table.sqlDql.generateSQL(_table.sqlWhere)

	defer _table.db.CloseDb()
	if err := _table.db.connect(); err != nil {
		log.Error(err)
		return nil
	}

	stmt, err := _table.db.getDb().Prepare(_table.sqlDql.sql)
	defer _table.db.CloseStmt(stmt)

	if err != nil{
		log.Error(err)
		return nil
	}

	return stmt.QueryRow(_table.sqlDql.values...)
}

func (_table *TableHandler) Count() int {
	_table.sqlDql.selectSql = "select count(*)"
	_table.sql, _table.values = _table.sqlDql.generateSQL(_table.sqlWhere)

	defer _table.db.CloseDb()
	if err := _table.db.connect(); err != nil {
		log.Error(err)
		return 0
	}

	var c int

	stmt, err := _table.db.getDb().Prepare(_table.sqlDql.sql)

	if err != nil{
		log.Error(err)
		return 0
	}

	err = stmt.QueryRow(_table.sqlDql.values...).Scan(&c)
	defer _table.db.CloseStmt(stmt)

	if err != nil{
		log.Error(err)
		return 0
	}

	return c
}

func (_table *TableHandler) SQL() (string, []interface{}) {
	return _table.sql, _table.values
}