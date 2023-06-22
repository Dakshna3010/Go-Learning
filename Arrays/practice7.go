//sort
package main
import ("fmt"
		"sort"
	)

func main(){
	n := []float64{3,1,4,2}

	fmt.Println(medianOne(n))
	fmt.Println("after medianOne",n)
	fmt.Println(medianTwo(n))
	fmt.Println("after medianTwo",n)
}
func medianOne(x []float64)float64{
	sort.Float64s(x)
	i:= len(x)/2
	if len(x)%2==1{
	return x[i/2]
	}
	return (x[i-1]+x[i])/2
}

func medianTwo(x []float64)float64{
	n:= make([]float64,len(x))
	copy(n,x)

	sort.Float64s(n)
	i:= len(n)/2
	if len(n)%2==1{
		return n[1/2]
	}
	return (n[i-1]+ n[i])/2
}
