package main
import "fmt"

type Stack struct{
	items []int
}
func (s *Stack)Push(val int){
		s.items = append(s.items, val)
}
func (s *Stack)Pop()int{
	if len(s.items) == 0{
		fmt.Println("Stack Underflow")
		return -1
	}else{
		a := s.items[len(s.items) - 1]
		s.items = s.items[0:(len(s.items)-1)]
		return a
	}
	
}

func main(){
	s := &Stack{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Println(s)
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())

}