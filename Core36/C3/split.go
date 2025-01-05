package main

import (
	"flag"
	"fmt"
)

var name string

func init() {
	flag.StringVar(&name, "name", "every", "the gretting object.")

}

func main() {
	flag.Parse()
	hello(name)
	fmt.Println("vim-go")
}
