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

	age1 :=m["James"]
	fmt.Println("The age is",age1)

	age2 := m["Q"]
	fmt.Println("That age of Q",age2)

	if v,ok:= m["Q"]; !ok{
		fmt.Println("There is no Q,and here is the Zero value of an int",v)
	}
}