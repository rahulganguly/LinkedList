package main

import "fmt"

type ListNode struct {
	data int
	next *ListNode
}

type LinkedList struct {
	head *ListNode
}

func DisplayReverse(head *ListNode){
	if head == nil {
		return
	}
	DisplayReverse(head.next)
	fmt.Printf("%d->",head.data)
}


func main(){
	ll := &LinkedList{}
	ll.head = &ListNode{data:10}
	ll.head.next = &ListNode{data:20}
	ll.head.next.next = &ListNode{data:30}
	ll.head.next.next.next = &ListNode{data:40}
	ll.head.next.next.next.next = &ListNode{data:50}
	DisplayReverse(ll.head)
}