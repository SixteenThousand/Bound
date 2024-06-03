package main

import(
    bound "github.com/SixteenThousand/Bound/src"
    "flag"
)

const FLAG_HELP string = "see --help"

func main() {
    var timerLength string
    flag.StringVar(&timerLength,"length","4m20s",FLAG_HELP)
    flag.Parse()
    bound.Run(timerLength)
}
