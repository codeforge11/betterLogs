package betterLogs

import (
	"bufio"
	"io"
	"log"
	"os"
	"time"
)

type Config struct {
	MainFolder       string //Main logs folder
	MainFileName     string //Location of the main log file
	MaxLine          int16  //The maximum number of lines the main log file can have
	OldLogsFolder    string //Location of the old logs files
	OldLogsFilesName string //Name of old logs files
}

var (
	logFile       *os.File
	loggerError   *log.Logger
	loggerMessage *log.Logger
)

func New(c Config) *Config {
	if c.MainFolder == "" {
		c.MainFolder = "logs"
	}
	if c.MainFileName == "" {
		c.MainFileName = "logsfile.txt"
	}
	if c.MaxLine == 0 {
		c.MaxLine = 150
	}
	if c.OldLogsFolder == "" {
		c.OldLogsFolder = c.MainFolder + "/oldLogs"
	}
	// if c.OldLogsFilesName==""{
	// 	c.OldLogsFilesName=""
	// }
	c.Init()
	return &c
}

func (c *Config) CheckLogFile() {

	file, err := os.OpenFile((c.MainFolder + "/" + c.MainFileName), os.O_RDWR, 0600)
	if err != nil {
		log.Fatalf("Failed to open log file: %s", err)
	}
	defer file.Close()

	lineCount := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineCount++
	}

	if lineCount >= int(c.MaxLine) {

		err = os.MkdirAll(c.MainFolder+"/"+c.OldLogsFolder, 0700)
		if err != nil {
			log.Fatalf("Failed to create "+c.MainFolder+"/"+c.OldLogsFolder+" directory: %s", err)
		}

		currentTime := time.Now().Format("2006-01-02_15-04-05")
		backupMainFileName := c.OldLogsFilesName + currentTime + "_logs.txt"

		backupFile, err := os.Create(c.MainFolder + "/" + c.OldLogsFolder + "/" + backupMainFileName)
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

func (c *Config) Init() {

	err := os.MkdirAll(c.MainFolder, 0700)
	if err != nil {
		log.Fatalf("Failed to create logs directory: %s", err)
	}

	logFile, err = os.OpenFile((c.MainFolder + "/" + c.MainFileName), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatalf("Failed to open error log file: %s", err)
	}

	loggerError = log.New(logFile, "", log.Ldate|log.Ltime)

	loggerMessage = log.New(logFile, "", log.Ldate|log.Ltime)

	c.CheckLogFile() //Check if the log file has reached the maximum number of lines
}

// Function to log a single error
func (c *Config) LogError(err error) {
	loggerError.Printf("|ERROR| %s", err)
}

// Function to log a single message
func (c *Config) LogMessage(message string) {
	loggerMessage.Printf("|LOG| %s", message)
}

// Function to log an error with extra text message
func (c *Config) LogErrow(err error, message string) {
	var finalMessage string
	if len(message) != 0 {
		finalMessage = message + "|/|" + err.Error()
	} else {
		finalMessage = err.Error()
	}
	loggerError.Printf("|LOG| %s", finalMessage)
}
