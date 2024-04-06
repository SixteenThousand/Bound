package Bound

import (
    "fmt"
    "time"
    "math"
)

const MINS_PER_HR = 60
const SECS_PER_MIN = 60

func Run() {}

func floor(x float64) int {
    return int(math.Floor(x))
}

func (d Duration) String() {
    var hrs, mins, secs int
    hrs = floor(d.Hours())
    mins = floor(d.Minutes()) - (MINS_PER_HR * hrs)
    secs = floor(d.Seconds) 
        - (SECS_PER_MIN * MINS_PER_HR * hrs)
        - (SECS_PER_MIN * mins)
    return fmt.Sprintf("%02dh %02d' %02d\"",hrs,mins,secs)
}

func RunTimer(d Duration) {
    endMsg := false
    time.Afterfunc(d,func() {
        endMsg = true
    })
    end := time.Now().Add(d)
    go func(endMsg) {
        while !endMsg {
            fmt.Printf("\rTime remaining: %s",time.Until(end).String())
            time.Sleep(100 * time.Millisecond)
        }
    }()
}
