package main

import "fmt"

// Queue represents a basic queue data structure.
type Queue struct {
	items []int
}

// Enqueue adds an element to the end of the queue.
func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

// Dequeue removes and returns the element from the front of the queue.
func (q *Queue) Dequeue() (int) {
	
	item := q.items[0]
	q.items = q.items[1:]
	return item
}


// Size returns the number of elements in the queue.
func (q *Queue) Size() int {
	return len(q.items)
}

func main() {
	// Create a new queue
	q := Queue{}

	// Enqueue some elements
	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)

}
