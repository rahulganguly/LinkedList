package main

import "fmt"

type ListNode struct {
	data int
	next *ListNode
}

type LinkedList struct {
	head *ListNode
}

func(ll *LinkedList)Insert(data int){
	newNode := &ListNode{data:data}
	current := ll.head

	if current != nil && current.data > data {
		newNode.next = current
		ll.head = newNode
		return
	}

	current = ll.head
	for current.next != nil && current.next.data < data {
		current = current.next
	}
	newNode.next = current.next
	current.next = newNode
}

func(ll *LinkedList)Display() {
	current := ll.head
	for current != nil {
		fmt.Printf("%d->", current.data)
		current = current.next
	}
	fmt.Println()
}


func main() {
	ll := &LinkedList{}
	ll.head = &ListNode{data:10}
	ll.head.next = &ListNode{data:20}
	ll.head.next.next = &ListNode{data:30}
	ll.head.next.next.next = &ListNode{data:40}
	ll.Insert(35)
	ll.Display()
	ll.Insert(5)
	ll.Display()
}