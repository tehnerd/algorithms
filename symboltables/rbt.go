package symboltables

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

type RBT struct {
	root *RBTNode
}

func (rbtn *RBTNode) size() int {
	return rbtn.elemCount
}

func (rbtn *RBTNode) isRed() bool {
	if rbtn == nil {
		return false
	} else {
		return rbtn.color == RED
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
	leftNode.right = rbtn
	leftNode.color = rbtn.color
	rbtn.color = RED
	return leftNode
}

func (rbtn *RBTNode) flipColors() {
	if rbtn == nil {
		return
	}
	rbtn.color = RED
	if rbtn.left != nil {
		rbtn.left.color = BLACK
	}
	if rbtn.right != nil {
		rbtn.right.color = BLACK
	}
}

func (rbtn *RBTNode) moveRedLeft() *RBTNode {
	rbtn.flipColors()
	if rbtn.right == nil {
		return rbtn
	}
	if rbtn.right.left.isRed() {
		rbtn.right = rbtn.right.rotateRight()
		rbtn = rbtn.rotateLeft()
	}
	return rbtn
}

func (rbtn *RBTNode) moveRedRight() *RBTNode {
	rbtn.flipColors()
	if rbtn.left == nil {
		return rbtn
	}
	if rbtn.left.left.isRed() {
		rbtn = rbtn.rotateRight()
	}
	return rbtn
}

func (rbt *RBT) DeleteMin() {
	if !rbt.root.left.isRed() && !rbt.root.right.isRed() {
		rbt.root.color = RED
	}
	rbt.root = rbt.root.deleteMin()
}

func (rbtn *RBTNode) deleteMin() *RBTNode {
	if rbtn.left == nil {
		return nil
	}
	if !rbtn.left.isRed() && !rbtn.left.left.isRed() {
		rbtn = rbtn.moveRedLeft()
	}
	rbtn.left = rbtn.left.deleteMin()
	return rbtn.balance()
}

func (rbt *RBT) DeleteMax() {
	if !rbt.root.left.isRed() && !rbt.root.right.isRed() {
		rbt.root.color = RED
	}
	rbt.root = rbt.root.deleteMax()
}

func (rbtn *RBTNode) deleteMax() *RBTNode {
	if rbtn.left.isRed() {
		rbtn = rbtn.rotateRight()
	}
	if rbtn.right == nil {
		return nil
	}
	if !rbtn.right.isRed() && !rbtn.right.left.isRed() {
		rbtn = rbtn.moveRedRight()
	}
	rbtn.right = rbtn.right.deleteMax()
	return rbtn.balance()
}

func (rbt *RBT) Delete(key int32) {
	if !rbt.root.left.isRed() && !rbt.root.right.isRed() {
		rbt.root.color = RED
	}
	rbt.root = rbt.root.deleteKey(key)
}

func (rbtn *RBTNode) deleteKey(key int32) *RBTNode {
	if key < rbtn.key {
		if !rbtn.left.isRed() && !rbtn.left.left.isRed() {
			rbtn = rbtn.moveRedLeft()
		}
		rbtn.left = rbtn.left.deleteKey(key)
	} else {
		if rbtn.left.isRed() {
			rbtn = rbtn.rotateRight()
		}
		if key == rbtn.key && rbtn.right == nil {
			return nil
		}
		if !rbtn.right.isRed() && !rbtn.right.left.isRed() {
			rbtn = rbtn.moveRedRight()
		}
		if key == rbtn.key {
			tmpNode := rbtn.right.findMin()
			rbtn.key = tmpNode.key
			rbtn.value = tmpNode.value
			rbtn.right = rbtn.deleteMin()
		} else {
			rbtn.right = rbtn.right.deleteKey(key)
		}
	}
	return rbtn.balance()
}

func (rbtn *RBTNode) findMin() *RBTNode {
	if rbtn.left != nil {
		return rbtn.left.findMin()
	}
	return rbtn
}

func (rbt *RBT) FindMin() int32 {
	min := rbt.root.findMin()
	return min.value
}

func (rbtn *RBTNode) balance() *RBTNode {
	if rbtn.right.isRed() {
		rbtn = rbtn.rotateLeft()
	}
	if rbtn.right.isRed() && !rbtn.left.isRed() {
		rbtn = rbtn.rotateLeft()
	}
	if rbtn.left.isRed() && rbtn.left.left.isRed() {
		rbtn = rbtn.rotateRight()
	}
	if rbtn.right.isRed() && rbtn.left.isRed() {
		rbtn.flipColors()
	}
	return rbtn

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
		return rbt.get(root.left, key)
	} else if key > root.key && root.right != nil {
		return rbt.get(root.right, key)
	} else {
		return -1
	}
}

func (rbt *RBT) Put(key, value int32) {
	rbt.root = rbt.put(rbt.root, key, value)
	rbt.root.color = BLACK
}

func (rbt *RBT) put(root *RBTNode, key int32, value int32) *RBTNode {
	if root == nil {
		return &RBTNode{key: key, value: value, color: RED}
	} else if key < root.key {
		root.left = rbt.put(root.left, key, value)
	} else if key > root.key {
		root.right = rbt.put(root.right, key, value)
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
