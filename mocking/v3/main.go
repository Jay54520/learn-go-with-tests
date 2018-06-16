package main

import (
	"io"
	"os"
	"fmt"
	"time"
)

func main()  {
	CountDown(os.Stdout)
}

const finalWord  = "Go!"
const countdownStart  = 3

func CountDown(writer io.Writer)  {
	for i := countdownStart; i > 0; i -- {
		time.Sleep(1 * time.Second)
		fmt.Fprintln(writer, i)
	}

	time.Sleep(1 * time.Second)
	fmt.Fprintf(writer, finalWord)
}