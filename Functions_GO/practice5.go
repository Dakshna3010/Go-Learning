package main
import "fmt"
//Attach method to a type
type person struct{
	first string
}

func(p person)speak(){
	fmt.Println("I am",p.first)

}

func main(){
	p1:= person{
		first: "james",
	}
	p2:= person{
		first: "jenny",
	}
	p1.speak()
	p2.speak()
}