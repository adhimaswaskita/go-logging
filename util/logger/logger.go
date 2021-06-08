package logger

import (
	"fmt"
	"io"
	"log"
	"os"
)

//CreateLogFile initiate the file for saving error log
func CreateLogFile(filename string) io.Writer {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}

	return file
}

//SetErrorLog save log to initiated file from CreateLogFile
func SetErrorLog(f io.Writer) {
	log.SetOutput(f)
}