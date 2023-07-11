package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l linkedList) printListData() {
	toPrint := l.head
	for l.length > 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Println()
}

func (l *linkedList) deleteWithValue(val int) {
	if l.length == 0 {
		return
	}

	if l.head.data == val {
		l.head = l.head.next
		l.length--
	}

	previousToDelete := l.head
	for previousToDelete.next.data != val {
		if previousToDelete.next.next == nil {
			return
		}
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	l.length--
}

func main() {
	myList := linkedList{}
	node1 := &node{data: 21}
	node2 := &node{data: 32}
	node3 := &node{data: 37}
	node4 := &node{data: 12}
	node5 := &node{data: 82}
	node6 := &node{data: 50}
	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)
	myList.prepend(node4)
	myList.prepend(node5)
	myList.prepend(node6)
	// fmt.Println(myList)
	myList.printListData()
	myList.deleteWithValue(32)
	myList.printListData()

	// data in linkedlist
	myList.printListData()
	myList.deleteWithValue(100)

	// deleting header
	myList.printListData()
	myList.deleteWithValue(50)

	// deleting empty list
	emptyList := linkedList{}
	myList.printListData()
	emptyList.deleteWithValue(1)
}
