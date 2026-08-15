package main

import "fmt"

type LinkList struct {
	Head *Node
}
type Node struct {
	Data int8
	Next *Node
}

func main() {
	ll := &LinkList{}

	appendll(ll, 2)
	appendll(ll, 3)
	prependll(ll, 1)
	prependll(ll, 4)

	node := ll.Head
	for node != nil {
		fmt.Println("linklist:", node.Data)
		node = node.Next
	}
	var i int8 = 0
	value := valAtIndexll(ll, i)
	fmt.Printf("index: %d value: %d\n", i, value)
	i = 2
	value = valAtIndexll(ll, i)
	fmt.Printf("index: %d value: %d\n", i, value)


	n := delAtIndexll(ll.Head, 1)
	for n != nil {
		fmt.Println("linklist2:", n.Data)
		n = n.Next
	}
}

func appendll(ll *LinkList, data int8) {
	// create a new node and set Next a nil 
	// if Next nil means: end of the LinkList
	newNode := &Node{Data: data, Next: nil}

	if ll.Head == nil {
		ll.Head = newNode; return
	}

	// loop through until last non-null node 
	currentNode := ll.Head
	for currentNode.Next != nil {
		currentNode = currentNode.Next
	}

	currentNode.Next = newNode
}

func prependll(ll *LinkList, data int8) { 
	existingHead := ll.Head

	if existingHead != nil {
		// pad the existing head node in new head node 
		newHead := &Node{Data: data, Next: existingHead}
		ll.Head = newHead
	} else {
		ll.Head = &Node{Data: data, Next: nil}; return
	}
}

func valAtIndexll(ll *LinkList, index int8) int8 {
	currentNode := ll.Head

	var value int8 = currentNode.Data // return head value on index zero

	for range index {
		// move next node
		currentNode = currentNode.Next
		// change to next node's value
		value = currentNode.Data
	}

	return value
}

func delAtIndexll(node *Node, value int8) *Node {
	head := node
	for head.Next != nil {
		if head.Next.Data == value {
			head.Next = head.Next.Next
			return head
		}
		head = head.Next
	}
	return head
}
