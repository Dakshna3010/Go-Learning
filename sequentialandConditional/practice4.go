package main
import "fmt"

func main(){

	m := map[string]int{
		"james": 42,
		"Moneypenny": 32,
	}

	for k,v:= range m{
		fmt.Printf("key is %v\t and value is %v \n",k,v)
	}
}