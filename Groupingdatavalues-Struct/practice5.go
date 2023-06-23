package main

import (
	"fmt"
	
)

func main(){
	p1:= struct{
		first string
		friends map[string]int
		favDrinks []string
	}{
		first: "james",
		friends: map[string]int{
			"jenny":27,
			"q":87,
			"Ian":47,

		},
		favDrinks:[]string{
			"martini",
			"water",
		},
	}
	for _,v:=range p1.friends{
		fmt.Println(p1.first,"-friends-",v)
	}
	for _,v:=range p1.favDrinks{
		fmt.Println(p1.first,"-drinks-",v)
	}
}