package bst

type binarySearchTree struct {
	value int
	left  *binarySearchTree
	right *binarySearchTree
}

func (tree *binarySearchTree) Insert(value int) *binarySearchTree {
	BST := binarySearchTree{value: value}
	for {
		if value < tree.value {
			if tree.left == nil {
				tree.left = &BST
				break
			} else {
				tree = tree.left
			}
		} else {
			if tree.right == nil {
				tree.right = &BST
				break
			} else {
				tree = tree.right
			}
		}
	}
	return tree
}

func (tree *binarySearchTree) Contains(value int) bool {
	temp := tree
	for temp != nil {
		if value < temp.value {
			temp = temp.left
		} else if value > temp.value {
			temp = temp.right
		} else {
			return true
		}
	}
	return false
}

func (tree *binarySearchTree) Remove(value int) *binarySearchTree {
	tree.RemoveHelper(value, nil)
	return tree
}

func (tree *binarySearchTree) RemoveHelper(value int, parent *binarySearchTree) {
	current := tree
	for current != nil {
		if value < current.value {
			parent = current
			current = current.left
		} else if value > current.value {
			parent = current
			current = current.right
		} else {
			if current.left != nil && current.right != nil {
				current.value = current.right.getMinValue()
				current.right.RemoveHelper(current.value, current)
			} else if parent == nil {
				if current.left != nil {
					current.value = current.left.value
					current.right = current.left.right
					current.left = current.left.left
				} else if current.right != nil {
					current.value = current.right.value
					current.left = current.right.left
					current.right = current.right.right
				} else {
					current.value = 0
				}
			} else if parent.left == current {
				if current.left != nil {
					parent.left = current.left
				} else {
					parent.left = current.right
				}
			} else if parent.right == current {
				if current.left != nil {
					parent.right = current.left
				} else {
					parent.right = current.right
				}
			}
			break
		}
	}
}

func (tree *binarySearchTree) getMinValue() int {
	if tree.left == nil {
		return tree.value
	}
	return tree.left.getMinValue()
}
