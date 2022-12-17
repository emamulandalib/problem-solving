package main

import "fmt"

type Heap struct {
	Db []int
}

func NewHeap() *Heap {
	return &Heap{}
}

func (h *Heap) InsertMax(el int) {
	h.Db = append(h.Db, el)
	i := h.len() - 1

	for i > 0 && h.Db[i] > h.Db[h.parent(i)] {
		h.swap(i, h.parent(i))
		i = h.parent(i)
	}
}

func (h *Heap) InsertMin(el int) {
	h.Db = append(h.Db, el)
	i := h.len() - 1

	for i > 0 && h.Db[i] < h.Db[h.parent(i)] {
		h.swap(i, h.parent(i))
		i = h.parent(i)
	}
}

func (h *Heap) SortAsc() {
	for i := h.len() - 1; i >= 0; i-- {
		h.swap(0, i)
		h.maxHeapify(i, 0)
	}
}

func (h *Heap) SortDesc() {
	for i := h.len() - 1; i >= 0; i-- {
		h.swap(0, i)
		h.minHeapify(i, 0)
	}
}

func (h *Heap) Display() {
	fmt.Println(h.Db)
}

func (h *Heap) parent(i int) int {
	return (i - 1) / 2
}

func (h *Heap) minHeapify(n int, i int) {
	smallest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && h.Db[left] < h.Db[smallest] {
		smallest = left
	}

	if right < n && h.Db[right] < h.Db[smallest] {
		smallest = right
	}

	if smallest != i {
		h.swap(i, smallest)
		h.minHeapify(n, smallest)
	}
}

func (h *Heap) maxHeapify(n int, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && h.Db[left] > h.Db[largest] {
		largest = left
	}

	if right < n && h.Db[right] > h.Db[largest] {
		largest = right
	}

	if largest != i {
		h.swap(i, largest)
		h.maxHeapify(n, largest)
	}
}

func (h *Heap) len() int {
	return len(h.Db)
}

func (h *Heap) swap(i, j int) {
	h.Db[i], h.Db[j] = h.Db[j], h.Db[i]
}

type MedianFinder struct {
	Db []int
}

func Constructor() *MedianFinder {
	return &MedianFinder{}
}

func (this *MedianFinder) AddNum(num int) {
	this.Db = append(this.Db, num)
}

func (this *MedianFinder) FindMedian() float64 {
	this.sortAsc()
	if this.isEven() {
		l := this.len() / 2
		f := l - 1
		return float64(this.Db[f]+this.Db[l]) / float64(2)
	} else {
		m := this.len() / 2
		return float64(this.Db[m])
	}
}

func (this *MedianFinder) isEven() bool {
	return this.len()%2 == 0
}

func (this *MedianFinder) sortAsc() {
	this.buildMax()
	for i := this.len() - 1; i >= 0; i-- {
		this.swap(0, i)
		this.heapify(i, 0)
	}
}

func (this *MedianFinder) buildMax() {
	for i := this.len() / 2; i >= 0; i-- {
		this.heapify(this.len(), i)
	}
}

func (this *MedianFinder) heapify(n, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && this.Db[left] > this.Db[largest] {
		largest = left
	}

	if right < n && this.Db[right] > this.Db[largest] {
		largest = right
	}

	if largest != i {
		this.swap(i, largest)
		this.heapify(n, largest)
	}
}

func (this *MedianFinder) swap(i, j int) {
	this.Db[i], this.Db[j] = this.Db[j], this.Db[i]
}

func (this *MedianFinder) len() int {
	return len(this.Db)
}

func (this *MedianFinder) Display() {
	fmt.Println(this.Db)
}
