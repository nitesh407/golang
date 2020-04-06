package main
import (
	"fmt"
	//"reflect"
)
/*
type Car struct{

	name string
	prize int 
}
*/
type Student struct{
	name string
	class int
}
type School struct{
	details Student
}
func main(){

	/*c1 := Car{name: "BMW", prize: 2000}

	c2 := Car{"VoksWogen", 1000}
	c2.prize = 1
	fmt.Println(c1, c2)

	//struct with pionters
	c3:= &Car{ "Maruti", 200}
	fmt.Println((*c3).name)

	if c1 == c2{
		fmt.Println("c1 and c2 are equal")
	}else{
		fmt.Println("c1 and c2 arent equal")
	}
	fmt.Println("C1 and c3 are equal : ", reflect.DeepEqual(c1, c3))
	*/
	//now we gonna see the concept of nested structs
	s := School{
		details: Student{"Nitesh", 12},
	}
	fmt.Println(s.details.name)

	// anonymus structure
	a := struct{
		name string
		class int
	}{
		name: "Nitesh",
		class :12,
	}

	fmt.Println("Anon struct : ",a)

}
