package  main
import "fmt"


func swap(a, b int)(int , int){

	return b, a
}
func main(){
	
	a, b := swap(10, 20)
	fmt.Println(a, b)

	//closures
	// These are the functions assigned to variables
	sum := func(x, y int)(int){
		return x + y
	}
	fmt.Println(sum(2, 3))
}	 


func init(){

	fmt.Println("This is init function ()")
}

func init(){

	fmt.Println("This is also one")
}
