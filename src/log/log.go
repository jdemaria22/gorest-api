package logging

import (
	"fmt"
	"time"
)

const (
	LOG_NAME    = "[GOREST-API]"
	FORMAT_DATE = "02-01-2006-15:04:05"
)

func Info(s interface{}) {
	fmt.Println(fmt.Sprintf("%s %s %s %s", LOG_NAME, "[INFO]", time.Now().UTC().Format(FORMAT_DATE), s))
}

func Debug(s interface{}) {
	fmt.Println(fmt.Sprintf("%s %s %s %s", LOG_NAME, "[DEBUG]", time.Now().UTC().Format(FORMAT_DATE), s))
}

func Error(s interface{}) {
	fmt.Println(fmt.Sprintf("%s %s %s %s", LOG_NAME, "[ERROR]", time.Now().UTC().Format(FORMAT_DATE), s))
}
