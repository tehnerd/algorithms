package queues

import (
	"algs/heap"
	"fmt"
)

type PQueueInt32 struct {
	queue []int32
}

func (pq *PQueueInt32) Len() int {
	return len(pq.queue)
}

func (pq *PQueueInt32) Compare(i, j int) int {
	if pq.queue[i] < pq.queue[j] {
		return -1
	} else if pq.queue[i] > pq.queue[j] {
		return 1
	} else {
		return 0
	}
}

func (pq *PQueueInt32) Swap(i, j int) {
	tmp := pq.queue[i]
	pq.queue[i] = pq.queue[j]
	pq.queue[j] = tmp
}

func (pq *PQueueInt32) DelMax() int32 {
	if len(pq.queue) == 0 {
		return -1
	}
	Max := pq.queue[0]
	pq.queue[0] = pq.queue[len(pq.queue)-1]
	pq.queue = pq.queue[:len(pq.queue)-1]
	heap.MaxHeapReheapify(pq, 0)
	return Max
}

func (pq *PQueueInt32) Insert(elem int32) {
	pq.queue = append(pq.queue, elem)
	heap.MaxHeapCheck(pq, len(pq.queue)-1)
}

func (pq *PQueueInt32) PrintElem() {
	for cntr := 0; cntr < len(pq.queue); cntr++ {
		fmt.Print(pq.queue[cntr], " ")
	}
}

func (pq *PQueueInt32) DequeAll() {
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
