package main

import "fmt"

func main(){
	Xi:= []int{42,43,44,45,46,47}
	for i,v := range Xi{
		fmt.Printf("index %v \t valve %v \n",i,v)
	}
}