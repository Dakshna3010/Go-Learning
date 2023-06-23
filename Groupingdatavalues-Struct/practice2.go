package main
import "fmt"

type person struct{
	first string
	last string
	FavIc []string
}

func main(){
	p1 := person{
		first : "james",
		last: "Bomd",
		FavIc: []string{"chocolate","Strawberry","Butterscoch"},

	}

	p2:=person{
		first: "joo",
		last: "Luc",
		FavIc: []string{"Chocolate","vanila","Blueberry"},
	}
	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.FavIc)
	fmt.Println(p2.FavIc)

	for _,v:=range p1.FavIc{
		fmt.Println(p1.first, "favourite is",v)
	}
	for _,v:=range p2.FavIc{
		fmt.Println(p1.first, "favourite is",v)
	}
}