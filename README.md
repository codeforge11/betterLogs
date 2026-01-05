
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
    // Attempt to read a file
    _, err := os.ReadFile("test.txt")
    if err != nil {
        // Log the error details to the default log file
        betterLogs.LogError(err)
        return
    }

    // Log a standard informational message
    betterLogs.LogMessage("File opened successfully")
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