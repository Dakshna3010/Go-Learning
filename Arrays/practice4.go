//append slice
package main

import "fmt"

func main(){
	xi:=[]int{41,42,43}
	fmt.Println(xi)

	xi = append(xi,45,46,47,34)
	fmt.Println(xi)
}