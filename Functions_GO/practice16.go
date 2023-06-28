package main

import (
	"fmt"
	"math"
)

type square struct{
	length float64
	width float64
}
type circle struct{
	radius float64
}

func (s square)Area() float64 {
	return	s.length * s.width
}
func(c circle)Area()float64{
	return math.Pi*math.Pow(c.radius,2)
}
type shape interface{
	Area()float64
}
func info(s shape)float64{
	return	s.Area()

}
func main(){
	c1:= circle{
		radius:4,
	}
	s1:=square{
		length : 5,
		width: 8,
	}
	fmt.Println(info(c1))
	fmt.Println(info(s1))
}