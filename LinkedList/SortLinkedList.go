package main

import "fmt"

type ListNode struct {
	data int
	next *ListNode
}

type LinkedList struct {
	head *ListNode
}

func Sort(head *ListNode) *ListNode {
	if head == nil || head.next == nil {
		return head
	}
	fast,slow := head,head
	for fast.next != nil && fast.next.next != nil {
		fast = fast.next.next
		slow = slow.next
	}
	firstTail := slow
	slow = slow.next
	firstTail.next = nil
	first,second := Sort(head),Sort(slow)
	return Merge(first,second)

}

func Merge(head1 *ListNode,head2 *ListNode) *ListNode {
	dummy := new(ListNode)

	for node := dummy; head1 != nil || head2 != nil; node = node.next{
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


func main() {
	ll := &LinkedList{}
	ll.head = &ListNode{data:10}
	ll.head.next = &ListNode{data:20}
	ll.head.next.next = &ListNode{data:5}
	ll.head.next.next.next = &ListNode{data:11}
	ll.head.next.next.next.next = &ListNode{data:22}
	response := Sort(ll.head)
	for response != nil {
		fmt.Printf("%d->",response.data)
		response = response.next
	}
}