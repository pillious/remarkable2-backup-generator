package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"strings"
)

var (
	backupsDir    string
	rsyncSrc      string
	rsyncLinkDest string
	doLogToFile   bool
	verbose       bool

	logFolderName        string = "/logs/"
	logFileExtension     string = ".logs"
	latestBackupFileName string = "/.latest_backup"
)

func main() {
	parseFlags()

	timestamp := getCurrIso8601()
	logFilePath := ""
	if doLogToFile {
		logFilePath = setupFileLogger(timestamp)
	}
	if rsyncSrc == "" {
		log.Fatal("Rsync source not specified")
	}

	rsyncDest := backupsDir + "/" + timestamp
	log.Printf("Creating backup: %s\n", rsyncDest)
	createBackupFolder(rsyncDest)

	latestBackupName := rsyncLinkDest
	if latestBackupName == "" {
		latestBackupName = getLatestBackupName()
	}

	makeBackup(rsyncSrc, rsyncDest, latestBackupName, logFilePath)
	updateLatestBackupName(rsyncDest)
	closeFileLogger()
}

func parseFlags() {
	flag.StringVar(&backupsDir, "backupsDir", ".", "The directory to store backups.")
	flag.StringVar(&rsyncSrc, "src", "", "The source directory for rsync.")
	flag.StringVar(&rsyncLinkDest, "linkDest", "", "The link-dest directory for rsync.")
	flag.BoolVar(&doLogToFile, "l", false, "Write logs to {backupsDir}/logs/ instead of STDOUT.")
	flag.BoolVar(&verbose, "v", false, "Enable verbose output.")

	flag.Parse()
}

func createBackupFolder(folderPath string) {
	createDir(folderPath)
}

func makeBackup(src string, dest string, linkDest string, logFilePath string) {
	rsyncArgs := []string{"-azHl"}
	if verbose {
		rsyncArgs = append(rsyncArgs, "--verbose")
	}
	if linkDest != "" {
		rsyncArgs = append(rsyncArgs, "--link-dest="+linkDest)
	}
	rsyncArgs = append(rsyncArgs, src, dest)

	cmd := exec.Command("rsync", rsyncArgs...)
	if doLogToFile {
		logFile, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			log.Fatal(err.Error())
		}
		defer logFile.Close()

		cmd.Stdout = logFile
		cmd.Stderr = logFile
	} else {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	if verbose {
		log.Println("rsync ", cmd.Args)
	}

	err := cmd.Run()
	if err != nil {
		deleteEmptyDir(dest)
		log.Panic("Error running rsync:", err)
	}

	log.Println("rsync completed successfully")
}

func getLatestBackupName() string {
	filePath := backupsDir + "/" + latestBackupFileName
	name, err := os.ReadFile(filePath)
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(name))
}

func updateLatestBackupName(name string) {
	filePath := backupsDir + "/" + latestBackupFileName
	file, err := os.Create(filePath)
	if err != nil {
		log.Panic(err.Error())
	}
	defer file.Close()

	_, err = file.WriteString(name)
	if err != nil {
		log.Panic(err.Error())
	}
}
