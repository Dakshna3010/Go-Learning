package main
import "fmt"

func main(){
	an:= make(map[string]int)
	an["Lucas"] = 28
	an["Steph"] = 37

	fmt.Println(an)
	fmt.Println(len(an))


	for k,v := range an{
		fmt.Println(k,v)
	}
	for _,v := range an{
		fmt.Println(v)
	}
	xi:=[]int{2,4,6,45}
	fmt.Println(xi)


	for i,v:=range xi{
		fmt.Println(i,v)
	}
}