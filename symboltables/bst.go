package symboltables

type BSTNode struct {
	key       int32
	value     int32
	left      *BSTNode
	right     *BSTNode
	elemCount int
}

func (bstn *BSTNode) size() int {
	return bstn.elemCount
}

type BST struct {
	root *BSTNode
}

func (bst *BST) Size() int {
	return bst.root.size()
}

func (bst *BST) Get(key int32) int32 {
	return bst.get(bst.root, key)
}

func (bst *BST) get(root *BSTNode, key int32) int32 {
	if key == root.key {
		return root.value
	} else if key < root.key && root.left != nil {
		return bst.get(root.left, key)
	} else if key > root.key && root.right != nil {
		return bst.get(root.right, key)
	} else {
		return -1
	}
}

func (bst *BST) Put(key, value int32) {
	bst.put(bst.root, key, value)
}

func (bst *BST) put(root *BSTNode, key int32, value int32) {
	if root == nil {
		bst.root = &BSTNode{key: key, value: value}
	} else if key < root.key {
		if root.left != nil {
			bst.put(root.left, key, value)
		} else {
			root.left = &BSTNode{key: key, value: value}
		}
	} else if key > root.key {
		if root.right != nil {
			bst.put(root.right, key, value)
		} else {
			root.right = &BSTNode{key: key, value: value}
		}
	} else {
		root.value = value
	}

}
