package main

import (
	"io"
	"os"
	"fmt"
	"time"
)

func main()  {
	CountDown(os.Stdout, time)
}

const finalWord  = "Go!"
const countdownStart  = 3

func CountDown(writer io.Writer, sleeper Sleeper)  {
	for i := countdownStart; i > 0; i -- {
		sleeper.Sleep(1 * time.Second)
		fmt.Fprintln(writer, i)
	}

	sleeper.Sleep(1 * time.Second)
	fmt.Fprintf(writer, finalWord)
}

type Sleeper interface {
	Sleep()
}