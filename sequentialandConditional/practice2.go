//for loop

package main

import (
	"fmt"
	"math/rand"
)

func main(){
	for i:=0;i<=5;i++{
		X := rand.Intn(10)
		Y := rand.Intn(10)
		fmt.Printf("The numbers X and y are %v and %v\n",X,Y)


		switch {
		case X<4 && Y<4:
			fmt.Printf("The numbers X and Y are %v and %v\n",X,Y)
		case X>6 && Y>6:
			fmt.Println("Both are greater than six")
		case X>=4 && X<=6:
			fmt.Println("X is between 4 and 6")
		case Y!=5:
			fmt.Printf("Y is not 5\n")
		default:
			fmt.Printf("None of the previous\n")
		}

	}
}