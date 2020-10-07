package main

import (
	"errors"
	"fmt"
)

// Node Struct - Local
type Node struct {
	item string
	next *Node
}

// Queue Struct Global
type Queue struct {
	front *Node
	back  *Node
	size  int
}

func (p *Queue) enqueue(name string) error {
	newNode := &Node{
		item: name,
		next: nil,
	}
	if p.front == nil {
		p.front = newNode

	} else {
		p.back.next = newNode

	}
	p.back = newNode
	p.size++
	return nil
}

func (p *Queue) dequeue() (string, error) {
	var item string

	if p.front == nil {
		return "", errors.New("empty queue")
	}

	item = p.front.item
	if p.size == 1 {
		p.front = nil
		p.back = nil
	} else {
		p.front = p.front.next
	}
	p.size--
	return item, nil
}

func (p *Queue) printAllNodes() error {
	currentNode := p.front
	if currentNode == nil {
		fmt.Println("Queue is empty.")
		return nil
	}
	fmt.Printf("%+v\n", currentNode.item)
	for currentNode.next != nil {
		currentNode = currentNode.next
		fmt.Printf("%+v\n", currentNode.item)
	}

	return nil
}

func (p *Queue) isEmpty() bool {
	return p.size == 0
}

func createQueue() *Queue {
	myQueue := &Queue{}
	myQueue.enqueue("Gabriel")
	myQueue.enqueue("Ada")
	myQueue.enqueue("Jaina")
	myQueue.enqueue("Proudmoore")
	return myQueue
}

func printContent(p *Queue) {
	p.printAllNodes()
}

func main() {
	myQueueT := createQueue()

	printContent(myQueueT)
	fmt.Println()

	myQueueT.dequeue()
	myQueueT.printAllNodes()
}
