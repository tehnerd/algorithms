package main

import (
	"algs/heap"
	"algs/sort"
	"fmt"
	"math/rand"
	"time"
)

func copyArray(array1 []int32) []int32 {
	array2 := make([]int32, len(array1))
	for cntr := 0; cntr < len(array1); cntr++ {
		array2[cntr] = array1[cntr]
	}
	return array2
}

func main() {
	fmt.Println("hello")
	initarray := make([]int32, 10)
	rand.Seed(time.Now().UnixNano())
	for cntr := 0; cntr < cap(initarray); cntr++ {
		initarray[cntr] = rand.Int31n(1000)
	}
	array1 := copyArray(initarray)
	fmt.Println(initarray)
	maxHeap := heap.BuildMaxHeapInt32(array1)
	minHeap := heap.BuildMinHeapInt32(array1)
	fmt.Println(maxHeap)
	fmt.Println(minHeap)
	fmt.Println("merge sort")
	sort.MergeSort(array1, 0, len(array1)-1)
	fmt.Println(array1)
	array2 := copyArray(initarray)
	fmt.Println("selection sort")
	sort.SelectionSort(array2)
	fmt.Println(array2)
	fmt.Println("insertion sort")
	array3 := copyArray(initarray)
	sort.InsertionSort(array3)
	fmt.Println(array3)
}
