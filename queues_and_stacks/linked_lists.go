package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

/*
 * Same idea in words:

 First attach the new node to the old tail.
 Then say “that new node is now the tail.”

 The first insert is special because there is no old tail to attach to,
 so both head and tail just become the new node.
*/

func (ll *LinkedList) insert(val int) {
	node := &Node{data: val}
	if ll.head == nil {
		ll.head = node
		ll.tail = node
	} else {
		ll.tail.next = node
		ll.tail = node
	}

	ll.length++

}

/*
From head, length-1 steps lands on the current tail, not the node before it.
List 1 -> 2 -> 3:

start at 1
i = 0 → 2
i = 1 → 3  (length-1 is 2, so two steps)

Then you set 3.next = nil (already nil) and tail = 3 (already tail).
Nothing is removed. You just return the last value and leave the list unchanged.
Other bugs
One-node list is not cleared.

If the list is only [1], the loop does not run, tail stays that node, head is never set to nil. The node is still in the list.
length is never decremented.

Even after you fix the walk, the next call uses a stale length. The loop can walk off the end and panic on new_tail.next when new_tail is nil.
-1 as “empty” is easy to confuse with real data.

Fine only if you promise values are never -1. (int, bool) or error is clearer.
The name hides what is deleted.

This is “remove last,” not a general delete.
*/

// func (ll *LinkedList) deletion() int {
// 	if ll.head == nil {
// 		return -1
// 	}
// 	node := ll.tail
// 	new_tail := ll.head
// 	for i := 0; i < ll.length-1; i++ {
// 		new_tail = new_tail.next
// 	}
// 	new_tail.next = nil
// 	ll.tail = new_tail

// 	return node.data
// }
//

// You need the predecessor of tail, plus two special cases: empty and single node.

/*
* Mental model
Deleting the tail is the reverse of append:

append: old tail’s next → new node, then move tail forward
delete tail: find node whose next is tail, set that next to nil, then move tail backward

You cannot start from tail and step back in a singly linked list. That is why you walk from head.
*/
func (ll *LinkedList) DeleteTail() (int, bool) {
	if ll.head == nil {
		return 0, false
	}

	val := ll.tail.data

	if ll.head == ll.tail {
		ll.head = nil
		ll.tail = nil
		ll.length = 0
		return val, true
	}

	prev := ll.head
	for prev.next != ll.tail {
		prev = prev.next
	}
	prev.next = nil
	ll.tail = prev
	ll.length--
	return val, true
}

// func (ll *LinkedList) traverse() bool {
// 	if ll.head == nil {
// 		return false
// 	}
// 	curr := ll.head

// 	for curr != nil {
// 		fmt.Printf("%d ", curr.data)
// 		curr = curr.next
// 	}
// 	return true
// }

func (ll *LinkedList) Traverse() {
	for curr := ll.head; curr != nil; curr = curr.next {
		fmt.Printf("%d ", curr.data)
	}
	fmt.Println()
}

// func (ll *LinkedList) isEmpty() bool {
// 	return ll.length == 0
// }

func (ll *LinkedList) IsEmpty() bool {
	return ll.head == nil
}

func (ll *LinkedList) size() int {
	return ll.length
}

func (ll *LinkedList) search(val int) bool {
	//  redundant check
	// if ll.head == nil {
	// 	return false
	// }

	for curr := ll.head; curr != nil; curr = curr.next {
		if curr.data == val {
			return true
		}
	}
	return false
}

func (ll *LinkedList) reverse() {
	if ll.head == nil || ll.length == 1 {
		return
	}

	// prev := ll.head
	// curr := prev.next

	// for curr != nil {
	// 	next := curr.next
	// 	curr.next = prev
	// 	// prev never moves
	// 	prev = curr
	// 	curr = next
	// }
	//

	ll.tail = ll.head

	var prev *Node
	curr := ll.head
	for curr != nil {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	ll.head = prev

	// head and tail are never updated
}

func ll_operations() {

}
