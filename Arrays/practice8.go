package main 
import "fmt"

func main(){
//craete an array
	a:=[5]int{}
//asigning values
	for i:=0;i<5;i++{
		a[i]=i
	}
	//print out
	for i,v:= range a{
		fmt.Printf("%v - %T - %v\n",v,v,i)
	}


}