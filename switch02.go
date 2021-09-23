package main

import (
    "fmt"
    "time"
)

func main() {
    
    // determine what time it is
    watch := time.Now() 
 
    other := 10
    // there is no condition here
    // read as, "switch true"
    switch {
    case watch.Hour() < 6:
        fmt.Println("Go back to sleep.")
    case watch.Hour() < 12:
        fmt.Println("Good morning!")
    case watch.Hour() < 18:
	    fmt.Println("Good afternoon.")
    case other == 10:  //  This will be ignored, go will execute the first match and stop the switch statement.
	    fmt.Println("10 will not be prrinted .....")
    default:
        fmt.Println("Counting sheep. Z-z-z-z-z-z-z")
    }
}

