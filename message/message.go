package message

import (
	"errors"
	"fmt"
)

func ShowMessage(msg string) (string, error) {

	if msg == "" {
		return "", errors.New("you can't send a empty message")
	}

	result := fmt.Sprintf("Hi, this is your message: %v", msg)
	return result, nil
}
