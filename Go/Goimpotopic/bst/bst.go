package main
import "fmt"


type Node struct{
	Key int 
    Left *Node
	Right *Node  
}


func (n *Node) Insert(k int){

	if n.Key < k {
		
		if n.Right == nil{
			n.Right = &Node{Key: k}
		}else {
			n.Right.Insert(k)
		}

	}else if n.Key > k{
		
		if n.Left == nil{
			n.Left = &Node{Key: k}
		}else {
			n.Left.Insert(k)
		}

	}
}
func main(){
	tree := &Node{Key: 100}
	tree.Insert(200)
	tree.Insert(300)

    fmt.Println(tree)
}