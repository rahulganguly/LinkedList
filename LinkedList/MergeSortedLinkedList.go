package main

import "fmt"


type ListNode struct {
	data int
	next *ListNode
}

type LinkedList struct {
	head *ListNode
}

func MergeLinkedList(head1 *ListNode,head2 *ListNode) *ListNode {
	dummy := new(ListNode)	
	for node := dummy ;head1 != nil || head2 != nil; node = node.next {

		if head1 == nil {
			node.next = head2
			break
		} else if head2 == nil {
			node.next = head1
			break
		} else if head1.data < head2.data {
			node.next = head1
			head1 = head1.next	
		} else {
			node.next = head2
			head2 = head2.next
		}		
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


	ll1 := &LinkedList{}
	ll1.head = &ListNode{data:11}
	ll1.head.next = &ListNode{data:21}
	ll1.head.next.next = &ListNode{data:31}
	ll1.head.next.next.next = &ListNode{data:41}
	ll1.head.next.next.next.next = &ListNode{data:51}
	ll1.head.next.next.next.next.next = &ListNode{data:61}

	result := MergeLinkedList(ll.head,ll1.head)

	for result != nil {
		fmt.Printf("%d->", result.data)
		result = result.next
	}
		
}