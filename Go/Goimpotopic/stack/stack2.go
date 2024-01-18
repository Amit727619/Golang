package main 

import "fmt"

type Stack []int 

func (st *Stack) Push(v int){
	*st = append(*st, v)
}
func main() {
	st := Stack{}
    st.Push(1)
	st.Push(2)


}