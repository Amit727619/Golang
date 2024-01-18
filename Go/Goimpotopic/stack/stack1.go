package main

import "fmt"

// Stack is a type representing a stack of integers.
type Stack []int

// Push adds a value to the top of the stack.
func (st *Stack) Push(v int) {
	*st = append(*st, v)
}

// Pop removes and returns the top value from the stack.
func (st *Stack) Pop() int {
	
	top := (*st)[len(*st)-1]
	*st = (*st)[:len(*st)-1]
	return top
}


func main() {
	// Create an empty stack
	st := Stack{}

	// Push values onto the stack
	st.Push(1)
	st.Push(2)
    fmt.Println("push on stack",st)
	// Pop values from the stack and print them
	fmt.Println("pop from stack")
	fmt.Println(st.Pop())
	fmt.Println(st.Pop())

}
