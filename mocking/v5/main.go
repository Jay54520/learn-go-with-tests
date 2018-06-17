package main

import (
	"io"
	"os"
	"fmt"
	"time"
)

func main()  {
	sleeper := &DefaultSleeper{}
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

type DefaultSleeper struct {

}

func (d *DefaultSleeper) Sleep()  {
	time.Sleep(1 * time.Second)
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep func(duration time.Duration)
}

func (c *ConfigurableSleeper) Sleep(duration time.Duration)  {
	c.sleep(duration)
}