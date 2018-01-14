package console

import (
	"fmt"
)

func Log(a ...interface{}) {
	fmt.Println(a...)
}

func Pizza(a ...interface{}) {
	fmt.Print("🍕 ")
	fmt.Println(a...)
}

func Beer(a ...interface{}) {
	fmt.Print("🍺 ")
	fmt.Println(a...)
}

func Unicorn(a ...interface{}) {
	fmt.Print("🦄 ")
	fmt.Println(a...)
}

func Poop(a ...interface{}) {
	fmt.Print("💩 ")
	fmt.Println(a...)
}
