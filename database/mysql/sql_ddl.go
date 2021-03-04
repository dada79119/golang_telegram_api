package mysql

import (
	"errors"
	"telegram_bot_api/util/log"
	"strconv"
)

type sqlDdl struct {
	tableName    	string
	columns 		[]string
	primaryKeySql 	string
	key 			string
	sql 		    string
}

func (_ddl *sqlDdl) increments(column string) Blueprint {
	if column == "" {
		log.Error(errors.New("increments column can not empty string"))
		return _ddl
	}
	_ddl.columns = append(_ddl.columns, "`" + column + "` int(11) unsigned not null auto_increment")
	_ddl.primaryKeySql = "primary key (` " + column + "`) "
	return _ddl
}

func (_ddl *sqlDdl) tinyint(column string, length int) Blueprint {
	if column == "" {
		log.Error(errors.New("tinyint column can not empty string"))
		return _ddl
	} else if length > 3 {
		log.Error(errors.New("tinyint length max 3"))
		return _ddl
	}

	_ddl.columns = append(_ddl.columns, "`" + column + "` tinyint(" + strconv.Itoa(length) +
		"," + ") not null default '0' ")
	return _ddl
}

func (_ddl *sqlDdl) int(column string, length int) Blueprint {
	if column == "" {
		log.Error(errors.New("varchar column can not empty string"))
		return _ddl
	}

	_ddl.columns = append(_ddl.columns, "`" + column + "` int(" + strconv.Itoa(length) + ") not null ")
	return _ddl
}

func (_ddl *sqlDdl) varchar(column string, length int) Blueprint {
	if column == "" {
		log.Error(errors.New("varchar column can not empty string"))
		return _ddl
	}

	_ddl.columns = append(_ddl.columns, "`" + column + "` varchar(" + strconv.Itoa(length) + ") not null default '' ")
	return _ddl
}

func (_ddl *sqlDdl) decimal(column string, length int, decimal int) Blueprint {
	if column == "" {
		log.Error(errors.New("decimal column can not empty string"))
		return _ddl
	}

	_ddl.columns = append(_ddl.columns, "`" + column + "` decimal(" + strconv.Itoa(length) +
							"," + strconv.Itoa(decimal) + ") not null ")
	return _ddl
}

func (_ddl *sqlDdl) text(column string) Blueprint {
	if column == "" {
		log.Error(errors.New("decimal column can not empty string"))
		return _ddl
	}

	_ddl.columns = append(_ddl.columns, "`" + column + "` text collate utf8mb4_unicode_ci not null ")
	return _ddl
}

func (_ddl *sqlDdl) timestamp(column string) Blueprint {
	if column == "" {
		log.Error(errors.New("decimal column can not empty string"))
		return _ddl
	}

	_ddl.columns = append(_ddl.columns, "`" + column + "` timestamp not null default '0000-00-00 00:00:00' ")
	return _ddl
}

func (_ddl *sqlDdl) comment(text string) Blueprint {
	if len(_ddl.columns) <= 0 || _ddl.columns == nil {
		log.Error(errors.New("comment columns can not nil or empty slice"))
		return _ddl
	}

	_ddl.columns = append(_ddl.columns, _ddl.columns[len( _ddl.columns) - 1] + " comment '" + text + "' ")
	return _ddl
}

func (_ddl *sqlDdl) generateSQL()  {

}

func (_ddl *sqlDdl) generateCreate(sqlWhere *sqlWhere) string {
	log.Sql(_ddl.sql)
	return _ddl.sql
}

func (_ddl *sqlDdl) generateDrop(sqlWhere *sqlWhere) string {
	log.Sql(_ddl.sql)
	return _ddl.sql
}

func (_ddl *sqlDdl) generateAlter(sqlWhere *sqlWhere) string {
	log.Sql(_ddl.sql)
	return _ddl.sql
}