package main
import "fmt"

type Queue struct{
	items []int
}
func (q *Queue)Enqueue(val int){
	q.items = append(q.items, val)
}

func (q *Queue)Dequeue()int{
	if len(q.items) == 0{
		fmt.Println("Queue is empty")
		return -1
	}else{
		a := q.items[0]
		q.items = q.items[1:]
		return a
	}
}

func main(){
	q := &Queue{}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	fmt.Println(q)
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q)
}