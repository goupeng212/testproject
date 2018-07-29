package main

import (
	"fmt"
	//"time"
)

type BinaryHeap struct {
	array []int
}

func NewBinaryHeap(arr []int) *BinaryHeap {
	var a []int
	biheap := &BinaryHeap{array: a}
	for _, item := range arr {
		biheap.Insert(item)
	}
	return biheap
}

func (heap *BinaryHeap) Print() {
	fmt.Println("Array Print:")
	for _, item := range heap.array {
		fmt.Print(item, ",")

	}
	fmt.Println("")
}
func (heap *BinaryHeap) Size() int {
	return len(heap.array)
}

func (heap *BinaryHeap) Insert(data int) {
	length := heap.Size()
	heap.array = append(heap.array, data)
	heap.filterup(length)
	//	heap.Print()
}
func (heap *BinaryHeap) filterup(start int) {
	current := start
	parent := (current - 1) / 2
	temp := heap.array[current]
	for current > 0 {
		//		fmt.Println("current:", current, " value:", heap.array[current], " parent:", parent, "  value:", heap.array[parent])

		if heap.array[current] > heap.array[parent] {
			//			fmt.Println("current:", current, " value:", heap.array[current], " parent:", parent, "  value:", heap.array[parent])
			heap.array[current] = heap.array[parent]
			heap.array[parent] = temp
			//			heap.Print()
			current = parent
			parent = (parent - 1) / 2
			//			fmt.Println("current:", current, " parent:", parent)
		} else {
			//			fmt.Println("enter else")
			break
		}
	}
	heap.array[current] = temp
}

func (heap *BinaryHeap) filterdown(start int, end int) {
	current := start
	leftChild := current*2 + 1
	temp := heap.array[current]
	//	fmt.Println("leftchild:", leftChild, "end:", end)
	for leftChild < end {
		fmt.Println("leftchild:", leftChild, "end:", end)

		if (leftChild+1 < end) && (heap.array[leftChild] < heap.array[leftChild+1]) {
			leftChild++
		}
		if heap.array[current] < heap.array[leftChild] {
			heap.array[current] = heap.array[leftChild]
			//			heap.array[leftChild] = temp
			current = leftChild
			leftChild = leftChild*2 + 1
			//			temp = heap.array[current]

			fmt.Println("leftchild:", leftChild, "end:", end)

			heap.Print()
		} else {
			break
		}
	}
	heap.array[current] = temp
	//	heap.Print()

}
func (heap *BinaryHeap) HeapSort() {
	for i := heap.Size() - 1; i > 0; i-- {
		heap.swap(0, i)
		heap.Print()
		heap.filterdown(0, i)

	}
}
func (heap *BinaryHeap) swap(i, j int) {
	tmp := heap.array[i]
	heap.array[i] = heap.array[j]
	heap.array[j] = tmp
}

//func main() {

//	arr := []int{82, 3, 21, 5, 2, 4, 6, 87, 22, 45, 78, 90}
//	biheap := NewBinaryHeap(arr)
//	biheap.Print()
//	biheap.HeapSort()
//	biheap.Print()
//}
