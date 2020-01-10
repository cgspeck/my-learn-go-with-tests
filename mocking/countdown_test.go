package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept += duration
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

const sleep = "sleep"
const write = "write"

type CountdownOperationsSpy struct {
	Calls []string
}

func (s *CountdownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func TestCountDown(t *testing.T) {
	t.Run("prints correct countdown", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeper := &SpySleeper{}

		Countdown(buffer, spySleeper)

		e := `3
2
1
Go!
`
		a := buffer.String()

		if a != e {
			t.Errorf("expected: %q actual: %q", e, a)
		}

		if spySleeper.Calls != 4 {
			t.Errorf("expected 4 calls to spySleeper, got: %d", spySleeper.Calls)
		}
	})

	t.Run("sleeps, then writes in a loop", func(t *testing.T) {
		opsSpy := &CountdownOperationsSpy{}

		Countdown(opsSpy, opsSpy)

		e := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}
		if !reflect.DeepEqual(e, opsSpy.Calls) {
			t.Errorf("expected: %v actual: %v", e, opsSpy.Calls)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}
