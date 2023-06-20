package main

import (
	"fmt"
	"math/rand"
)

func main(){
	X := rand.Intn(400)
	fmt.Printf("The number is %v \t \n",X)

	if X<=100{
		fmt.Println("The number is less than 100\t")
	}else if X>101 && X<=200{
		fmt.Println("The number is between 101 and 200")
	}else if X>201 && X<=300{
		fmt.Println("The number is between 201 and 300")
		
	}else{
		fmt.Println("The number is above 300")
	}

}


