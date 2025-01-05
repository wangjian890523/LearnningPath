package main

import "fmt"

var globalVar = initGlobalVar()

func initGlobalVar() string {
	fmt.Println("Initializing global variable.")
	return "Global variable"

}

func init() {
	fmt.Println("Init function 3.")
}
func init() {
	fmt.Println("Init function 1.")
}

func init() {
	fmt.Println("Init function 2.")
}

func main() {
	fmt.Println("main function. ")
	fmt.Println(globalVar)
}
