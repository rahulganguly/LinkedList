package main

import "fmt"

type DoublyLinkedListNode struct {
	data int
	next *DoublyLinkedListNode
	prev *DoublyLinkedListNode
}

type DoublyLinkedList struct {
	head *DoublyLinkedListNode
	tail *DoublyLinkedListNode
	size int
}

func(dll *DoublyLinkedList)DisplayFront(){
	current := dll.head
	for current != nil {
		fmt.Printf("%d->", current.data)
		current = current.next
	}
	fmt.Println()
}

func(dll *DoublyLinkedList)DisplayRear(){
	current := dll.tail
	for current != nil {
		fmt.Printf("%d->",current.data)
		current = current.prev
	}
	fmt.Println()
}

func(dll *DoublyLinkedList)InsertAtHead(data int){
	newNode := &DoublyLinkedListNode{data:data}
	if dll.head == nil {
		dll.head = newNode
		dll.tail = newNode
	} else {
		newNode.next = dll.head
		dll.head.prev = newNode
		dll.head = newNode
	}
	dll.size++
}

func(dll *DoublyLinkedList)InsertAtTail(data int){
	if dll.head == nil {
		dll.InsertAtHead(data)
	} else {
		newNode := &DoublyLinkedListNode{data:data}
		dll.tail.next = newNode
		newNode.prev = dll.tail
		dll.tail = newNode
		dll.size++	
	}
}

func(dll *DoublyLinkedList)InsertAtPosition(data int, position int){
	if position == 1 {
		dll.InsertAtHead(data)
		return
	} else if position > dll.size {
		dll.InsertAtTail(data)
		return
	} else {
		newNode := &DoublyLinkedListNode{data:data}
		current := dll.head
		
		for pos := 1; pos < position ; pos++ {		
			current = current.next
		}

		current.prev.next = newNode
		newNode.prev = current.prev		
		newNode.next = current
		current.prev = newNode
		dll.size++
	}
}

func(dll *DoublyLinkedList)DeletePosition(position int){
	if position == 1 {
		dll.DeleteHead()
		return
	} else if position > dll.size {
		dll.DeleteTail()
		return
 	} else {
 		current := dll.head
 		for pos := 1; pos < position-1; pos++{
 			current = current.next
 		}

 		current.next = current.next.next
 		current.next.prev = current
 		dll.size--
 	}
}

func(dll *DoublyLinkedList)DeleteHead(){
	dll.head = dll.head.next
	dll.head.prev = nil
	dll.size--
}

func(dll *DoublyLinkedList)DeleteTail(){
	dll.tail = dll.tail.prev
	dll.tail.next = nil
	dll.size--
}



func main(){
	ll := &DoublyLinkedList{}
	ll.InsertAtHead(40)
	ll.InsertAtHead(30)
	ll.InsertAtHead(20)
	ll.InsertAtHead(10)
	ll.DisplayFront()
	ll.DisplayRear()
	ll.InsertAtTail(70)
	ll.InsertAtTail(80)
	ll.InsertAtTail(90)
	ll.DisplayFront()
	ll.DisplayRear()
	ll.InsertAtPosition(50,5)
	ll.InsertAtPosition(60,6)
	ll.DisplayFront()
	ll.DisplayRear()
	ll.DeleteHead()
	ll.DeleteHead()
	ll.DisplayFront()
	ll.DisplayRear()
	ll.DeleteTail()
	ll.DeleteTail()
	ll.DisplayFront()
	ll.DisplayRear()
	ll.DeletePosition(2)
	ll.DeletePosition(2)
	ll.DisplayFront()
	ll.DisplayRear()
}
