package main
import "fmt"

type Student struct{
	name string
	age int
}

type Address struct{
	h_no int
	street string
}

func (s Student)show(){
	fmt.Println("Name : ", s.name)
	fmt.Println("Age : ", s.age)
}

func (a Address)show(){
	fmt.Println("House number : ", a.h_no)
	fmt.Println("House street : ", a.street)
}

func main(){

	st := Student{"Nitesh", 21}
	fmt.Println(st)
	st.show()

	ad := Address{83, "Keshav colony"}
	ad.show()
}