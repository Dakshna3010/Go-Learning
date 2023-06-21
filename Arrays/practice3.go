//blank identifier to not use a returned value
package main

import "fmt"

func main(){
	XS:= []string{"goa","Bihar","Jammu","Kerala","Karnataka"}
	fmt.Println(XS)
	fmt.Printf("%T\n",XS)

	for _,v := range XS{
		fmt.Printf("%v \n",v)
	}

	fmt.Println(len(XS))
}
