package main
import "fmt"

//an anonymous function
//func(p parameters(s)) (r return(s)){code}

func main(){
	foo()

	func(){
		fmt.Println("Anonymus func ran")
	}()
	func (s string)  {
		fmt.Println("This is anonymous func showing my name ",s)
	}("Dakshina")
}
func foo(){
	fmt.Println("Foo ran")
}