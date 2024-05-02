package bound

import (
    "fmt"
    "time"
)

const MINS_PER_HR = 60
const SECS_PER_MIN = 60

func Run() {
    println("We got here!")
    durr, err := time.ParseDuration("2m")
    if err != nil {
        fmt.Println("Oh no, we got here.")
        panic(1)
    }
    RunTimer(durr)
}

func DurationToString(d time.Duration) string {
    var hrs, mins, secs int
    hrs = Floor(d.Hours())
    mins = Floor(d.Minutes()) - (MINS_PER_HR * hrs)
    secs = (Floor(d.Seconds()) - 
        (SECS_PER_MIN * MINS_PER_HR * hrs) -
        (SECS_PER_MIN * mins))
    return fmt.Sprintf("%02dh %02d' %02d\"",hrs,mins,secs)
}

func RunTimer(d time.Duration) {
    end := time.Now().Add(d)
    endMsg := make(chan int,1)
    time.AfterFunc(d,func() {
        endMsg <- 1
    })
    go func() {
        for {
            if len(endMsg) > 0 {
                break
            }
            fmt.Printf("\rTime remaining: %s",time.Until(end).String())
            time.Sleep(100 * time.Millisecond)
        }
    }()
    println(<- endMsg)
}
