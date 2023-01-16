package main

import (
	"log"
	"os"
	"path"

	"github.com/go-pg/pg/v10"
	"github.com/kkdai/linebot-ptt-beauty/bots"
	"github.com/kkdai/linebot-ptt-beauty/models"
	"github.com/kkdai/linebot-ptt-beauty/utils"
)

var logger *log.Logger
var meta = &models.Model{}
var logRoot = "logs"
var no_db = true

func main() {
	logFile, _ := initLogFile()
	defer logFile.Close()

	url := os.Getenv("DATABASE_URL")
	options, _ := pg.ParseURL(url)
	db := pg.Connect(options)
	meta.Db = db
	defer db.Close()

	logger = utils.GetLogger(logFile)
	meta.Log = logger

	meta.Log.Println("Start to init Line Bot...")
	bots.InitLineBot(meta, bots.ModeHTTP, "", "")
	meta.Log.Println("...Exit")
}

func initLogFile() (logFile *os.File, err error) {
	logfilename := "pttbeauty.log"
	logFileName := path.Base(logfilename)
	logFilePath := path.Join(logRoot, logFileName)
	if _, err := os.Stat(logRoot); os.IsNotExist(err) {
		os.Mkdir(logRoot, 0755)
	}
	return os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
}
