package pkg

import (
	"fmt"
	"log"
	"os"
	"time"
)

type CustomLogginAndErrorInterface interface {
	Log(string, error)
	Error() string
}

type CustomLogginAndError struct {
	Code    string
	Message string
}

func (c *CustomLogginAndError) Log(code string, message error) {
	file, err := os.OpenFile("log.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()
	t := time.Now()
	logMessage := fmt.Sprintf("%v: %s: %v\n", t.Format("2006-01-02 15:04:05"), code, message)
	_, writeErr := file.WriteString(logMessage)
	if writeErr != nil {
		log.Println(err)
		return
	}
}
func (c *CustomLogginAndError) Error() string {
	file, err := os.OpenFile("Error.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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

func NewCustomLogginAndError() CustomLogginAndErrorInterface {
	return &CustomLogginAndError{}
}
