package image

import (
	"fmt"
	"strconv"
	"time"
	"linebot/config"
	"linebot/util/file"
)

// 產生圖片上傳路徑
func GeneratePhotoPath(imageType string, transType string) (string, string) {

	// get unix time (nanosecond => millisecond)
	unixTime := time.Now().UnixNano() / int64(time.Millisecond)
	// int64 to string
	strTime := strconv.FormatInt(unixTime, 10)

	saveDir := fmt.Sprintf("%s//%s", config.PathConfig.SaveDirectoryRoot, transType)
	uploadDir := fmt.Sprintf("%s//%s", config.PathConfig.HttpDirectoryRoot, transType)

	if err := file.CreateDirectoryIfNotExist(saveDir); err != nil {
		return "", ""
	}

	return fmt.Sprintf("%s%s.%s", saveDir, strTime, imageType), fmt.Sprintf("%s%s.%s", uploadDir, strTime, imageType)
}

func ReturnPhotoPath(path string) string{
	return  config.ServerConfig.Host +"image/" + path

}
