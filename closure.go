package main
import "fmt"

func return_anno()func(string){
	return func(name string){
		fmt.Println(name)
	}
}
func return_rem()func()int{
	i := 0
	return func()int{
		i++
		return i
	}
}
func main(){

	v := func(name string){
		fmt.Println(name)
	}
	v("Nitesh")

	a := return_anno()
	a("Nitesh")

	b := return_rem()
	fmt.Println(b())    //b is a funcion here. So, we need to put parenthasis like a calling statement
	fmt.Println(b())
}