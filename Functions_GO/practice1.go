package main 
import "fmt"
// func (r receiver)identifier (p parameter(s)) (return(s)){.......}
func main(){
	foo()
	bar("dakshina")
	s:=aloha("Dakshina")
	fmt.Println(s)
	s1,n:= dog("Roxie",2)
	fmt.Println(s1,n)
}
func foo(){
	fmt.Println("I am from foo")
}

func bar(s string){
	fmt.Println("My name is",s)

}

func aloha(s string)string{
	return fmt.Sprint("My name is",s)
}

func dog(name string,age int)(string,int){
	age*=2
	return fmt.Sprint(name,"Ths age "),age
}