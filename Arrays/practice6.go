//multidimensional slice
package main
import "fmt"

func main(){
	jb:= []string{"james","Bond","martin","Chocolate"}
	jm:= []string{"jenny","Moneypenny","Guiness","Wolverian Tracks"}
	fmt.Println(jb)
	fmt.Println(jm)

	xxs:=[][]string{jb,jm}
	fmt.Println(xxs)
}