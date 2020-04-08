package main
import "fmt"

type List struct{
	head *Node
	tail *Node
}

func (l *List)First()*Node{
	return l.head
}

func (l *List)Last()*Node{
	return l.tail
}

func (l *List)Push(value int){
	node := &Node{value: value}

	if l.head == nil{
		l.head = node
	}else{
		l.tail.next = node
		node.prev = l.tail
	}
	l.tail = node
}

type Node struct{
	value int 
	next *Node
	prev *Node
}

func (n *Node)Next()*Node{
	return n.next
}
func (n *Node)Prev()*Node{
	return n.prev
}

func main(){
	l := &List{}
	l.Push(1)
	l.Push(2)
	l.Push(3)
	l.Push(4)

	n := l.First()
	//traversing
	for{
		fmt.Println(n)
		n = n.Next()
		if n == nil{
			break
		}
	}

	//printing the list in reverse order
	n1 := l.Last()
	for{
		fmt.Println(n1)
		n1 = n1.Prev()
		if n1 == nil{
			break
		}
	}

}