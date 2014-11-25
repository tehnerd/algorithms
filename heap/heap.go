package heap

type Comparable interface {
	Len() int
	/*
		Compare must returns -1 if Comparable[i]<Comparable[j]
		1 if [i] > [j]
		and 0 if equal
	*/
	Compare(i, j int) int
	Swap(i, j int)
}

func HeapParent(index int) int {
	if index == 0 {
		return -1
	} else {
		return (index - 1) / 2
	}
}

func BuildMaxHeap(heap Comparable) {
	for cntr := 0; cntr < heap.Len(); cntr++ {
		MaxHeapCheck(heap, cntr)
	}
}

func BuildMinHeap(heap Comparable) {
	for cntr := 0; cntr < heap.Len(); cntr++ {
		MinHeapCheck(heap, cntr)
	}
}

func MaxHeapCheck(MaxHeap Comparable, pos int) {
	if HeapParent(pos) == -1 {
		return
	} else if MaxHeap.Compare(HeapParent(pos), pos) == -1 {
		MaxHeap.Swap(pos, HeapParent(pos))
		MaxHeapCheck(MaxHeap, HeapParent(pos))
	}
}

func MaxHeapReheapify(MaxHeap Comparable, pos int) {
	if 2*(pos+1)-1 >= MaxHeap.Len() {
		return
	}
	if MaxHeap.Compare(pos, 2*(pos+1)-1) == -1 {
		MaxHeap.Swap(pos, 2*(pos+1)-1)
		MaxHeapReheapify(MaxHeap, 2*(pos+1)-1)
	}
	if 2*(pos+1) >= MaxHeap.Len() {
		return
	}
	if MaxHeap.Compare(pos, 2*(pos+1)) == -1 {
		MaxHeap.Swap(pos, 2*(pos+1))
		MaxHeapReheapify(MaxHeap, 2*(pos+1))
	}
	return
}

func MinHeapReheapify(MinHeap Comparable, pos int) {
	if 2*(pos+1)-1 >= MinHeap.Len() {
		return
	}
	if MinHeap.Compare(pos, 2*(pos+1)-1) == 1 {
		MinHeap.Swap(pos, 2*(pos+1)-1)
		MinHeapReheapify(MinHeap, 2*(pos+1)-1)
	}
	if 2*(pos+1) >= MinHeap.Len() {
		return
	}
	if MinHeap.Compare(pos, 2*(pos+1)) == 1 {
		MinHeap.Swap(pos, 2*(pos+1))
		MinHeapReheapify(MinHeap, 2*(pos+1))
	}
	return
}

func MinHeapCheck(MinHeap Comparable, pos int) {
	parentPos := HeapParent(pos)
	if parentPos == -1 {
		return
	} else if MinHeap.Compare(parentPos, pos) == 1 {
		MinHeap.Swap(pos, parentPos)
		MinHeapCheck(MinHeap, parentPos)
	}
}
