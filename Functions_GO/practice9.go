package main
import "fmt"

func main(){
	fmt.Printf("%T\n",add)
	fmt.Printf("%T\n",subtract)
	fmt.Printf("%T\n",math)

	x := math(42,16,add)
	fmt.Println(x)

	y:=math(42,16,subtract)
	fmt.Println(y)

}
func math(a int, b int, f func(int,int)int)int{
	return f(a,b)
}

func add(a int,b int)int{
	return a + b
}
func subtract(a int, b int)int{
	return a - b
}