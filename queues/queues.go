package queues

import (
	"algs/heap"
	"fmt"
)

type PQueue struct {
	queue []int32
}

func (pq *PQueue) DelMax() int32 {
	if len(pq.queue) == 0 {
		return -1
	}
	Max := pq.queue[0]
	pq.queue[0] = pq.queue[len(pq.queue)-1]
	pq.queue = pq.queue[:len(pq.queue)-1]
	heap.MaxHeapReheapify(pq.queue, 0)
	return Max
}

func (pq *PQueue) Insert(elem int32) {
	pq.queue = append(pq.queue, elem)
	heap.MaxHeapInsert(pq.queue, elem, len(pq.queue)-1)
}

func (pq *PQueue) PrintElem() {
	for cntr := 0; cntr < len(pq.queue); cntr++ {
		fmt.Print(pq.queue[cntr], " ")
	}
}

func (pq *PQueue) DequeAll() {
	elem := int32(0)
	privElem := int32(-1)
	for len(pq.queue) != 0 {
		elem = pq.DelMax()
		if privElem != -1 {
			if privElem < elem {
				fmt.Println("ERROR IN PQ")
			}
			privElem = elem
		}
	}
	fmt.Println("pq complited; check errors (if any) above")
}
