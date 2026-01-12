package betterLogs

import (
	"bufio"
	"io"
	"log"
	"os"
	"time"
)

var (
	logFile       *os.File
	loggerError   *log.Logger
	loggerMessage *log.Logger
	mainFolder    string = "logs"
	filePath      string = mainFolder + "/logsfile.txt" //Location of the main log file
	maxLine       int16  = 150                          //The maximum number of lines the main log file can have
	oldLogsPath          = mainFolder + "/old_logs"     //Location of the old logs files
	finalMessage  string
)

func CheckLogFile() {

	file, err := os.OpenFile(filePath, os.O_RDWR, 0600)
	if err != nil {
		log.Fatalf("Failed to open log file: %s", err)
	}
	defer file.Close()

	lineCount := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineCount++
	}

	if lineCount >= int(maxLine) {
		if _, err := os.Stat(oldLogsPath); os.IsNotExist(err) {
			err = os.Mkdir(oldLogsPath, 0700)
			if err != nil {
				log.Fatalf("Failed to create "+oldLogsPath+" directory: %s", err)
			}
		}

		currentTime := time.Now().Format("2006-01-02_15-04-05")
		backupFileName := oldLogsPath + currentTime + "_logs.txt"

		backupFile, err := os.Create(backupFileName)
		if err != nil {
			log.Fatalf("Failed to create backup log file: %s", err)
		}
		defer backupFile.Close()

		_, err = file.Seek(0, 0)
		if err != nil {
			log.Fatalf("Failed to seek log file: %s", err)
		}

		_, err = io.Copy(backupFile, file)
		if err != nil {
			log.Fatalf("Failed to copy log file: %s", err)
		}

		err = file.Truncate(0)
		if err != nil {
			log.Fatalf("Failed to truncate log file: %s", err)
		}
	}
}

func init() {
	var err error

	err = os.MkdirAll(mainFolder, 0700)
	if err != nil {
		log.Fatalf("Failed to create logs directory: %s", err)
	}

	logFile, err = os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatalf("Failed to open error log file: %s", err)
	}

	loggerError = log.New(logFile, "ERROR: ", log.Ldate|log.Ltime)

	loggerMessage = log.New(logFile, "", log.Ldate|log.Ltime)

	CheckLogFile() //Check if the log file has reached the maximum number of lines
}

// Function to log a single error
func LogError(err error) {
	loggerError.Printf("|-| %s", err)
}

// Function to log a single message
func LogMessage(message string) {
	loggerMessage.Printf("|-| %s", message)
}

// Function to log an error with extra text message
func LogErrow(err error, message string) {

	if len(message) != 0 {
		finalMessage = message + "|-|" + err.Error()
	} else {
		finalMessage = err.Error()
	}
	loggerError.Printf("|-| %s", finalMessage)
}
