package main

import (
	"log"
	"os"
	"time"
)

var logFile *os.File

func setupFileLogger(fileName string) string {
	folderPath := backupsDir + logFolderName
	createDir(folderPath)

	filePath := folderPath + fileName + logFileExtension

	var err error
	logFile, err = os.OpenFile(filePath, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Panic(err)
	}

	log.SetOutput(logFile)

	return filePath
}

func closeFileLogger() {
	if logFile != nil {
		log.SetOutput(os.Stdout)
		err := logFile.Close()
		if err != nil {
			log.Println("Error closing log file:", err)
		}
		log.Printf("Logs written to %s\n", logFile.Name())
	}
}

func createDir(path string) {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return
	}

	err := os.MkdirAll(path, 0755)
	if err != nil {
		log.Panic(err.Error())
	}
}

func deleteEmptyDir(path string) {
	if verbose {
		log.Printf("Deleting empty backup directory: %s\n", path)
	}
	err := os.Remove(path)
	if err != nil {
		log.Panic(err.Error())
	}
}

func getCurrIso8601() string {
	currentIsoTimeStr := time.Now().UTC().Format(time.RFC3339)
	return currentIsoTimeStr
}
