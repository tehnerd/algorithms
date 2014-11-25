package symboltables

const (
	RED   bool = true
	BLACK bool = false
)

type ComparableKV interface {
	/*
		a.Compare(b); returns -1 if a<b
		1 if a>b
		0 if equal
	*/
	Compare(ckv ComparableKV) int
}

type RBTNode struct {
	kv    ComparableKV
	left  *RBTNode
	right *RBTNode
	color bool
}

type RBT struct {
	root      *RBTNode
	elemCount int64
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
	rbt.root = rbt.root.deleteMin(rbt)
}

func (rbtn *RBTNode) deleteMin(rbt *RBT) *RBTNode {
	if rbtn.left == nil {
		rbt.elemCount -= 1
		return nil
	}
	if !rbtn.left.isRed() && !rbtn.left.left.isRed() {
		rbtn = rbtn.moveRedLeft()
	}
	rbtn.left = rbtn.left.deleteMin(rbt)
	return rbtn.balance()
}

func (rbt *RBT) DeleteMax() {
	if !rbt.root.left.isRed() && !rbt.root.right.isRed() {
		rbt.root.color = RED
	}
	rbt.root = rbt.root.deleteMax(rbt)
}

func (rbtn *RBTNode) deleteMax(rbt *RBT) *RBTNode {
	if rbtn.left.isRed() {
		rbtn = rbtn.rotateRight()
	}
	if rbtn.right == nil {
		rbt.elemCount -= 1
		return nil
	}
	if !rbtn.right.isRed() && !rbtn.right.left.isRed() {
		rbtn = rbtn.moveRedRight()
	}
	rbtn.right = rbtn.right.deleteMax(rbt)
	return rbtn.balance()
}

func (rbt *RBT) Delete(ckv ComparableKV) {
	if !rbt.root.left.isRed() && !rbt.root.right.isRed() {
		rbt.root.color = RED
	}
	rbt.root = rbt.root.deleteKey(ckv, rbt)
}

func (rbtn *RBTNode) deleteKey(ckv ComparableKV, rbt *RBT) *RBTNode {
	if ckv.Compare(rbtn.kv) == -1 {
		if !rbtn.left.isRed() && !rbtn.left.left.isRed() {
			rbtn = rbtn.moveRedLeft()
		}
		rbtn.left = rbtn.left.deleteKey(ckv, rbt)
	} else {
		if rbtn.left.isRed() {
			rbtn = rbtn.rotateRight()
		}
		if ckv.Compare(rbtn.kv) == 0 && rbtn.right == nil {
			rbt.elemCount -= 1
			return nil
		}
		if !rbtn.right.isRed() && !rbtn.right.left.isRed() {
			rbtn = rbtn.moveRedRight()
		}
		if ckv.Compare(rbtn.kv) == 0 {
			tmpNode := rbtn.right.findMin()
			rbtn.kv = tmpNode.kv
			rbtn.right = rbtn.deleteMin(rbt)
		} else {
			rbtn.right = rbtn.right.deleteKey(ckv, rbt)
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

func (rbt *RBT) FindMin() ComparableKV {
	min := rbt.root.findMin()
	return min.kv
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

func (rbt *RBT) Len() int64 {
	return rbt.elemCount
}

/*
	ComparableKV could have lots of fields. in this construstion we can pass
	ckv which consists only fields, used by Compare implementation, and if
	we cand find such a node, that ckv.Compare(node)==0, we returns that node
	with all the fields. For example:
	Comparable kv is (key,value) struct, we can pass (key,nil) and if in rbt
	exists node with the same key, we returns (key,value)
*/
func (rbt *RBT) Get(key ComparableKV) ComparableKV {
	return rbt.get(rbt.root, key)
}

func (rbt *RBT) get(root *RBTNode, key ComparableKV) ComparableKV {
	if key.Compare(root.kv) == 0 {
		return root.kv
	} else if key.Compare(root.kv) == -1 && root.left != nil {
		return rbt.get(root.left, key)
	} else if key.Compare(root.kv) == 1 && root.right != nil {
		return rbt.get(root.right, key)
	} else {
		var noneExist ComparableKV
		return noneExist
	}
}

func (rbt *RBT) Put(kv ComparableKV) {
	rbt.root = rbt.put(rbt.root, kv)
	rbt.root.color = BLACK
}

func (rbt *RBT) put(root *RBTNode, kv ComparableKV) *RBTNode {
	if root == nil {
		rbt.elemCount += 1
		return &RBTNode{kv: kv, color: RED}
	} else if kv.Compare(root.kv) == -1 {
		root.left = rbt.put(root.left, kv)
	} else if kv.Compare(root.kv) == 1 {
		root.right = rbt.put(root.right, kv)
	} else {
		root.kv = kv
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
