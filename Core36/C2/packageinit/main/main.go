package main

import (
	"fmt"

	"github.com/wangjian890523/LearnningPath/Core36/C2/packageinit/a"
)

func init() {
	fmt.Println("initializing  main package")

}

func main() {
	fmt.Println("Mian funcitons ")
	a.AFunc()
}
