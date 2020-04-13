package main
import "fmt"

type Tree struct{
	node *Node
}

func (t * Tree)insert(val int)*Tree{

	if t.node == nil{
		node := &Node{value : val}
		t.node = node
	}else{
		t.node.insert(val)
	}
	return t
}
type Node struct{

	value int
	Left *Node
	Right *Node
}

func (n *Node)insert(val int){

	if val <= n.value{
		if n.Left == nil{
			n.Left = &Node{value : val}
		}else{
			n.Left.insert(val)
		}
	}else{
		if n.Right == nil{
			n.Right = &Node{value : val}
		}else{
			n.Right.insert(val)
		}
	}
}

func Print_Node(n *Node){

	if n == nil{
		return
	}
	fmt.Println(n.value)
	Print_Node(n.Left)
	Print_Node(n.Right)
}

func (n *Node)exists(val int)bool{

	if n == nil{
		return false
	}
	if n.value == val{
		return true
	}
	if val <= n.value{
		return n.Left.exists(val)
	}else{
		return n.Right.exists(val)
	}
}

func main(){

	t := &Tree{}
	t.insert(10)
	t.insert(8)
	t.insert(20)
	t.insert(9)
	t.insert(0)
	t.insert(15)
	t.insert(25)

	Print_Node(t.node)

	fmt.Println(t.node.exists(10))
}	