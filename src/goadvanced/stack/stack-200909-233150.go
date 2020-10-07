package main

import (
	"errors"
	"fmt"
)

// Node struct
type Node struct {
	item string
	next *Node
}

type stack struct {
	top  *Node
	size int
}

func (p *stack) push(name string) error {
	newNode := &Node{
		item: name,
		next: nil,
	}
	if p.top == nil {
		p.top = newNode
	} else {
		newNode.next = p.top
		p.top = newNode
	}
	p.size++
	return nil
}

func (p *stack) pop() (string, error) {
	var item string

	if p.top == nil {
		return "", errors.New("Empty Stack")
	}

	item = p.top.item
	if p.size == 1 {
		p.top = nil
	} else {
		p.top = p.top.next
	}
	p.size--
	return item, nil
}

func (p *stack) printAllNodes() error {
	currentNode := p.top
	if currentNode == nil {
		fmt.Println("Stack is empty.")
		return nil
	}
	fmt.Printf("%+v\n", currentNode.item)
	for currentNode.next != nil {
		currentNode = currentNode.next
		fmt.Printf("%+v\n", currentNode.item)
	}
	return nil
}

var (
	parentheses = map[string]string{
		"(": ")",
		"{": "}",
		"[": "]",
	}
	openParentheses  = "({["
	closeParentheses = ")}]"
)

func main() {
	myStack := &stack{nil, 0}
	fmt.Println("Initializing Stack")
	fmt.Println()

	fmt.Println("Add in Stacks....")
	myStack.push("Gabriel")
	myStack.push("Ada")
	myStack.push("Kai")
	myStack.push("Rawr!")
	fmt.Println("Stacks added...")
	fmt.Println()

	fmt.Println("Printing all Nodes...")
	myStack.printAllNodes()
	fmt.Println()

	fmt.Println("Popping Top Node...")
	myStack.pop()

	fmt.Println("Printing all Nodes...")
	myStack.printAllNodes()
	fmt.Println()

	fmt.Println("Printing Stack without function")

	currentNode := myStack.top
	if currentNode == nil {
		fmt.Println("Stack is empty.")
	} else {
		fmt.Printf("%+v\n", currentNode.item)
		for currentNode.next != nil {
			currentNode = currentNode.next
			fmt.Printf("%+v\n", currentNode.item)
		}

	}
}
