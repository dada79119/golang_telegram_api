package mysql

import (
	"errors"
	"telegram_bot_api/util/log"
	"reflect"
	"regexp"
	"strings"
)

func NewDB() Database {
	return new(databaseHandler)
}

func Model(s interface{}) Table {
	return newTable(s, NewDB(), false)
}

func generateTableName(reflectType reflect.Type) string {
	if reflectType == nil {
		log.Error(errors.New("generateTableName reflectType can not nil"))
		return ""
	}

	name := reflectType.Name()
	tableName := ""
	isUppercase := regexp.MustCompile(`^[A-Z]+$`).MatchString
	isNumberCase := regexp.MustCompile(`^[0-9]+$`).MatchString

	for i, r := range name {
		if isUppercase(string(r)) && i != 0 {
			tableName += "_"
		}
		if isNumberCase(string(r)) && i != 0 {
			tableName += "_"
		}

		tableName += strings.ToLower(string(r))
	}

	return tableName
}

func newSchema() Schema {
	schemaHandler := new(SchemaHandler)
	schemaHandler.db = NewDB()
	return schemaHandler
}

func newTable(model interface{}, db Database, isSubQuery bool) Table {
	if model == nil {
		log.Error(errors.New("newTable table can not nil"))
		return nil
	} else if tableName, isString := model.(string); !isString && reflect.TypeOf(model).Kind() != reflect.Ptr  {
		log.Error(errors.New("newTable table kind can only be struct point or string"))
		return nil
	} else {
		table := new(TableHandler)
		table.db = db
		table.isSubQuery = isSubQuery
		table.sqlDml = new(sqlDml)
		table.sqlDql = new(sqlDql)
		table.sqlWhere = new(sqlWhere)

		if !isString {
			reflectType := reflect.TypeOf(model).Elem()
			table.sqlDql.setReflectType(reflectType).
				setReflectValue(reflect.ValueOf(model).Elem())
			table.sqlDml.setReflectType(reflectType).
						setReflectValue(reflect.ValueOf(model).Elem())
			tableName = generateTableName(reflectType)

			tags := parseStructTag(reflectType)
			table.sqlWhere.tags = tags
			table.sqlDql.tags = tags
		}

		table.sqlDml.tableName = tableName
		table.sqlDql.tableName = tableName
		table.sqlWhere.tableName = tableName

		return table
	}
}

func parseStructTag(reflectType reflect.Type) []string {
	var tags []string
	if reflectType != nil {
		for i := 0; i < reflectType.NumField(); i++ {
			field := reflectType.Field(i)
			column := field.Tag.Get("table")
			tags = append(tags, column)
		}
	}
	return tags
}