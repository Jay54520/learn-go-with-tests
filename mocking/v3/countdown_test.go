package main

import (
	"testing"
	"bytes"
)

func TestCountDown(t *testing.T)  {
	buffer := &bytes.Buffer{}

	CountDown(buffer)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep()  {
	s.Calls++
}