package main
import "fmt"

func main(){
	x()
}
var x =func (){
	for i :=0;i<10;i++{
		fmt.Println(i)
	}
}