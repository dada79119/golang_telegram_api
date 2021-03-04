package mysql

type SchemaHandler struct {
	db  	   Database
	//sqlDdl     *sqlDdl
}

func (_schema *SchemaHandler) Create(tableName string, closure func(Blueprint)) {
	// ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COLLATE=utf8mb4_unicode_ci

	//blueprint := newBlueprint(tableName)
	//closure(blueprint)
	//blueprint.(*BlueprintHandler).generateSQL()
}