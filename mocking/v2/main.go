package main

import (
	"io"
	"os"
	"fmt"
)

func main()  {
	CountDown(os.Stdout)
}

func CountDown(writer io.Writer)  {
	for i := 3; i > 0; i -- {
		fmt.Fprintln(writer, i)
	}
	fmt.Fprintf(writer, "Go!")
}