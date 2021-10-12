package main

import (
	"fmt"
	"pbrt-go/mymath/vector3d"
)

func main() {
	str := text()
	v := vector3d.NewVector3d(3, 4, 5)
	fmt.Println(str)
	fmt.Println(v)
	fmt.Println(v.LengthSq())
}

func text() string {
	return "Hi there"
}
