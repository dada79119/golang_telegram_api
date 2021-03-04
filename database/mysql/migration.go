package mysql

import (
	"bufio"
	"io/ioutil"
	"telegram_bot_api/util"
	"telegram_bot_api/util/file"
	"telegram_bot_api/util/log"
	"os"
	"strings"
)

type Migration interface {
	Up()
	Down()
}

func deleteMigration(err error, migrationFileName string) {
	log.Error(err)
	log.Error(os.Remove(migrationFileName))
}

func CreateMigration() {
	filePath := file.GetCurDirPath() + "/database/migration/"
	timeDate := util.TimeNow().Format("20060101150405")
	name := os.Args[2]
	fileName := timeDate + "_" + name
	migrationFileName := filePath + fileName + ".go"
	structName := strings.Replace(name, "_", "", -1) + timeDate
	structName = strings.ToUpper(structName[0:1]) + structName[1:]
	migrationFactoryPath := file.GetCurDirPath() + "/database/mysql/migration_factory.go"

	var (
		f *os.File
		fileRead *os.File
		fileWrite *os.File
	)

	if f, err := os.OpenFile(migrationFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666); err != nil {
		log.Error(err)
		return
	} else if _, err = f.WriteString("package migration\n\ntype " + structName + " struct {}\n\n" +
										"func (t *" + structName + ") Up() {\n}\n\n" +
										"func (t *" + structName + ") Down() {\n}"); err != nil {
		deleteMigration(err, migrationFileName)
		return
	} else if fileRead, err = os.OpenFile(migrationFactoryPath, os.O_RDWR, 0666); err != nil {
		deleteMigration(err, migrationFileName)
		return
	} else if b, err := ioutil.ReadAll(fileRead); err != nil {
		deleteMigration(err, migrationFileName)
		return
	} else if fileWrite, err := os.OpenFile(migrationFactoryPath, os.O_RDWR|os.O_TRUNC, 0666); err != nil {
		deleteMigration(err, migrationFileName)
		return
	} else {
		newPgf := ""
		line := 1
		scanner := bufio.NewScanner(strings.NewReader(string(b)))
		for scanner.Scan() {
			if line == 3 && strings.Index(scanner.Text(), "import") == -1{
				newPgf += "import \"like_life_linebot/migration\"\n"
			} else if scanner.Text() == "}" {
				newPgf += "    \"" + fileName + "\": &migration." + structName + "{},\n"
			}
			newPgf += scanner.Text() + "\n"
			line++
		}

		_, err = fileWrite.WriteString(newPgf)
		if err != nil {
			deleteMigration(err, migrationFileName)
		}
	}

	defer closeFile(f)
	defer closeFile(fileRead)
	defer closeFile(fileWrite)
}

func Migrate() {
	createMigrationTable()
	//filePath := file.GetCurDirPath() + "/database/migration/"
//files := file.GetAllFilesInDirectory(filePath)
//sort.Strings(files)

//var typeRegistry = make(map[string]reflect.Type)
//myTypes := []interface{}{"Ggffgdgf20190909174522"}
//
//reflect.TypeOf("Ggffgdgf20190909174522")

//v := reflect.New(typeRegistry["migration.Ggffgdgf20190909174522"]).Elem()
//fmt.Println(v.Kind())
}

func closeFile(f *os.File){
	if f != nil {
		_ = f.Close()
	}
}

func createMigrationTable() {
	NewDB().Transaction(func(database Database) {
		if _, err := database.getTx().Exec("create table `migrations` (" +
			"`id` int(11) not null auto_increment," +
			"`migration` varchar(250) collate utf8mb4_unicode_ci not null default ''," +
			"`batch` int(11) not null default '0'," +
			"`created_at` timestamp not null default current_timestamp," +
			"primary key (`id`) engine=InnoDB default charset=utf8mb4 collate=utf8mb4_unicode_ci;"); err != nil {
			log.Error(err)
			log.Error(database.getTx().Rollback())
			return
		} else {
			log.Error(database.getTx().Commit())
		}
	})
}