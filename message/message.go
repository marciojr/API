package message

import (
	"fmt"
)

func ShowMessage(msg string) string {
	result := fmt.Sprintf("Hi, this is your message: %v", msg)
	return result
}
