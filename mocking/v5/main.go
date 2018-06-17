package main

import (
	"io"
	"os"
	"fmt"
	"time"
)

func main()  {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	CountDown(os.Stdout, sleeper)
}

const finalWord  = "Go!"
const countdownStart  = 3

func CountDown(writer io.Writer, sleeper Sleeper)  {
	for i := countdownStart; i > 0; i -- {
		sleeper.Sleep()
		fmt.Fprintln(writer, i)
	}

	sleeper.Sleep()
	fmt.Fprintf(writer, finalWord)
}

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep func(duration time.Duration)
}

func (c *ConfigurableSleeper) Sleep()  {
	c.sleep(c.duration)
}