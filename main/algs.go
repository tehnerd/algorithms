package main

import (
	"algs/heap"
	"algs/queues"
	//"algs/sort"
	"algs/symboltables"
	"fmt"
	"math/rand"
	//build in sort
	//bsort "sort"
	"algs/graphs"
	"bufio"
	"os"
	"strconv"
	"strings"
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

func timeDecor3(intArrayFunc func([]int32, int) (int32, int),
	array []int32, intVal int) (int32, int) {
	t1 := time.Now()
	rval1, rval2 := intArrayFunc(array, intVal)
	t2 := time.Now()
	fmt.Println(t2.UnixNano() - t1.UnixNano())
	return rval1, rval2
}

func timeDecor4(testFunc func(int32) int32, tval int32) {
	t1 := time.Now()
	rval := testFunc(tval)
	t2 := time.Now()
	fmt.Println(rval)
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

func sameElems(array []int32) int {
	ammountOfSame := 0
	for cntr := 0; cntr < len(array)-1; cntr++ {
		if array[cntr] == array[cntr+1] {
			ammountOfSame++
		}
	}
	return ammountOfSame
}

func testArray(array []int32) {
	if isSorted(array) {
		fmt.Println("sorted")
	} else {
		fmt.Println("not sorted")
	}
	fmt.Println(sameElems(array))
}

func copyArray(array1 []int32) []int32 {
	array2 := make([]int32, len(array1))
	for cntr := 0; cntr < len(array1); cntr++ {
		array2[cntr] = array1[cntr]
	}
	return array2
}

func ReadGraph() *graphs.Graph {
	graph := new(graphs.Graph)
	fd, err := os.Open(os.Args[1])
	defer fd.Close()
	if err != nil {
		os.Exit(1)
	}
	reader := bufio.NewReader(fd)
	line, err := reader.ReadString('\n')
	for err == nil {
		fields := strings.Fields(line)
		if len(fields) > 0 {
			vertice := fields[0]
			vertice = strings.Split(vertice, ":")[0]
			for cntr := 1; cntr < len(fields); cntr += 2 {
				distance, _ := strconv.Atoi(fields[cntr+1])
				graph.AddEdgeUnidirectDistance(vertice, fields[cntr],
					int32(distance))
			}
		}
		line, err = reader.ReadString('\n')
	}
	return graph
}

type TestComparable []int32

func (tc TestComparable) Compare(i, j int) int {
	if tc[i] < tc[j] {
		return -1
	} else if tc[i] > tc[j] {
		return 1
	} else {
		return 0
	}
}

func (tc TestComparable) Len() int {
	return len(tc)
}

func (tc TestComparable) Swap(i, j int) {
	tmp := tc[i]
	tc[i] = tc[j]
	tc[j] = tmp
}

func main() {
	fmt.Println("hello")
	initarray := make([]int32, 100)
	rand.Seed(time.Now().UnixNano())
	for cntr := 0; cntr < cap(initarray); cntr++ {
		initarray[cntr] = rand.Int31n(int32(len(initarray)))
	}
	testArray(initarray)
	//array0 := copyArray(initarray)
	//perc, num := timeDecor3(sort.FindPercentile, array0, 50)
	//fmt.Println(perc, " ", num)

	/*
		array1 := copyArray(initarray)
		//maxHeap := heap.BuildMaxHeapInt32(array1)
		//minHeap := heap.BuildMinHeapInt32(array1)
		fmt.Println("##### percentile ######")
		array0 := copyArray(initarray)
		perc, num := timeDecor3(sort.FindPercentile, array0, 50)
		fmt.Println(perc, " ", num)
		fmt.Println("#####  merge sort  #####")
		timeDecor2(sort.MergeSort, array1, 0, len(array1)-1)
		testArray(array1)
		timeDecor2(sort.MergeSort, array1, 0, len(array1)-1)
		fmt.Println(array1[num])
		fmt.Println("#####  selection sort  #####")
		array2 := copyArray(initarray)
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
	*/
	array6 := make(TestComparable, 0)
	for cntr := 0; cntr < len(initarray); cntr++ {
		array6 = append(array6, initarray[cntr])
	}
	fmt.Println(array6)
	heap.BuildMinHeap(array6)
	fmt.Println(array6)

	fmt.Println("##### priority queue #####")
	var pq queues.PQueueInt32
	for cntr := 0; cntr < len(initarray); cntr++ {
		(&pq).Insert(initarray[cntr])
	}
	(&pq).DequeAll()
	/*
			fmt.Println("##### linked lists #####")
			var llist symboltables.LList
			for cntr := 0; cntr < len(initarray); cntr++ {
				(&llist).Add(initarray[cntr], int32(cntr))
			}
			timeDecor4((&llist).Get, tkey)
			fmt.Println(tkey)

		fmt.Println("##### binary search ####")
		var bsstruct symboltables.BinarySearchST
		t1 := time.Now()
		for cntr := 0; cntr < len(initarray); cntr++ {
			(&bsstruct).Append(initarray[cntr], int32(cntr))
		}
		t2 := time.Now()
		fmt.Println(t2.UnixNano() - t1.UnixNano())
		//bsort.Sort(&bsstruct)
		timeDecor4((&bsstruct).Search, tkey)
	*/
	fmt.Println("##### Binary Search Tree #####")
	tkey := rand.Int31n(int32(len(initarray)))
	var bst symboltables.BST
	fmt.Println("create time")
	t1 := time.Now()
	for cntr := 0; cntr < len(initarray); cntr++ {
		(&bst).Put(initarray[cntr], int32(cntr))
	}
	t2 := time.Now()
	fmt.Println(t2.UnixNano() - t1.UnixNano())
	timeDecor4((&bst).Get, tkey)
	timeDecor4((&bst).Get, tkey)
	//	fmt.Println((&bst).Get(tkey))
	fmt.Println("##### Red Black Tree #####")
	var rbt symboltables.RBT
	fmt.Println("create time")
	t1 = time.Now()
	for cntr := 0; cntr < len(initarray); cntr++ {
		(&rbt).Put(initarray[cntr], int32(cntr))
	}
	t2 = time.Now()
	fmt.Println(t2.UnixNano() - t1.UnixNano())
	timeDecor4((&rbt).Get, tkey)
	timeDecor4((&rbt).Get, tkey)
	fmt.Println("rbt delete min,max")
	(&rbt).Put(int32(-100), int32(2307))
	fmt.Println((&rbt).Get(int32(-100)))
	fmt.Println((&rbt).FindMin())
	(&rbt).DeleteMin()
	fmt.Println((&rbt).FindMin())
	fmt.Println((&rbt).Get(int32(-100)))
	(&rbt).Put(int32(1<<31-1), int32(51574168))
	fmt.Println((&rbt).Get(int32(1<<31 - 1)))
	(&rbt).DeleteMax()
	fmt.Println((&rbt).Get(int32(1<<31 - 1)))

	(&rbt).Put(int32(-100), int32(2307))
	fmt.Println((&rbt).FindMin())
	(&rbt).Delete(int32(-100))
	fmt.Println((&rbt).FindMin())
	//	fmt.Println((&bst).Get(tkey))
	fmt.Println("GRAPHS")
	var g1 graphs.Graph
	(&g1).AddEdge("a", "b")
	(&g1).AddEdge("a", "c")
	(&g1).AddEdge("a", "g")
	(&g1).AddEdge("a", "z")
	(&g1).AddEdge("g", "z")
	(&g1).Print()
	(&g1).RemoveEdge("a", "g")
	(&g1).Print()
	fmt.Println((&g1).Vertices())
	fmt.Println(graphs.IsConnected(&g1, "a"))
	fmt.Println(graphs.PathTo(&g1, "a", "g"))
	bfp := new(graphs.BreadthFirstPath)
	fmt.Println(bfp.BFPHasPathTo(&g1, "a", "g"))
	if len(os.Args) > 1 {
		g := ReadGraph()
		if len(os.Args) >= 3 {
			var spf = new(graphs.SPF)
			spf.Init(g)
			spf.SP(os.Args[2])
			fmt.Println(spf.SPFDist())
		}
	}
}
