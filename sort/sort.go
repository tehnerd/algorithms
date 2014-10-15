package sort

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
		if minIndex != index {
			array[minIndex] = array[index]
			array[index] = minElem
		}
	}
}
