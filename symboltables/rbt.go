package symboltables

import "fmt"

const (
	RED   bool = true
	BLACK bool = false
)

type RBTNode struct {
	key       int32
	value     int32
	left      *RBTNode
	right     *RBTNode
	color     bool
	elemCount int
}

func (rbtn *RBTNode) size() int {
	return rbtn.elemCount
}

func isRed2(rbtn *RBTNode) bool {
	if rbtn == nil {
		return false
	} else {
		return (*rbtn).color == RED
	}
}

func (rbtn *RBTNode) isRed() bool {
	if rbtn == nil {
		return false
	} else {
		return (*rbtn).color == RED
	}
}

func (rbtn *RBTNode) rotateLeft() *RBTNode {
	rightNode := rbtn.right
	rbtn.right = rightNode.left
	rightNode.left = rbtn
	rightNode.color = rbtn.color
	rbtn.color = RED
	return rightNode
}

func (rbtn *RBTNode) rotateRight() *RBTNode {
	leftNode := rbtn.left
	rbtn.left = leftNode.right
	leftNode.left = rbtn
	leftNode.color = rbtn.color
	rbtn.color = RED
	return leftNode
}

func (rbtn *RBTNode) flipColors() {
	rbtn.color = RED
	rbtn.left.color = BLACK
	rbtn.right.color = BLACK
}

type RBT struct {
	root *RBTNode
}

func (rbt *RBT) Size() int {
	return rbt.root.size()
}

func (rbt *RBT) Get(key int32) int32 {
	return rbt.get(rbt.root, key)
}

func (rbt *RBT) get(root *RBTNode, key int32) int32 {
	if key == root.key {
		return root.value
	} else if key < root.key && root.left != nil {
		fmt.Println("further")
		return rbt.get(root.left, key)
	} else if key > root.key && root.right != nil {
		fmt.Println("further")
		return rbt.get(root.right, key)
	} else {
		return -1
	}
}

func (rbt *RBT) Put(key, value int32) {
	rbt.put(rbt.root, key, value)
	rbt.root.color = BLACK
}

func (rbt *RBT) put(root *RBTNode, key int32, value int32) *RBTNode {
	if root == nil {
		rbt.root = &RBTNode{key: key, value: value, color: RED}
		root = rbt.root
	} else if key < root.key {
		if root.left != nil {
			root.left = rbt.put(root.left, key, value)
		} else {
			root.left = &RBTNode{key: key, value: value, color: RED}
		}
	} else if key > root.key {
		if root.right != nil {
			root.right = rbt.put(root.right, key, value)
		} else {
			root.right = &RBTNode{key: key, value: value, color: RED}
		}
	} else {
		root.value = value
	}
	if root.right.isRed() && !root.left.isRed() {
		root = root.rotateLeft()
	}
	if root.left.isRed() && root.left.left.isRed() {
		root = root.rotateRight()
	}
	if root.right.isRed() && root.left.isRed() {
		root.flipColors()
	}
	return root
}
