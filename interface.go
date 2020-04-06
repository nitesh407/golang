package main
import "fmt"

type Circle interface{

	circumpherence() int
	area() int
}
type C1 struct{
	radius int
}

func (c C1)circumpherence()int{

	return 2* 3 * c.radius
} 
func (c C1)area()int{
	return 3 * c.radius * c.radius
}

func main(){

	var c_inter Circle
	c_inter = C1{3}
	//c_inter := C1{3}      //this will work as well
	fmt.Println("Circum : ", c_inter.circumpherence())
	fmt.Println("Circum : ", c_inter.area())

	fmt.Println("The dynamic value of interface Circle is :", c_inter)
	fmt.Printf("The dynamic type of the interface is : %T", c_inter)
	fmt.Println()
}