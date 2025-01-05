package a

import (
	"fmt"

	"github.com/wangjian890523/LearnningPath/Core36/C2/packageinit/b"
)

func init() {
	fmt.Println("initializing  package a.")
}

func AFunc() {
	fmt.Println("Function in package a ")
	b.BFunc()
}
