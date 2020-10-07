package main

import (
	"fmt"
)

// Node is a structure
type Node struct {
	item string
	next *Node
}

type linkedList struct {
	head *Node
	size int
}

func (p *linkedList) addNode(name string) error {
	newNode := &Node{
		item: name,
		next: nil,
	}
	if p.head == nil {
		p.head = newNode
	} else {
		currentNode := p.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = newNode
	}
	p.size++
	return nil
}

func (p *linkedList) printAllNodes() error {
	currentNode := p.head
	if currentNode == nil {
		fmt.Println("Linked list is empty.")
		return nil
	}
	fmt.Printf("%+v\n", currentNode.item)
	for currentNode.next != nil {
		currentNode = currentNode.next
		fmt.Printf("%+v\n", currentNode.item)
	}

	return nil
}

// Remove Node
func (p *linkedList) Remove(item string) error {
	currentNode := p.head
	for currentNode != nil {
		if currentNode.next.item == item {
			currentNode.next = currentNode.next.next
			return nil
		}
		currentNode = currentNode.next
	}
	return nil
}

// add at Position
func (p *linkedList) addPos(item string, position int) error {
	if position < 0 || position > p.size {
		return fmt.Errorf("Index out of bounds")
	}
	addNode := Node{item, nil}
	if position == 0 {
		addNode.next = p.head
		p.head = &addNode
		return nil
	}
	node := p.head
	j := 0
	for j < position-2 {
		j++
		node = node.next
	}
	addNode.next = node.next
	node.next = &addNode
	p.size++
	return nil
}

// get function
func (p *linkedList) get(index int) (string, error) {
	currentNode := p.head
	if p.head == nil {
		return "", fmt.Errorf("Empty Linked list")
	}
	if index > 0 && index <= p.size {
		for i := 1; i <= index-1; i++ {
			currentNode = currentNode.next
		}
		item := currentNode.item
		return item, nil
	}
	return "", fmt.Errorf("Invalid index")
}

func main() {
	myList := &linkedList{nil, 0}
	fmt.Println("Created Linked List")
	fmt.Println()

	fmt.Print("Adding nodes to the linked list...\n\n")
	myList.addNode("Mary")
	myList.addNode("Jaina")
	myList.addNode("Gabriel")
	myList.addNode("Rabbya")
	fmt.Println("Showing all nodes in the linked list...")
	myList.printAllNodes()
	fmt.Printf("There are %+v elements in the list in total. \n", myList.size)
	fmt.Println()

	fmt.Println("Removing Nodes...")
	myList.Remove("Gabriel")
	fmt.Println()
	myList.printAllNodes()

	fmt.Println("Add Node at Position..")
	myList.addPos("Gabriel", 2)
	fmt.Println()
	myList.printAllNodes()
}
