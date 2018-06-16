package main

import (
	"io"
	"os"
	"fmt"
)

func main()  {
	CountDown(os.Stdout)
}

const finalWord  = "Go!"
const countdownStart  = 3

func CountDown(writer io.Writer)  {
	for i := countdownStart; i > 0; i -- {
		fmt.Fprintln(writer, i)
	}
	fmt.Fprintf(writer, finalWord)
}