package web

import (

	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io"
	"strings"
	"linebot/util/log"
)

var (
	//key 需要 16 或 32
	key = "7yaja 7yaja7yaja"
	aesKey = []byte(key)

)

func StrToMd5(str string) (string){
	if str == ""{
		return ""
	}
	data := []byte(str + key)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has)
}

func StrToUnixCrypt(str string) (string){
	sum := sha256.Sum256([]byte(str))
	var builder strings.Builder
	_, err := fmt.Fprintf(&builder, "%x", sum)
	log.Error(err)
	return UnixCrypt(str, builder.String())
}

func KeyEncrypt(cryptoText string) (string, error)  {
	keyBytes := sha256.Sum256([]byte(aesKey))
	return encrypt(keyBytes[:], cryptoText)
}

// encrypt string to base64 crypto using AES
func encrypt(key []byte, text string) (string, error)  {
	plaintext := []byte(text)

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Error(err)
		return "", err
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		log.Error(err)
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func KeyDecrypt(cryptoText string) (string, error) {
	keyBytes := sha256.Sum256([]byte(aesKey))
	return decrypt(keyBytes[:], cryptoText)
}

// decrypt from base64 to decrypted string
func decrypt(key []byte, cryptoText string) (string, error)  {
	ciphertext, err := base64.StdEncoding.DecodeString(cryptoText)
	if err != nil {
		log.Error(err)
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Error(err)
		return "", err
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		log.Verbose("ciphertext too short")
		return "", fmt.Errorf("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)
	return fmt.Sprintf("%s", ciphertext), nil
}

func HashCheck(dbHashPassword string, password string)bool{
	passwordByte := []byte(password)

	// Hashing the password with the default cost of 10
	_, err := bcrypt.GenerateFromPassword(passwordByte, bcrypt.DefaultCost)
	log.Error(err)

	err = bcrypt.CompareHashAndPassword([]byte(dbHashPassword), passwordByte)

	if err == nil{
		return true
	}

	return false
}
