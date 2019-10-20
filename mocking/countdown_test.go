package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

const (
	sleep = "sleep"
	write = "write"
)

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

func TestCountdown(t *testing.T) {
	t.Run("prints 3 to Go!", func(t *testing.T) {
		buf := bytes.Buffer{}
		sleeper := CountdownOperationsSpy{}

		Countdown(&buf, &sleeper)

		want := "3\n2\n1\nGo!\n"
		got := buf.String()

		if got != want {
			t.Errorf("Countdown() = %q, want %q", got, want)
		}
	})

	t.Run("sleep before every print", func(t *testing.T) {
		spySleepPrinter := CountdownOperationsSpy{}
		Countdown(&spySleepPrinter, &spySleepPrinter)

		got := spySleepPrinter.Calls
		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v calls, want %v", got, want)
		}
	})
}

type SpyTIme struct {
	durationSlept time.Duration
}

func (s *SpyTIme) Sleep(d time.Duration) {
	s.durationSlept = d
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := SpyTIme{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("slept for %v, wanted to sleep for %v", spyTime.durationSlept, sleepTime)
	}
}
