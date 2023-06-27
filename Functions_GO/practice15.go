package main
import "fmt"
type person struct{
	first string
	age int
}
func (p person) speak(){
	fmt.Println("My name is",p.first,"and my age is ", p.age)
}
func main(){
	p1:=person{
		first :"jenny",
		age: 27,
	}
	p1.speak()
}