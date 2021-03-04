package mysql

import (
	"strings"
)

func addPrefixTableName(tableName string, column string, tags []string) string {
	if strings.Index(column, ".") >= 0 {
		return column
	}

	tn := ""
	if len(tags) > 0 {
		for _, v := range tags {
			if v == column {
				tn = "`"+ tableName + "`"+ "."
				break
			}
		}
	} else {
		tn = "`"+ tableName + "`"+ "."
	}

	return tn + column
}
