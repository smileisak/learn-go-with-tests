package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	write          = "write"
	sleep          = "sleep"
	finalWord      = "Go!"
	countdownStart = 3
)

// Sleeper interface for general purpose
type Sleeper interface {
	Sleep()
}

// DefaultSleeper in the implementation of Sleeper with real sleep
type DefaultSleeper struct{}

// Sleep metheods to real sleep
func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

// CountdownOperationsSpy type
type CountdownOperationsSpy struct {
	Calls []string
}

// Sleep implementation
func (s *CountdownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

// ConfigurableSleeper is a sleeper with controlled time
type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

// Sleep function
func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

// SpyTime for unit tests
type SpyTime struct {
	durationSlept time.Duration
}

// Sleep on spy
func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

// Countdown count until 3 and go
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
