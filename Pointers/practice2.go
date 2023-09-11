package main
import "fmt"

func intdelta(n *int){
	*n=50
}

func main(){
	a:=42
	fmt.Println(a)
	intdelta(&a)
	fmt.Println(a)

}