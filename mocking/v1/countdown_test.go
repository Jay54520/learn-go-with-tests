package main

import (
	"testing"
	"bytes"
)

func TestCountDown(t *testing.T)  {
	buffer := &bytes.Buffer{}

	CountDown(buffer)

	got := buffer.String()
	want := "3"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}