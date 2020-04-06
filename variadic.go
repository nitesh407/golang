package main
import "fmt"

//variadic function that returns the multiplication 
func mul(nums... int)(int){

	x := 1
	for _, num := range nums{
		x *= num
	}
	return x
}

func main(){

	fmt.Println(mul(1, 2, 3, 4, 5))
	fmt.Println(mul())

	s := []int{1, 2, 3, 4, 5}
	fmt.Println(mul(s...))
}