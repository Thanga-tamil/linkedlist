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
	prependll(ll, 0)

	node := ll.Head
	for node != nil {
		fmt.Println("linklist:", node.Data)
		node = node.Next
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

	if ll.Head != nil {
		newHead := &Node{Data: data, Next: ll.Head}
		ll.Head = newHead
	} else {
		ll.Head = &Node{Data: data, Next: nil}; return
	}

}
