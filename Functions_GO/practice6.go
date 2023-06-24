package main
import "fmt"
//interface in Go defines a set of method signatures
//Polymorphism is the ability of a VALUE of a certain TYPE to also be of another TYPE
type person struct{
	first string
}
type secretagent struct{
	person
	itk bool
}


func(p person)speak(){
	fmt.Println("I am",p.first)

}
func(sa secretagent)speak(){
	fmt.Println("I am a secretagent",sa.first)

}
type human interface{
	speak()
}
func saysomething(h human){
	h.speak()
}

func main(){
	sa1:= secretagent{
		person: person{
		first: "james",
	},
	itk: true,
	}
	p2:= person{
		first: "jenny",
	}
	sa1.speak()
	p2.speak()
	saysomething(sa1)
	saysomething(p2)
}