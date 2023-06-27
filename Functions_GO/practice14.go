package main
import "fmt"

func main(){
	defer fmt.Println(3)
	defer fmt.Println(2)
	defer fmt.Println(1)
	fmt.Println(0)
}
/*"defer"multiple fun in main
Show that a deferred func runs after the func conatining it exist
determine the order in which the multiple defer func runs*/
// deferred func runs in LIFO order
