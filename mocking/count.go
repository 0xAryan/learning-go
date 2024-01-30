package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3


func Countdown(out io.Writer, sleeper Sleeper) {

	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprintln(out, finalWord)
}


// // Any objec that will have a method name Sleep()
// // will be qualified as Sleeper type interface
// type Sleeper interface {
// 	Sleep()
// }

// // This is a genuin struct to implement real sleep
// type DefaultSleeper struct{}

// func (d *DefaultSleeper) Sleep() {
// 	time.Sleep(1 * time.Second)
// }


type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration) // specially this part 
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}


// need to revisit the custom time thing can't understand
func main() {
	// sleeper := &DefaultSleeper{}
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}