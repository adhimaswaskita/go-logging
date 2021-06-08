package main

import (
	"github.com/adhimaswaskita/go-logging/service"
	"github.com/adhimaswaskita/go-logging/util/logger"
	"log"
)

func main() {
	f := logger.CreateLogFile("log.txt")
	logger.SetErrorLog(f)

	_, err := service.Divide(10, 0)
	if err != nil {
		log.Println("error : ", err)
	}
}