package main
import "fmt"

func main(){

	func(){
		fmt.Println("Hello, Nitesh")
	}()  //we are calling the func here only

	func(name string){
		fmt.Println(name)
	}("Hello")  //we are calling the func here only
}


