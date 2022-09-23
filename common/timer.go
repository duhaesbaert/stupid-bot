package common

import (
	"fmt"
)

type Timer struct {
	minute int
	second int
}

// Newtimer initializes a Timer with the determined minutes. If the minutes passed as argument are greater than 59,
// it will set minutes to 59 and seconds to 59.
func Newtimer(minutes int) Timer {
	seconds := 0
	if minutes > 59 {
		minutes = 59
		seconds = 59
	}
	return Timer{
		minute: minutes,
		second: seconds,
	}
}

// Countdown is exporting the functionality to subtract the time from the timer, 1 second each request.
func (t Timer) Countdown() Timer {
	return t.subtractSecond()
}

// subtractSecond subtracts 1 second each time it is called, if seconds are 0, subtracts a minute and returns to 59.
func (t Timer) subtractSecond() Timer {
	if t.second == 0 {
		if t.minute > 0 {
			t = t.subtractMinute()
			t.second = 59
		}
	} else {
		t.second--
	}
	return t
}

// subtractMinute subtracts 1 from minutes, and returns if the countdown should continue.
func (t Timer) subtractMinute() Timer {
	t.minute--
	return t
}

func (t Timer) ShowNormalizedTime() string {
	var minutes string
	if t.minute < 10 {
		minutes = "0"
	}
	minutes += fmt.Sprintf("%d", t.minute)

	var seconds string
	if t.second < 10 {
		seconds = "0"
	}
	seconds += fmt.Sprintf("%d", t.second)

	return fmt.Sprintf("%sm%ss", minutes, seconds)
}

// TimerMinutes returns the current minute from the Timer.
func (t Timer) TimerMinutes() int {
	return t.minute
}
