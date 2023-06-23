package main
import "fmt"

func main(){

	am:= map[string]int{
		"todd": 41,
		"henry":16,
		"padget": 14,
	}

	fmt.Println("The age of henry was",am["henry"])
	fmt.Printf("%#v\n",am)

	an:= make(map[string]int)
	an["Lucas"] = 28
	an["Steph"] = 37

	fmt.Println(an)
	fmt.Println(len(an))
}