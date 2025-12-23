package main

import "fmt"

type ListNode struct {
	data int
	next *ListNode
}

type LinkedList struct {
	head *ListNode
}

func(ll *LinkedList) FindMiddle(){
	fast := ll.head
	slow := ll.head

	for fast.next != nil &&	fast.next.next != nil {
		fast = fast.next.next
		slow = slow.next
	}

	fmt.Printf("%d", slow.data)
	fmt.Println()
}


func main(){

	ll := &LinkedList{}
	ll.head = &ListNode{data:10}
	ll.head.next = &ListNode{data:20}
	ll.head.next.next = &ListNode{data:30}
	ll.head.next.next.next = &ListNode{data:40}
	ll.head.next.next.next.next = &ListNode{data:50}
	ll.FindMiddle()


	ll1 := &LinkedList{}
	ll1.head = &ListNode{data:11}
	ll1.head.next = &ListNode{data:21}
	ll1.head.next.next = &ListNode{data:31}
	ll1.head.next.next.next = &ListNode{data:41}
	ll1.head.next.next.next.next = &ListNode{data:51}
	ll1.head.next.next.next.next.next = &ListNode{data:61}
	ll1.FindMiddle()
	
}