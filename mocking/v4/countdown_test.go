package main

import (
	"testing"
	"bytes"
	"reflect"
)

func TestCountDown(t *testing.T) {
	t.Run("prints 5 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}

		CountDown(buffer, &CountdownOperationsSpy{})

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})

	t.Run("sleep after every print", func(t *testing.T) {
		spySLeepPrinter := &CountdownOperationsSpy{}
		CountDown(spySLeepPrinter, spySLeepPrinter)

		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(spySLeepPrinter.Calls, want) {
			t.Errorf("wanted calls %v got %v", want, spySLeepPrinter.Calls)
		}
	})
}

type CountdownOperationsSpy struct {
	Calls []string
}

func (s *CountdownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
	return
}

func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const write = "write"
const sleep = "sleep"
