package visual

import (
	"fmt"
)

func DrawSort(index1, index2, arrayLength int) {
	for cntr := 0; cntr < arrayLength; cntr++ {
		if cntr == index1 || cntr == index2 {
			fmt.Print("|")
		} else {
			fmt.Print(".")
		}
	}
	fmt.Println("")
}
