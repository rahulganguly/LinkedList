package main

import "fmt"

type ListNode struct {
	data int
	next *ListNode
}

type LinkedList struct {
	head *ListNode
}

func MergeKLists(ll []*ListNode) *ListNode {

	if ll == nil || len(ll) == 0 {
		return nil
	}
	for len(ll) > 1 {
		 ll1 := ll[0]
		 ll2 := ll[1]
		 ll = ll[2:]
		 merged := MergeList(ll1,ll2)
		 ll = append(ll,merged)
	}
	return ll[0]
}

func MergeList(ll1 *ListNode,ll2 *ListNode) *ListNode {
	dummy := new(ListNode)

	for node := dummy; ll1 != nil || ll2 != nil; node = node.next {
		if ll1 == nil {
			node.next = ll2
			break
		} else if ll2 == nil {
			node.next = ll1
			break
		} else if ll1.data < ll2.data {
			node.next = ll1
			ll1 = ll1.next
		} else {
			node.next = ll2
			ll2 = ll2.next
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



	ll2 := &LinkedList{}
	ll2.head = &ListNode{data:12}
	ll2.head.next = &ListNode{data:22}
	ll2.head.next.next = &ListNode{data:32}
	ll2.head.next.next.next = &ListNode{data:42}
	ll2.head.next.next.next.next = &ListNode{data:52}

	ll3 := &LinkedList{}
	ll3.head = &ListNode{data:11}
	ll3.head.next = &ListNode{data:21}
	ll3.head.next.next = &ListNode{data:31}
	ll3.head.next.next.next = &ListNode{data:41}
	ll3.head.next.next.next.next = &ListNode{data:51}

	var listNodes []*ListNode
	listNodes = append(listNodes,ll.head,ll1.head,ll2.head,ll3.head)

	result := MergeKLists(listNodes)
	for result != nil {
		fmt.Printf("%d->", result.data)
		result = result.next
	}
}