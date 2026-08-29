package main

import "fmt"

// fixed size array implementation
// items is a dynamic array
// but queue capacity is tracked via maxCapacity
type Queue struct {
	items []int
}

func (q *Queue) dequeue() int {
	if q.isEmpty() {
		return -1
	}
	// fmt.Println(q.items, len(q.items), cap(q.items))
	n := q.items[0]
	// array slicing decreases len & cap of an array
	// q.items = q.items[1:]

	// copy(dst, src) copies from src into dst, starting at index 0.
	// It copies min(len(dst), len(src)) elements.
	// The front value 5 is gone from the used part.
	// The slice length is still 3, so that leftover 3 is still “in” the
	// slice. That’s why you need the next line.
	// doesn't allocate memory

	copy(q.items, q.items[1:])
	// fmt.Println(q.items, len(q.items), cap(q.items))

	// this only changes the length, not the array
	// len was 3, now 2. The slice becomes the first 2 slots.
	// The extra 3 at index 2 is still in memory, but it is
	// past len, so it is not part of our queue.
	//
	q.items = q.items[:len(q.items)-1]
	// fmt.Println(q.items, len(q.items), cap(q.items))
	return n
}

func (q *Queue) enqueue(val int) {
	if len(q.items) < cap(q.items) {
		// append method increases capacity of an array
		q.items = append(q.items, val)
	} else {
		return
	}

}

func (q Queue) peek() int {
	if !q.isEmpty() {
		return q.items[0]
	} else {
		return -1
	}
}

func (q *Queue) isEmpty() bool {
	return len(q.items) == 0
}

func (q Queue) isFull() bool {
	return len(q.items) == cap(q.items)
}

func QueueOperations() {
	// make takes 3 arguments here
	// type, len and capacity
	// len & cap functions can be used to check respective values
	q := Queue{make([]int, 0, 10)}
	q.enqueue(5)
	q.enqueue(4)
	q.enqueue(3)

	fmt.Println(q.items, len(q.items))

	fmt.Println(q.dequeue())
	fmt.Println(q.isEmpty())
	fmt.Println(q.isFull())
	fmt.Println(q.dequeue())
	fmt.Println(q.items, len(q.items))

	q.enqueue(10)
	q.enqueue(25)
	q.enqueue(13)
	fmt.Println(q.items, len(q.items))

	fmt.Println(q.dequeue())
	fmt.Println(q.items, len(q.items))

	fmt.Println(q.peek())
	fmt.Println(q.items, len(q.items))

}
