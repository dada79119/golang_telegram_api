package mysql

import (
	"fmt"
	"telegram_bot_api/util/log"
	_string "telegram_bot_api/util/string"
	"reflect"
	"strconv"
	"strings"
)

type sqlDql struct {
	reflectType  	reflect.Type
	reflectValue 	reflect.Value
	joins		    []string
	orderColumns    []string
	groupSql	    string
	havingSql	    string
	limitSql   	    string
	unionSql  	    string
	offsetSql  	    string
	rawLockSql  	string
	selectSql 	    string
	onSql    	    string
	tableName 	 	string
	columns 	    []string
	columns2 	    []string
	questionMarks   []string
	values 		    []interface{}
	sql 		    string
	tags 			[]string
}

func (_dql *sqlDql) setReflectType(t reflect.Type) *sqlDql {
	_dql.reflectType = t
	return _dql
}

func (_dql *sqlDql) setReflectValue(v reflect.Value) *sqlDql {
	_dql.reflectValue = v
	return _dql
}

func (_dql *sqlDql) generateOn(column1 string, operator string, column2 string) {

	_dql.onSql = " on " + _string.TrimQuotes(column1) + " " + _string.TrimQuotes(operator) +
				" " + _string.TrimQuotes(column2) + " "
}

func (_dql *sqlDql) generateJoin(joinTable string, column1 string, operator string, column2 string, status string) {
	_dql.joins = append(_dql.joins, " " + status + " join " + _string.TrimQuotes(joinTable) +
		" on " + addPrefixTableName(_dql.tableName, column1, _dql.tags) + " " + _string.TrimQuotes(operator) + " " +
		addPrefixTableName(_dql.tableName, column2, _dql.tags) + " ")
}

func (_dql *sqlDql) generateJoinWithWhere(joinTable string, closure func(Join), status string) {
	subTable := newTable(joinTable, nil ,true)
	closure(subTable.(Join))
	subTable.getDql().generateSQL(subTable.getWhere())

	_dql.joins = append(_dql.joins, " " + status + " join " +
		joinTable + " " + subTable.getDql().onSql + " " + strings.Join(subTable.getWhere().wheres[:], " ") + " ")

	_dql.values = append(_dql.values, append(subTable.getWhere().values)...)
}

func (_dql *sqlDql) generateUnion(model interface{}) {
	t := model.(Table)
	d := t.getDql()
	d.generateSQL(t.getWhere())
	_dql.unionSql += " union all " + d.sql
}

func (_dql *sqlDql) generateOrderBy(columns []string, value []string) {
	for i := 0 ; i < len(columns) ; i++{
		_dql.orderColumns = append(_dql.orderColumns,columns[i] + " " + _string.TrimQuotes(value[i]))
	}
}

func (_dql *sqlDql) generateGroupBy(column string) {
	_dql.groupSql = " group by " + addPrefixTableName(_dql.tableName, column, _dql.tags)
}

func (_dql *sqlDql) generateHaving(column string, operator string, value interface{}) {
	_dql.havingSql = " having " + addPrefixTableName(_dql.tableName, column, _dql.tags) + " " +
				_string.TrimQuotes(operator) + " ? "

	_dql.values = append(_dql.values, value)
}

func (_dql *sqlDql) generateLimit(offset int, limit int) {
	_dql.limitSql = " limit " + strconv.Itoa(offset) + " , " + strconv.Itoa(limit) + " "
}

func (_dql *sqlDql) generateOffset(offset int) {
	_dql.offsetSql = " offset " + strconv.Itoa(offset) + " "
}

func (_dql *sqlDql) generateRowLock() {
	_dql.rawLockSql = " for update; "
}

func (_dql *sqlDql) setSelectColumn(columns []string) interface{} {
	_dql.columns = columns
	return _dql
}

func (_dql *sqlDql) parseStructTag() {
	if _dql.reflectType != nil {
		for i := 0; i < _dql.reflectType.NumField(); i++ {
			field := _dql.reflectType.Field(i)
			column := field.Tag.Get("table")
			_dql.tags = append(_dql.tags, column)
		}
	}
}

func (_dql *sqlDql) generateSelect() {
	if _dql.selectSql == "" && len(_dql.columns) > 0 {
		_dql.selectSql = "select "
		for _, item := range _dql.columns {
			_dql.selectSql += " " + addPrefixTableName(_dql.tableName, item, _dql.tags) + ","
		}

		_dql.selectSql = _dql.selectSql[0:len(_dql.selectSql) - 1]
	}
}

func (_dql *sqlDql) generateSQL(sqlWhere *sqlWhere) (string, []interface{}) {
	_dql.generateSelect()

	_dql.sql = ""
	if _dql.selectSql != "" {
		_dql.sql += _dql.selectSql + " from " + _dql.tableName
	}

	if _dql.joins != nil {
		_dql.sql += strings.Join(_dql.joins[:], " ")
	}

	if sqlWhere.wheres != nil {
		if _dql.selectSql != "" {
			_dql.sql += " where 1=1 "
		}

		_dql.sql += strings.Join(sqlWhere.wheres[:], " ")
		_dql.values = append(_dql.values, append(sqlWhere.values)...)
	}

	if _dql.groupSql != ""{
		_dql.sql += _dql.groupSql
	}

	if _dql.havingSql != ""{
		_dql.sql += _dql.havingSql
	}

	if len(_dql.orderColumns) > 0 {
		_dql.sql += " order by " + strings.Join(_dql.orderColumns[:], ",")
	}

	if _dql.limitSql != "" {
		_dql.sql += _dql.limitSql
	}

	if _dql.unionSql != "" {
		_dql.sql += _dql.unionSql
	}

	if _dql.rawLockSql != "" {
		_dql.sql += _dql.unionSql
	}

	fmt.Println(_dql.sql)

	log.Sql(_dql.sql)
	log.Sql(_dql.values)
	return _dql.sql, _dql.values
}