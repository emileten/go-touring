package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T comparable] struct {
	next *List[T]
	val  T
}

// if nothing next, returns 'zero' value of T.
func (l List[T]) get() T {
	return l.val
}

func (l *List[T]) insert(newVal T) {
	// 5 lines to do a deep copy... There is probably a better way ! 
	copyList := new(List[T])
	copyNext := l.next
	copyVal := l.val
	copyList.next = copyNext
	copyList.val = copyVal
	l.next = copyList
	l.val = newVal
}

func (l *List[T]) remove() {
	// 5 lines to do a deep copy..
	copyList := new(List[T])
	copyNext := l.next
	copyVal := l.val
	copyList.next = copyNext
	copyList.val = copyVal
	l.val = copyList.next.val
	l.next = copyList.next.next
}

func main() {
	var myLinkedList = new(List[string])
	// prints nothing
	fmt.Print(myLinkedList.get())
	myLinkedList.insert("hyerin\n")
	// prints hyerin
	fmt.Print(myLinkedList.get())
	myLinkedList.remove()
	// prints nothing
	fmt.Print(myLinkedList.get())
	myLinkedList.insert("hyerin\n")
	// prints hyerin
	fmt.Print(myLinkedList.get())	
	myLinkedList.insert("emile\n")
	// prints emile
	fmt.Print(myLinkedList.get())
	myLinkedList.remove()
	// prints hyerin
	fmt.Print(myLinkedList.get())
}