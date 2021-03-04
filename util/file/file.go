package file

import (
	"io/ioutil"
	"telegram_bot_api/util/log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

/*
	判斷路徑/文件是否存在
*/
func IsFileExists(path string) bool {
	if path == ""{
		return false
	}

	//----取得文件訊息
	_, err := os.Stat(path)
	//----表示沒有找到
	if err != nil {
		log.Error(err)
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

/*
	判斷路徑是否為文件夾
*/
func IsDirectory(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		log.Error(err)
		return false
	}
	return s.IsDir()
}

/*
	判斷資料夾路徑是否存在，如果不存在則創建資料夾，並給予權限 0777
*/
func CreateDirectoryIfNotExist(directory string) error{
	if exist := IsFileExists(directory); !exist{
		return os.MkdirAll(directory, 0777)
	}
	return nil
}

/*
	刪除指定資料夾
*/
func DeleteDirectory(directory string) {
	if IsFileExists(directory) {
		os.RemoveAll(directory)
	}
}

/*
	取得資料夾內所有的檔案
*/
func GetAllFilesInDirectory(searchDirectory string) []string{
	paths := make([]string, 0)
	files, _ := ioutil.ReadDir(searchDirectory)

	for _, file := range files {
		if file.IsDir() || file.Name() == ".DS_Store" {
			continue
		}else{
			paths = append(paths, file.Name())
		}
	}

	return paths
}

/*
	取得資料夾內指定數量的檔案
*/
func GetNFilesInDirectory(searchDirectory string, targetDirectory string, fileCount int) []string{
	paths := make([]string, 0)
	files, _ := ioutil.ReadDir(searchDirectory)

	for _, file := range files {
		if file.IsDir(){
			continue
		}else{
			paths = append(paths,targetDirectory + "//" + file.Name())
			if len(paths) >= fileCount{
				break
			}
		}
	}

	return paths
}


/*
	取得當前執行檔路徑
*/
func GetCurDirPath() (rst string) {
	var osType = runtime.GOOS
	if osType == "linux"{
		// os.Args 提供原始命令行参数访问功能。注意，切片中的第一个参数是该程序的路径，并且 os.Args[1:]保存所有程序的的参数
		file, _ := exec.LookPath(os.Args[0])
		path, _ := filepath.Abs(file)
		rst = filepath.Dir(path)
	}else{
		rst, _ = os.Getwd()
	}
	return strings.Replace(rst, "\\", "/", -1)
}