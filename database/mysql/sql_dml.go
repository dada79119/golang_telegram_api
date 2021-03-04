package mysql

import (
	"telegram_bot_api/util/log"
	"reflect"
	"strings"
)

type sqlDml struct {
	reflectType  	reflect.Type
	reflectValue 	reflect.Value
	joins		    []string
	orderColumns    []string
	groupSql	    string
	havingSql	    string
	limitSql   	    string
	unionSql  	    string
	offsetSql  	    string
	selectSql 	    string
	onSql    	    string
	tableName    	string
	columns 	    []string
	questionMarks   []string
	values 		    []interface{}
	sql 		    string
}

func (_dml *sqlDml) setReflectType(t reflect.Type) *sqlDml {
	_dml.reflectType = t
	return _dml
}

func (_dml *sqlDml) setReflectValue(v reflect.Value) *sqlDml {
	_dml.reflectValue = v
	return _dml
}

func (_dml *sqlDml) clearArray() {
	_dml.columns = _dml.columns[:0]
	_dml.questionMarks = _dml.questionMarks[:0]
	_dml.values = _dml.values[:0]
}

func (_dml *sqlDml) parseStructTag(closure func(string, interface{})) {
	for i := 0; i < _dml.reflectType.NumField(); i++ {
		field := _dml.reflectType.Field(i)
		column := field.Tag.Get("table")
		if column == "" {
			continue
		}
		value :=  _dml.reflectValue.Field(i).Interface()
		closure(column, value)
	}
}

func (_dml *sqlDml) generateInsertSQL() (string, []interface{}) {
	_dml.clearArray()

	_dml.parseStructTag(func(column string, value interface{}) {
		_dml.columns = append(_dml.columns, column)
		_dml.questionMarks = append(_dml.questionMarks, "?")
		_dml.values = append(_dml.values, value)
	})

	_dml.sql = "insert into " + _dml.tableName +
		" (" + strings.Join(_dml.columns[:], ",") + ")values(" +
		strings.Join(_dml.questionMarks[:], ",") + ")"

	log.Sql(_dml.sql)
	log.Sql(_dml.values)
	return _dml.sql, _dml.values
}

func (_dml *sqlDml) generateUpdateSQL(columns []string, sqlWhere *sqlWhere) (string, []interface{}) {
	_dml.clearArray()

	_dml.parseStructTag( func(column string, value interface{}) {
		for _, columnName := range columns {
			if column == columnName{
				_dml.columns = append(_dml.columns, column + " = ?")
				_dml.values = append(_dml.values, value)
				continue
			}
		}
	})

	_dml.sql = "update " + _dml.tableName + " set " + strings.Join(_dml.columns[:], ",")

	if sqlWhere.wheres != nil {
		_dml.sql += " where 1=1 " + strings.Join(sqlWhere.wheres[:], " ")
		_dml.values = append(_dml.values, append(sqlWhere.values)...)
	}

	log.Sql(_dml.sql)
	log.Sql(_dml.values)
	return _dml.sql, _dml.values
}

func (_dml *sqlDml) generateDeleteSQL(sqlWhere *sqlWhere) (string, []interface{}) {
	_dml.clearArray()
	_dml.sql = "delete from " + _dml.tableName

	if sqlWhere.wheres != nil {
		_dml.sql += " where 1=1 " + strings.Join(sqlWhere.wheres[:], " ")
		_dml.values = append(_dml.values, append(sqlWhere.values)...)
	}

	log.Sql(_dml.sql)
	log.Sql(_dml.values)
	return _dml.sql, _dml.values
}