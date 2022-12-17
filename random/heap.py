from operator import ne


class Heap():
    def __init__(self) -> None:
        self.heap = []
    
    def add(self, el):
        self.heap.append(el)
    
    def insert(self, el):
        self.heap.append(el)
        i = len(self.heap) - 1

        while i > 0 and self.heap[i] > self.heap[self.parent(i)]:
            self.heap[i], self.heap[self.parent(i)] = self.heap[self.parent(i)], self.heap[i]
            i = self.parent(i)
    
    def parent(self, i):
        return (i - 1) // 2
    
    def max_heapify(self, n, i):
        largest = i
        left = 2 * i + 1
        right = 2 * i + 2

        if left < n and self.heap[left] > self.heap[largest]:
            largest = left

        if right < n and self.heap[right] > self.heap[largest]:
            largest = right
        
        if largest != i:
            self.heap[i], self.heap[largest] = self.heap[largest], self.heap[i]
            self.max_heapify(n, largest)
    
    def build_max(self):
        n = len(self.heap)
        for i in range(n // 2, -1, -1):
            self.max_heapify(n, i)
    
    def sort_asc(self):
        self.build_max()

        n = len(self.heap)
        for i in range(n - 1, 0, -1):
            self.heap[0], self.heap[i] = self.heap[i], self.heap[0]
            self.max_heapify(i, 0)

    def min_heapify(self, n, i):
        largest = i
        left = 2 * i + 1
        right = 2 * i + 2

        if left < n and self.heap[left] < self.heap[largest]:
            largest = left

        if right < n and self.heap[right] < self.heap[largest]:
            largest = right
        
        if largest != i:
            self.heap[i], self.heap[largest] = self.heap[largest], self.heap[i]
            self.min_heapify(n, largest)
    
    def build_min(self):
        n = len(self.heap)
        for i in range(n // 2, -1, -1):
            self.min_heapify(n, i)
    
    def sort_desc(self):
        self.build_min()
        n = len(self.heap)
        for i in range(n - 1, -1, -1):
            self.heap[0], self.heap[i] = self.heap[i], self.heap[0]
            self.min_heapify(i, 0)

    def delete(self, el):
        self.build_max()
    
    def output(self):
        return self.heap

# h = Heap()
# for i in [3, 4, 2, 8, 5, 9, 1]:
#     h.add(i)

# h.build_max()
# h.insert(10)
# h.insert(100)
# # h.sort_desc()
# print(h.output())