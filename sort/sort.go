package sort

import (
	"algs/heap"
)

func MergeSort(Array []int32, low, high int) {
	if low < high {
		middle := (low + high) / 2
		MergeSort(Array, low, middle)
		MergeSort(Array, middle+1, high)
		Merge(Array, low, middle, high)
	}
}

func Merge(Array []int32, low, middle, high int) {
	buffer1 := make([]int32, middle-low+1)
	buffer2 := make([]int32, high-middle)
	for cntr := 0; cntr < len(buffer1); cntr++ {
		buffer1[cntr] = Array[low+cntr]
	}
	for cntr := 0; cntr < len(buffer2); cntr++ {
		buffer2[cntr] = Array[middle+1+cntr]
	}
	cntr := low
	for len(buffer1) > 0 && len(buffer2) > 0 {
		if buffer1[0] < buffer2[0] {
			Array[cntr] = buffer1[0]
			buffer1 = buffer1[1:]
		} else {
			Array[cntr] = buffer2[0]
			buffer2 = buffer2[1:]
		}
		cntr++
	}
	if len(buffer1) > 0 {
		for cntr2 := 0; cntr2 < len(buffer1); cntr2++ {
			Array[cntr+cntr2] = buffer1[cntr2]
		}
	}
	if len(buffer2) > 0 {
		for cntr2 := 0; cntr2 < len(buffer2); cntr2++ {
			Array[cntr+cntr2] = buffer2[cntr2]
		}
	}
}

func SelectionSort(array []int32) {
	for index := 0; index < len(array); index++ {
		minElem := array[index]
		minIndex := index
		for elemPointer := index; elemPointer < len(array); elemPointer++ {
			if array[elemPointer] < minElem {
				minElem = array[elemPointer]
				minIndex = elemPointer
			}
		}
		swapElem(array, index, minIndex)
	}
}

func swapElem(array []int32, index1 int, index2 int) {
	tmpVal := array[index1]
	array[index1] = array[index2]
	array[index2] = tmpVal
}

func InsertionSort(array []int32) {
	for index1 := 1; index1 < len(array); index1++ {
		for index2 := index1; index2 > 0; index2-- {
			if array[index2] > array[index2-1] {
				break
			}
			swapElem(array, index2, index2-1)
		}
	}
}

func ShellSort(array []int32) {
	h := 1
	for h < len(array)/3 {
		h = 3*h + 1
	}
	for h >= 1 {
		for index1 := h; index1 < len(array); index1++ {
			for index2 := index1; index2 >= h; index2 -= h {
				if array[index2] > array[index2-h] {
					break
				}
				swapElem(array, index2, index2-h)
			}
		}
		h = h / 3
	}
}

func QuickSort(array []int32, low, high int) {
	if low < high {
		j := partition(array, low, high)
		QuickSort(array, low, j-1)
		QuickSort(array, j+1, high)
	} else {
		return
	}

}

func partition(array []int32, low, high int) int {
	indexLow := low
	indexHigh := high + 1
	partItem := array[low]
	for {
		indexLow += 1
		indexHigh -= 1
		for array[indexLow] < partItem {
			if indexLow == high {
				break
			}
			indexLow += 1
		}
		for partItem < array[indexHigh] {
			if indexHigh == low {
				break
			}
			indexHigh -= 1
		}
		if indexLow >= indexHigh {
			break
		}
		swapElem(array, indexLow, indexHigh)
	}
	swapElem(array, low, indexHigh)
	return indexHigh
}

func FindPercentile(array []int32, percentile int) (int32, int) {
	elemNum := int(float64(len(array)) * (float64(percentile) / 100))
	low := 0
	high := len(array) - 1
	for low < high {
		part := partition(array, low, high)
		if part == elemNum {
			return array[part], elemNum
		} else if part > elemNum {
			high = part - 1
		} else if part < elemNum {
			low = part + 1
		}
	}
	return array[elemNum], elemNum
}
