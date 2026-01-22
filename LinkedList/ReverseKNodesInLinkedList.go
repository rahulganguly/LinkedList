package main

import "fmt"

type ListNode struct {
	data int
	next *ListNode
}

type LinkedList struct {
	head *ListNode
}

func(ll *LinkedList)ReverseKNodes(k int) *ListNode{
	if ll.head == nil || k < 1{
		return ll.head
	}
	
	dummy := &ListNode{next:ll.head}
	prevGroup := dummy

	for {
		kth := prevGroup
		for i :=0; i < k && kth !=nil ; i++ {
			kth = kth.next
		}

		if kth == nil {
			break
		}

		groupNext := kth.next
		prev,curr := kth.next,prevGroup.next


		for curr != groupNext {
			tmp := curr.next
			curr.next = prev
			prev = curr
			curr = tmp
		}

		tmp := prevGroup.next
		prevGroup.next = kth
		prevGroup = tmp 

	}
	return dummy.next
}



func main(){
	ll := &LinkedList{}
	ll.head = &ListNode{data:10}
	ll.head.next = &ListNode{data:20}
	ll.head.next.next = &ListNode{data:30}
	ll.head.next.next.next = &ListNode{data:40}
	ll.head.next.next.next.next = &ListNode{data:50}
	ll.head.next.next.next.next.next = &ListNode{data:60}
	ll.head.next.next.next.next.next.next = &ListNode{data:10}

	res1 := ll.ReverseKNodes(2)

	fmt.Println()
	for res1 != nil {
		fmt.Printf("%d->", res1.data)
		res1 = res1.next
	}
}