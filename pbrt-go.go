package main

import (
	"fmt"
	"pbrt-go/mymath"
)

func main() {
	str := text()
	v := mymath.NewVector3(3, 4, 5)
	fmt.Println(str)
	fmt.Println(v)
	fmt.Println(v.LengthSq())
}

func text() string {
	return "Hi there"
}
