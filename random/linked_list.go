package main

import (
	"errors"
	"fmt"
)

type Node struct {
	Data int
	Next *Node
}

type QueueUsingLinkedList struct {
	head *Node
	tail *Node
	Size int
	Len  int
}

func NewQueueUsingLinkedList(size int) *QueueUsingLinkedList {
	return &QueueUsingLinkedList{
		Size: size,
	}
}

func (q *QueueUsingLinkedList) Enqueue(el int) error {
	if q.IsFull() {
		return errors.New("Queue is full")
	}

	node := &Node{
		Data: el,
	}

	if q.head == nil {
		q.head = node
		q.tail = node
	}

	q.tail.Next = node
	q.tail = node
	q.Len += 1
	return nil
}

func (q *QueueUsingLinkedList) Dequeue() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("Queue is empty")
	}
	d := q.head.Data
	q.head = q.head.Next
	q.Len -= 1
	return d, nil
}

func (q *QueueUsingLinkedList) IsEmpty() bool {
	return q.Len == 0
}

func (q *QueueUsingLinkedList) IsFull() bool {
	return q.Size == q.Len
}

func (q *QueueUsingLinkedList) Peek() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("queue is empty")
	}

	return q.head.Data, nil
}

func (q *QueueUsingLinkedList) Display() {
	if q.IsEmpty() {
		fmt.Println("queue is empty")
	}

	current := q.head

	for current != nil {
		fmt.Println(current.Data)
		current = current.Next
	}
}
