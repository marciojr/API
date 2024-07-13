package message

import (
	"regexp"
	"testing"
)

func TestShowMessage(t *testing.T) {
	msg := "Golang is realy cool"
	want := regexp.MustCompile(`\b` + msg + `\b`)
	result, err := ShowMessage(msg)

	if !want.MatchString(result) || err != nil {
		t.Fatalf(`Message name("Golang") = %q, %v, want match for %#q, nil`, result, err, want)
	}
}

func TestShowEmptyMessage(t *testing.T) {
	msg := ""
	result, err := ShowMessage(msg)

	if result != "" || err == nil {
		t.Fatalf(`Empty Message ("") = %q, %v, want match for ""`, result, err)
	}
}
