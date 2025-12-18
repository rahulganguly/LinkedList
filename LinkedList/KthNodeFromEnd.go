package main

import "fmt"

type ListNode struct {
	data int
	next *ListNode
}

type LinkedList struct {
	head *ListNode
}

func(ll *LinkedList)FindKthNodeFromEnd(endNodePos int){
	current := ll.head
	startNode := ll.head
	for position := 1; position <= endNodePos;position++{
		current = current.next
	} 
	for current != nil {
		startNode = startNode.next
		current = current.next
	}
	fmt.Printf("%d", startNode.data)

}



func main(){
	ll := &LinkedList{}
	ll.head = &ListNode{data:10}
	ll.head.next = &ListNode{data:20}
	ll.head.next.next = &ListNode{data:30}
	ll.head.next.next.next = &ListNode{data:40}
	ll.head.next.next.next.next = &ListNode{data:50}
	ll.head.next.next.next.next.next = &ListNode{data:60}
	ll.head.next.next.next.next.next.next = &ListNode{data:70}
	ll.FindKthNodeFromEnd(2)
}