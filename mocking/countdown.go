package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	finalWord      = "Go!"
	countdownStart = 3
)

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func Countdown(b io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(b, i)
		sleeper.Sleep()
	}

	fmt.Fprint(b, finalWord)
}

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}
