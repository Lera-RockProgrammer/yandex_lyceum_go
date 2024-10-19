package main

import (
	"fmt"
	"lesson/geometry"
	"lesson/students"
)

func main() {
	student1 := students.Student{Name: "pupka"}
	fmt.Println(student1)
	rect := geometry.Rectangle{Width: 10, Height: 5}
	area := rect.Area()
	scaledW, scaledH := rect.Scale(2.0)
	fmt.Println(area, scaledH, scaledW)
}
