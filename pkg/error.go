package pkg

import (
	"fmt"
	"os"
	"time"
)

type CustomError struct {
	Code    string
	Message string
	// CreateAt time.Time
}

func (c *CustomError) Error() string {
	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Sprintf("Error opening log file: %v", err)
	}
	defer file.Close()
	t := time.Now()
	logMessage := fmt.Sprintf("%v: %s at: %v\n", t.Format("2006-01-02 15:04:05"), c.Code, c.Message)
	_, writeErr := file.WriteString(logMessage)
	if writeErr != nil {
		return fmt.Sprintf("Error writing to log file: %v", writeErr)
	}

	return logMessage
}
