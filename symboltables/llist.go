package symboltables

import (
	"fmt"
)

type LList struct {
	key   int32
	value int32
	next  *LList
}

func (llist *LList) Add(key int32, value int32) {
	for llist.next != nil {
		if llist.key != key {
			llist = llist.next
		} else {
			break
		}
	}
	if llist.key != key {
		llist.next = new(LList)
		llist = llist.next
		llist.key = key
		llist.value = value
	} else {
		llist.value = value
	}
}

func (llist *LList) Print() {
	for llist.next != nil {
		fmt.Println("key:", llist.key, " value:", llist.value)
		llist = llist.next
	}
	fmt.Println("key:", llist.key, " value:", llist.value)
}

func (llist *LList) Get(key int32) int32 {
	for llist != nil {
		if llist.key == key {
			return llist.value
		}
		llist = llist.next
	}
	return -1
}

func (llist *LList) Search(value int32) int32 {
	for llist != nil {
		if llist.value == value {
			return llist.value
		}
		llist = llist.next
	}
	return -1
}
