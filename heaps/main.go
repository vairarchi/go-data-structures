package main

import "fmt"

// MaxHeap struct has a slice that holds the array
type MaxHeap struct {
	array []int
}

// Insert adds an element to the heap
func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)
}

// Extract returns the largest keys and removes ir from the heap
func (h *MaxHeap) Extract() int {
	extracted := h.array[0]
	lastIndex := len(h.array) - 1
	// when the array is empty
	if len(h.array) == 0 {
		fmt.Println("Cannot extract because array length is 0")
		return -1
	}
	// take out the last index and put it in the root
	h.array[0] = h.array[lastIndex]
	h.array = h.array[:lastIndex]

	h.maxHeapifyUp(0)

	return extracted
}

// maxHeapify will heapify from bottom up
func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

// maxHeapifyDown will heapify from top to bottom
func (h *MaxHeap) maxHeapifyDown(index int) {
	lastIndex := len(h.array) - 1
	leftIndex, rightIndex := left(index), right(index)
	childToCompare := 0

	// loop while index has atleast one child
	for leftIndex < lastIndex {
		if leftIndex == lastIndex { // when index is the only child
			childToCompare = leftIndex
		} else if h.array[leftIndex] > h.array[rightIndex] { // when left child is larger
			childToCompare = leftIndex
		} else { // when right child is larger
			childToCompare = rightIndex
		}
		// compare array value of current index to larger child and swap if smaller
		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			leftIndex, rightIndex = left(index), right(index)
		} else {
			return
		}
	}
}

// get the parent index
func parent(index int) int {
	return (index - 1) / 2
}

// get the left child imdex
func left(index int) int {
	return 2*index + 1
}

// get the right child imdex
func right(index int) int {
	return 2*index + 2
}

// swap keys in the array
func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

func main() {
	m := &MaxHeap{}
	fmt.Println(m)

	buildHeap := []int{23, 65, 49, 2, 3, 56, 65, 32, 449, 31, 95, 70, 42, 3, 5, 6, 24}
	for _, v := range buildHeap {
		m.Insert(v)
		fmt.Println(m)
	}

	for i := 0; i < 5; i++ {
		m.Extract()
		fmt.Println("Heap after Extraction #", i+1, ":", m)
	}
}
