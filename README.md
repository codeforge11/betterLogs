# betterLogs
[![wakatime](https://wakatime.com/badge/user/f21d1d72-d48f-4c76-8d7d-4781e81e04ec/project/02ce9844-9f0c-4f06-b571-49ad56065708.svg)](https://wakatime.com/badge/user/f21d1d72-d48f-4c76-8d7d-4781e81e04ec/project/02ce9844-9f0c-4f06-b571-49ad56065708)
[![GitHub tag](https://img.shields.io/github/v/tag/codeforge11/betterLogs?style=flat-square)](https://github.com/codeforge11/betterLogs/tags)

**betterLogs** is a lightweight, high-performance logging [GO](https://github.com/golang/go) library designed for simplicity. It provides an intuitive way to generate logs and manage automated archiving with minimal resource overhead.

## Getting Started

### Prerequisites

- **Go version**: To use betterLogs, you must have Go 1.24.0 or higher installed.
- **Minimal Go knowledge**: Only basic knowledge of Go is required to get started.

### Installation

With [Go's module support](https://go.dev/wiki/Modules#how-to-use-modules), simply import betterLogs in your code and Go will automatically fetch it during build:

```go
import "github.com/codeforge11/betterLogs"
```

## Example

```go
package main

import (
	"os"

	"github.com/codeforge11/betterLogs"
)

func main() {

	cfg := betterLogs.Config{
		MainFileName:     "YourOwnMainFileName",
		MainFolder:       "MainFolderName",
		OldLogsFilesName: "YourOldLogsMainFileName",
		OldLogsFolder:    "YourOldLogsFolderName",
		MaxLine:          20,
	}
	Logger := betterLogs.New(cfg)

	// Attempt to read a file
	_, err := os.ReadFile("test.txt")
	if err != nil {
		// Log the error details to the default log file
		Logger.LogError(err)
		return
	}
	// Or
	_, err = os.ReadFile("test.txt")
	if err != nil {
		// Log the error details with your own message
		Logger.LogErrow("Your own optional message", err)
		return
	}

	// Log a standard informational message
	Logger.LogMessage("File opened successfully")
}
```


## More about functions

```go
LogError(err)
```
Records error object output. A lightweight function designed to capture error states without additional formatting.

```go
LogMessage("Example comment") 
```
Provides a versatile way to log custom text. It supports strings of any length, making it ideal for tracking application flow and milestones.

```go
LogErrow("Your own optional message", err)
```
This enables wrapping errors with custom descriptive messages, making it easier to identify the source of the issue.

```go
CheckLogFile()
```
This function checks if the main log file has exceeded the maximum line limit and performs archiving if necessary. It is optional to call this manually, as the library automatically performs this check every time when your application starts.

## Logs Customizations
Name | About |Default
------------- | -------------|-------------
MainFileName  | The name of the primary log file where current errors and messages are stored.| logsfile.txt
MainFolder  | The root directory where your active log files will be created.| logs
OldLogsFilesName| Optional prefix or pattern for archived log files (defaults to date-based naming).| -
OldLogsFolder  |The subfolder where logs are moved after reaching the line limit | [your logs folder]/oldLogs
MaxLine|The maximum number of lines per file before it is automatically moved to the archive.| 150
---------------------------------------