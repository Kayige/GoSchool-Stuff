package main

import (
	"errors"
	"fmt"
)

// Queue Struct Global
type Queue struct {
	front *Node
	back  *Node
	size  int
}

func (p *Queue) enqueue(name string, priority int) error {
	newNode := &Node{}
	if p.front == nil {
		p.front = newNode
		p.back = newNode
	} else {
		if p.front.priority > priority {
			newNode.next = p.front

		} else {
			currentNode := p.front

			for currentNode.next != nil && currentNode.next.priority <= priority {
				currentNode = currentNode.next
			}
			newNode.next = currentNode.next
			currentNode.next = newNode
		}

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
