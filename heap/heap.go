package heap

import (
	"fmt"
)

func PrintInt32Heap(Heap []int32) {
	for cntr := 0; cntr < len(Heap)/2; cntr++ {
		fmt.Println(Heap[cntr])
		if (2*cntr + 2) < len(Heap) {
			fmt.Println(fmt.Sprintf("%d %d", Heap[2*cntr+1], Heap[2*cntr+2]))
		} else {
			fmt.Println(fmt.Sprintf("%d", Heap[2*cntr+1]))
		}
	}
}

func HeapInt32Parent(index int) int {
	if index == 0 {
		return -1
	} else {
		return (index - 1) / 2
	}
}

func BuildMaxHeapInt32(Heap []int32) []int32 {
	MaxHeap := make([]int32, len(Heap))
	for cntr := 0; cntr < len(Heap); cntr++ {
		MaxHeapInsert(MaxHeap, Heap[cntr], cntr)
	}
	return MaxHeap
}

func BuildMinHeapInt32(Heap []int32) []int32 {
	MinHeap := make([]int32, len(Heap))
	for cntr := 0; cntr < len(Heap); cntr++ {
		MinHeapInsert(MinHeap, Heap[cntr], cntr)
	}
	return MinHeap
}

func swapElem(array []int32, pos1, pos2 int) {
	tmp := array[pos1]
	array[pos1] = array[pos2]
	array[pos2] = tmp
}

func MaxHeapInsert(MaxHeap []int32, newElem int32, pos int) {
	MaxHeap[pos] = newElem
	if HeapInt32Parent(pos) == -1 {
		return
	} else if MaxHeap[HeapInt32Parent(pos)] < newElem {
		swapElem(MaxHeap, pos, HeapInt32Parent(pos))
		MaxHeapInsert(MaxHeap, newElem, HeapInt32Parent(pos))
	}
}

func MaxHeapReheapify(MaxHeap []int32, pos int) {
	if 2*(pos+1)-1 >= len(MaxHeap) {
		return
	}
	if MaxHeap[pos] < MaxHeap[2*(pos+1)-1] {
		swapElem(MaxHeap, pos, 2*(pos+1)-1)
		MaxHeapReheapify(MaxHeap, 2*(pos+1)-1)
	}
	if 2*(pos+1) >= len(MaxHeap) {
		return
	}
	if MaxHeap[pos] < MaxHeap[2*(pos+1)] {
		swapElem(MaxHeap, pos, 2*(pos+1))
		MaxHeapReheapify(MaxHeap, 2*(pos+1))
	}
	return
}

func MinHeapReheapify(MaxHeap []int32, pos int) {
	if 2*(pos+1)-1 >= len(MaxHeap) {
		return
	}
	if MaxHeap[pos] > MaxHeap[2*(pos+1)-1] {
		swapElem(MaxHeap, pos, 2*(pos+1)-1)
		MaxHeapReheapify(MaxHeap, 2*(pos+1)-1)
	}
	if 2*(pos+1) >= len(MaxHeap) {
		return
	}
	if MaxHeap[pos] > MaxHeap[2*(pos+1)] {
		swapElem(MaxHeap, pos, 2*(pos+1))
		MaxHeapReheapify(MaxHeap, 2*(pos+1))
	}
	return
}

func MinHeapInsert(MinHeap []int32, newElem int32, pos int) {
	MinHeap[pos] = newElem
	parentPos := HeapInt32Parent(pos)
	if parentPos == -1 {
		return
	} else if MinHeap[parentPos] > newElem {
		tmp := newElem
		MinHeap[pos] = MinHeap[parentPos]
		MinHeap[parentPos] = tmp
		MinHeapInsert(MinHeap, newElem, parentPos)
	}
}
