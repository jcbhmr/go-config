package main

import (
	"fmt"
	"log"
	"runtime"
	"flag"
)

var goosFlag = flag.Bool("goos", false, "print only runtime.GOOS")
var goarchFlag = flag.Bool("goarch", false, "print only runtime.GOARCH")

func main() {
	log.SetFlags(0)
	flag.Parse()
	if *goosFlag {
		fmt.Println(runtime.GOOS)
	} else if *goarchFlag {
		fmt.Println(runtime.GOARCH)
	} else {
		fmt.Println(runtime.GOOS + "/" + runtime.GOARCH)
	}
}
