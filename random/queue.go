package main

import "errors"

type Queue struct {
	DB   []int
	Size int
}

func NewQueue(size int) *Queue {
	return &Queue{
		DB:   []int{},
		Size: size,
	}
}

func (q *Queue) Enqueue(el int) error {
	if q.Len()+1 > q.Size {
		return errors.New("Queue is full")
	}
	q.DB = append(q.DB, el)
	return nil
}

func (q *Queue) Len() int {
	return len(q.DB)
}

func (q *Queue) Dequeue() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("Queue is empty")
	}

	item := q.DB[0]
	q.DB = q.DB[1:]
	return item, nil
}

func (q *Queue) IsEmpty() bool {
	return q.Len() == 0
}

func (q *Queue) IsFull() bool {
	return q.Size == q.Len()
}

func (q *Queue) Peek() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("Queue is empty")
	}
	return q.DB[0], nil
}
