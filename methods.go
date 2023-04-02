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

//pointer type reciever method
func (s *Student)Update_name(name string){                 // it has to be pointer type, to change the internal values of the struct
	s.name = name
}

func main(){

	st := Student{"Nitesh", 21}
	fmt.Println(st)
	st.show()
	st.Update_name("Pawan")
	fmt.Println("After updation :", st)

	ad := Address{83, "Keshav colony"}
	ad.show()


}