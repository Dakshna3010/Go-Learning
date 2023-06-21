package main
import "fmt"

func main(){

for i:=0;i<100;i++{
	if x := rand.Intn(5); x==3{
		fmt.Printf("iteration %v \t is %v\n",i,x)
	}

	}
}