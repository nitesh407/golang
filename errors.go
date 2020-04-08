package main
import (
	"fmt"
	"errors"
)

type Emp struct{
	name, role string
}

func (e *Emp)get_salary(role string)(int, error){         //error should be the last argument, its good practice
	if role==""{
		return 0, errors.New("There is not such role")
	}
	switch role{
	case "Develpoer":
		return 24500, nil
	case "Test Engineer":
		return 30000, nil
	default:
		//we can be very precise by using the nested function
		return 0, errors.New(fmt.Sprintf("There isn't such role: %v in the company", role))
	}
}

func main(){

	emp := Emp{"Nitesh", ""}
	if s, err := emp.get_salary(emp.role); err != nil{
		fmt.Println(err)
	}else{
		fmt.Println("The salary is :", s)
	}
}