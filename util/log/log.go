package log

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"runtime"
)

func Verbose(text string){
	_, fn, line, _ := runtime.Caller(1)
	fmt.Printf("[Verbose] %s ,line %d: %s", fn, line, text)
}

func Error(err error){
	if err != nil {
		_, fn, line, _ := runtime.Caller(1)
		fmt.Printf("[ERROR] %s ,line %d: %s \n", fn, line, err.Error())
	}
}

func Sql(text interface{}){
	if gin.Mode() != gin.DebugMode {
		_, fn, line, _ := runtime.Caller(1)
		fmt.Printf("[SQL] %s ,line %d: %v \n", fn, line, text)
	}
}