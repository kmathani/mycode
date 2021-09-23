package main

import (
    "fmt"
    "runtime"
    "strings"
)

func main() {

	var tVal  string = "hello"
           // init gover                
    var gover string = runtime.Version()
    switch  {
    case strings.Contains(gover,"go1.17"):                 // if matches "go1.17", do the code below then STOP
        fmt.Println("You are using the latest version of GoLang")
    case strings.Contains(gover,"go1.16"),strings.Contains(gover, "go1.16.5"),strings.Contains(gover, "go1.15"):       // if matches "go1.16", "go1.16.5", OR "go1.15" 
        fmt.Println("You are using an older, but acceptable version of GoLang")
    case strings.Contains(tVal, "hello"):
	    fmt.Println("Does it Print ?")
    default:                       // in all other cases run the code below
        fmt.Println("I cannot make a recommendation.")
    }
}

