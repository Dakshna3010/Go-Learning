package main 
import "fmt"

func main(){
	fmt.Println(printsquare(square,3))
}
func printsquare(f func(int)int,a int)string{
	x:= f(a)
	return fmt.Sprintf("The number %v squared is %v",a,x)
}
func square(n int)int{
	return n*n
}