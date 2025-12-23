package main

import "fmt"

type ListNode struct {
	data int
	next *ListNode
}

type LinkedList struct {
	head *ListNode
}

func findIntersectionNode(head1,head2 *ListNode){
	len1 := findLength(head1)
	len2 := findLength(head2)

	if len1 > len2 {		
		for ;len1 > len2; len1-- {
			head1 = head1.next
		}
	} else if len2 > len1 {
		for ;len2 > len1; len2--{
			head2 = head2.next
		}
	}

	for head1 != head2 {
		head1 = head1.next
		head2 = head2.next
	}

	fmt.Printf("%d",head1.data)
}

func findLength(head *ListNode) int {
	current := head
	counter := 0
	for current != nil {
		current = current.next
		counter++
	}
	return counter
}

func main(){
	newNode := &ListNode{data:111}
	newNode.next = &ListNode{data:211}
	newNode.next.next = &ListNode{data:311}

	ll := &LinkedList{}
	ll.head = &ListNode{data:10}
	ll.head.next = &ListNode{data:20}
	ll.head.next.next = &ListNode{data:30}
	ll.head.next.next.next = &ListNode{data:40}
	ll.head.next.next.next.next = newNode


	ll1 := &LinkedList{}
	ll1.head = &ListNode{data:11}
	ll1.head.next = &ListNode{data:21}
	ll1.head.next.next = &ListNode{data:31}
	ll1.head.next.next.next = &ListNode{data:41}
	ll1.head.next.next.next.next = &ListNode{data:51}
	ll1.head.next.next.next.next.next = &ListNode{data:61}
	ll1.head.next.next.next.next.next.next = newNode
	findIntersectionNode(ll.head,ll1.head)
}