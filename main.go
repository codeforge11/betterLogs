package betterLogs

import (
	"bufio"
	"io"
	"log"
	"os"
	"time"
)

type Config struct {
	MainFolder  string
	FilePath    string //Location of the main log file
	MaxLine     int16  //The maximum number of lines the main log file can have
	OldLogsPath string //Location of the old logs files

}

var (
	logFile       *os.File
	loggerError   *log.Logger
	loggerMessage *log.Logger
	finalMessage  string
)

func New(c Config) *Config {
	if c.MainFolder == "" {
		c.MainFolder = "logs"
	}
	if c.FilePath == "" {
		c.FilePath = c.MainFolder + "/logsfile.txt"
	}
	if c.MaxLine == 0 {
		c.MaxLine = 150
	}
	if c.OldLogsPath == "" {
		c.OldLogsPath = c.MainFolder + "/oldLogs"
	}
	c.Init()
	return &c
}

func (c *Config) CheckLogFile() {

	file, err := os.OpenFile(c.FilePath, os.O_RDWR, 0600)
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
		if _, err := os.Stat(c.OldLogsPath); os.IsNotExist(err) {
			err = os.Mkdir(c.OldLogsPath, 0700)
			if err != nil {
				log.Fatalf("Failed to create "+c.OldLogsPath+" directory: %s", err)
			}
		}

		currentTime := time.Now().Format("2006-01-02_15-04-05")
		backupFileName := c.OldLogsPath + currentTime + "_logs.txt"

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

func (c *Config) Init() {

	err := os.MkdirAll(c.MainFolder, 0700)
	if err != nil {
		log.Fatalf("Failed to create logs directory: %s", err)
	}

	logFile, err = os.OpenFile(c.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatalf("Failed to open error log file: %s", err)
	}

	loggerError = log.New(logFile, "ERROR: ", log.Ldate|log.Ltime)

	loggerMessage = log.New(logFile, "", log.Ldate|log.Ltime)

	c.CheckLogFile() //Check if the log file has reached the maximum number of lines
}

// Function to log a single error
func (c *Config) LogError(err error) {
	loggerError.Printf("|-| %s", err)
}

// Function to log a single message
func (c *Config) LogMessage(message string) {
	loggerMessage.Printf("|-| %s", message)
}

// Function to log an error with extra text message
func (c *Config) LogErrow(err error, message string) {

	if len(message) != 0 {
		finalMessage = message + "|-|" + err.Error()
	} else {
		finalMessage = err.Error()
	}
	loggerError.Printf("|-| %s", finalMessage)
}
