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
	// log here
	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return c.Error()
	}
	t := time.Now()
	file.WriteString(
		fmt.Sprintf("%v: %s at: %v\n", t.Format("2006-01-02 15:04:05"), c.Code, c.Message),
	)
	return fmt.Sprintf("%v: %s at: %v", t.Format("2006-01-02 15:04:05"), c.Code, c.Message)
}
