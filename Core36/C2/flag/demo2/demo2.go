package main

import (
	"flag"
	"fmt"
	"os"
)

var name string

func init() {
	flag.StringVar(&name, name, "everyone", "the greeting object")

}

func main() {

	flag.CommandLine.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", "qustion")
		flag.PrintDefaults()
	}
	flag.Parse()
	fmt.Printf("hello, %s!\n", name)
}
