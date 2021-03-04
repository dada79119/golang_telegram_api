package crypt

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"
)

// 生成32位MD5
func GetMD5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

// return len=8  salt
func GetRandomSalt() string {
	return GetRandomString(8)
}

//生成随机字符串
func GetRandomString(lenght int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < lenght; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

//產生隨機6碼驗證(大小寫英文+數字)
func GenerateRandomCode(lenght int) string{
	chars := [62]string{
		"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k",
		"l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v",
		"w", "x", "y", "z",
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K",
		"L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V",
		"W", "X", "Y", "Z",
	}

	// 產生隨機6碼驗證
	rand.Seed(time.Now().UnixNano())
	code := ""
	for i := 0; i < lenght; i++ {
		randInt := rand.Intn(len(chars))
		code += chars[randInt]
	}

	return code
}