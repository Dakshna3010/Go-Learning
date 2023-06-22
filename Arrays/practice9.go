package main  
import "fmt"

func main(){

	a:= []int{}

	for i:=42;i<=51;i++{
		a=append(a, i)
	}

	fmt.Println("The value of a is %v \n",a)


	for i,v:=range a{
		fmt.Printf("The index is %v and the type is %T and the value is %v\n",i,v,v)
	}

}