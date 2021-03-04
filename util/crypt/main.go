package crypt

import (
	"linebot/util/crypt/web"
)

func StrToMd5(str string) (string){
	return web.StrToMd5(str)
}

func StrToUnixCrypt(str string) (string){
	return web.StrToUnixCrypt(str)
}

func KeyEncrypt(cryptoText string) (string, error) {
	return  web.KeyEncrypt(cryptoText)
}

func KeyDecrypt(cryptoText string) (string, error) {
	return  web.KeyDecrypt(cryptoText)
}