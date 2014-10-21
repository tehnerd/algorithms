package main

import (
	"algs/queues"
	"algs/sort"
	"fmt"
	"math/rand"
	"time"
)

func timeDecor(intArrayFunc func([]int32), array []int32) {
	t1 := time.Now()
	intArrayFunc(array)
	t2 := time.Now()
	fmt.Println(t2.UnixNano() - t1.UnixNano())

}

func timeDecor2(intArrayFunc func([]int32, int, int), array []int32, start, end int) {
	t1 := time.Now()
	intArrayFunc(array, start, end)
	t2 := time.Now()
	fmt.Println(t2.UnixNano() - t1.UnixNano())

}

func isSorted(array []int32) bool {
	for cntr := 0; cntr < len(array)-1; cntr++ {
		if array[cntr+1] < array[cntr] {
			return false
		}
	}
	return true
}

func testArray(array []int32) {
	if isSorted(array) {
		fmt.Println("sorted")
	} else {
		fmt.Println("not sorted")
	}
}

func copyArray(array1 []int32) []int32 {
	array2 := make([]int32, len(array1))
	for cntr := 0; cntr < len(array1); cntr++ {
		array2[cntr] = array1[cntr]
	}
	return array2
}

func drawSort(index1, index2, arrayLength int) {
	for cntr := 0; cntr < arrayLength; cntr++ {
		if cntr == index1 || cntr == index2 {
			fmt.Print('|')
		} else {
			fmt.Print('.')
		}
	}
	fmt.Println("")
}

func main() {
	fmt.Println("hello")
	initarray := make([]int32, 10000)
	rand.Seed(time.Now().UnixNano())
	for cntr := 0; cntr < cap(initarray); cntr++ {
		initarray[cntr] = rand.Int31n(20000000)
	}
	testArray(initarray)
	array1 := copyArray(initarray)
	//maxHeap := heap.BuildMaxHeapInt32(array1)
	//minHeap := heap.BuildMinHeapInt32(array1)
	fmt.Println("#####  merge sort  #####")
	timeDecor2(sort.MergeSort, array1, 0, len(array1)-1)
	testArray(array1)
	timeDecor2(sort.MergeSort, array1, 0, len(array1)-1)
	array2 := copyArray(initarray)
	fmt.Println("#####  selection sort  #####")
	timeDecor(sort.SelectionSort, array2)
	testArray(array2)
	timeDecor(sort.SelectionSort, array2)
	fmt.Println("#####  insertion sort  #####")
	array3 := copyArray(initarray)
	timeDecor(sort.InsertionSort, array3)
	testArray(array3)
	timeDecor(sort.InsertionSort, array3)
	fmt.Println("#####  shell sort  #####")
	array4 := copyArray(initarray)
	timeDecor(sort.ShellSort, array4)
	testArray(array4)
	timeDecor(sort.ShellSort, array4)
	fmt.Println("#####  quick sort  #####")
	array5 := copyArray(initarray)
	timeDecor2(sort.QuickSort, array5, 0, len(array5)-1)
	testArray(array5)
	timeDecor2(sort.QuickSort, array5, 0, len(array5)-1)
	fmt.Println("##### heap sort #####")
	array6 := copyArray(initarray)
	timeDecor(sort.HeapSort, array6)
	testArray(array6)
	timeDecor(sort.HeapSort, array6)
	fmt.Println("##### priority queue #####")
	var pq queues.PQueue
	for cntr := 0; cntr < len(initarray); cntr++ {
		(&pq).Insert(initarray[cntr])
	}
	(&pq).DequeAll()
}
