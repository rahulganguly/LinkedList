package main

import "fmt"

type ListNode struct {
	data int
	next *ListNode
}

type LinkedList struct {
	head *ListNode
	size int
}


func(ll *LinkedList)Display(){
	current := ll.head
	for current != nil {
		fmt.Printf("%d->",current.data)
		current = current.next
	}
	fmt.Println()
}


func(ll *LinkedList)InsertAtHead(data int){
	newNode := &ListNode{data:data}
	newNode.next = ll.head
	ll.head = newNode
	ll.size++
}

func(ll *LinkedList)InsertAtTail(data int){
	newNode := &ListNode{data:data}
	if ll.head == nil {
		ll.head = newNode
	} else {
		current := ll.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
	ll.size++
}

func(ll *LinkedList)InsertAtPosition(data int,position int){
	if position == 1 {
		ll.InsertAtHead(data)
	} else if position > ll.size {
		ll.InsertAtTail(data)
	} else {
		newNode := &ListNode{data:data}
		current := ll.head
		for pos := 1 ; pos < position - 1 ; pos ++ {
			current = current.next			
		}
		newNode.next = current.next
		current.next = newNode
		ll.size++
	}
}

func(ll *LinkedList)DeleteHead(){
	if ll.head != nil {
		ll.head = ll.head.next
		ll.size --
	}	
}

func(ll *LinkedList)DeleteTail(){
	current := ll.head
	for current.next.next != nil {
		current = current.next		
	}
	current.next = nil
	ll.size--
}


func(ll *LinkedList)DeleteAtPosition(position int){
	if position == 1 {
		ll.DeleteHead()
	} else if position > ll.size {
		ll.DeleteTail()
	} else {
		current := ll.head
		for pos := 1; pos < position -1; pos++{
			current = current.next
		}
		current.next = current.next.next
		ll.size--
	}

}

func main(){

	ll := &LinkedList{}
	ll.InsertAtHead(30)
	ll.InsertAtHead(20)
	ll.InsertAtHead(10)
	ll.InsertAtTail(80)
	ll.InsertAtTail(90)
	ll.InsertAtTail(100)
	ll.InsertAtPosition(50,4)
	ll.InsertAtPosition(40,4)
	ll.InsertAtPosition(60,6)
	ll.Display()
	ll.DeleteHead()
	ll.DeleteHead()
	ll.Display()
	ll.DeleteTail()
	ll.DeleteTail()
	ll.Display()
	ll.DeleteAtPosition(2)
	ll.DeleteAtPosition(2)
	ll.Display()
}