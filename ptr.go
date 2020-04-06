package main
import "fmt"

func Array_ptr(arr *[3]int){

	arr[0] = 0
}

func main(){
	a := 12
	var p *int
	p = &a
	//p := &a
	//var p = &a
	fmt.Println(a)
	fmt.Println(&a == p)
	fmt.Println(*p)

	arr := [3]int{1, 2, 3}
	Array_ptr(&arr)
	fmt.Println(arr)

}