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
	fmt.Fprintf(writer, "3")
}