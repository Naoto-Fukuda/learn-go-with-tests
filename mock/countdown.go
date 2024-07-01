package mock

import (
    "fmt"
    "io"
    "os"
		"time"
)

const write = "write"
const sleep = "sleep"

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

type SpyTime struct {
	duration time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration){
	s.duration = duration
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

type Sleeper interface {
	Sleep()
}

// 構造体にCallsフィールドを用意
type SpySleeper struct {
	Calls int
}

type CountdownOperationsSpy struct {
	Calls []string
}

func (s *CountdownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperationsSpy) Write([]byte) (n int, err error){
	s.Calls = append(s.Calls, write)
	return
}

// 構造体にメソッドを紐付け、Callsフィールドに加算
func (s *SpySleeper) Sleep() {
	s.Calls++
}

const finalWord = "Go!"
const countdownStart = 3

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