package main

import "fmt"

func main() {
	// TestHeap()
	TestMedianFinder()
}

func TestMedianFinder() {
	f := MedianFinder{}
	nums := []int{-1, -2, -3, -4}
	for _, n := range nums {
		f.AddNum(n)
		fmt.Println(f.FindMedian())
		fmt.Println()
	}
	f.Display()
}

func TestHeap() {
	h := NewHeap()
	nums := []int{-1, -2, -3, -4}
	for _, n := range nums {
		h.InsertMax(n)
	}
	h.SortAsc()
	h.Display()
}

func TestQueueUsingLinkedList() {
	q := NewQueueUsingLinkedList(10)

	for i := 1; i <= 20 && !q.IsFull(); i++ {
		_ = q.Enqueue(i)
	}

	fmt.Println(q.IsFull())
	fmt.Println(q.Peek())

	for !q.IsEmpty() {
		d, _ := q.Dequeue()
		fmt.Println(d)
	}
}

func TestQueue() {
	q := NewQueue(11)

	for i := 1; i <= 10 && !q.IsFull(); i++ {
		q.Enqueue(i)
	}

	for !q.IsEmpty() {
		n, _ := q.Dequeue()
		fmt.Println(n)
	}

	fmt.Println(q.IsFull())
}
