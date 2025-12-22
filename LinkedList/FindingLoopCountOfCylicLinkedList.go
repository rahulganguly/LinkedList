package main

import "fmt"

type ListNode struct {
	data int
	next *	ListNode
}

type LinkedList struct {
	head *ListNode
}

func(ll *LinkedList)GetCyclicLinkedListCount(){
	fast := ll.head
	slow := ll.head
	isCyclic := false
	for fast != nil && fast.next != nil {
		fast = fast.next.next
		slow = slow.next
		if fast == slow {
			isCyclic = true
			break
		}
	}

	counter := 1
	if isCyclic {		
		for fast.next != slow {
			counter++
			fast = fast.next
		}
	}

	fmt.Printf("%d", counter)
}



func main(){
	newNode := &ListNode{data:100}
	ll := &LinkedList{}
	ll.head = &ListNode{data:10}
	ll.head.next = &ListNode{data:20}
	ll.head.next.next = &ListNode{data:30}
	ll.head.next.next.next = &ListNode{data:40}
	ll.head.next.next.next.next = newNode
	newNode.next = &ListNode{data:50}
	newNode.next.next = &ListNode{data:60}
	newNode.next.next.next = &ListNode{data:70}
	newNode.next.next.next.next = newNode

	ll.GetCyclicLinkedListCount()

}