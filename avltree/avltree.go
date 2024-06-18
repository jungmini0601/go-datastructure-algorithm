package avltree

type AvlTree struct {
	root *Node
}

func (t *AvlTree) Insert(value *Node) {
	t.root = t.root.insert(value)
}

func (t *AvlTree) Search(value int) *Node {
	return t.root.search(value)
}

type Node struct {
	value  int
	left   *Node
	right  *Node
	height int
}

func NewAVLTree(value int) *Node {
	return &Node{
		value:  value,
		left:   nil,
		right:  nil,
		height: 0,
	}
}

func (n *Node) updateHeight() {
	left, right := -1, -1
	if n.left != nil {
		left = n.left.height
	}

	if n.right != nil {
		right = n.right.height
	}

	if left > right {
		n.height = left + 1
	} else {
		n.height = right + 1
	}
}

func (n *Node) leftRotation() *Node {
	rightChild := n.right
	n.right = rightChild.left
	rightChild.left = n
	n.updateHeight()
	rightChild.updateHeight()
	return rightChild
}

func (n *Node) rightRotation() *Node {
	leftChild := n.left
	n.left = leftChild.right
	leftChild.right = n
	n.updateHeight()
	leftChild.updateHeight()
	return leftChild
}

func (n *Node) balanceFactor() int {
	left, right := -1, -1
	if n.left != nil {
		left = n.left.height
	}

	if n.right != nil {
		right = n.right.height
	}

	return left - right
}

func (n *Node) rebalance() *Node {
	balanceFactor := n.balanceFactor()
	// balanceFactor가 양수이면 왼쪽의 높이가 더 큰 것
	// balanceFactor의 절대 값이 2보다 크다면 균형이 무너진 것
	if balanceFactor > 1 {
		// 왼쪽 자식노드의 BF 값이 음수이면 우측이 더 큰 것이므로 자식 노드를 회전시켜줌
		if n.left != nil && n.left.balanceFactor() < 0 {
			n.left = n.left.leftRotation()
		}
		// LL 회전
		n = n.rightRotation()
	} else if balanceFactor < -1 {
		if n.right != nil && n.right.balanceFactor() > 0 {
			n.right = n.right.rightRotation()
		}
		// RR 회전
		n = n.leftRotation()
	}
	return n
}

func (n *Node) insert(newNode *Node) *Node {
	if n == nil {
		return newNode
	}

	if newNode.value < n.value {
		n.left = n.left.insert(newNode)
	} else if newNode.value > n.value {
		n.right = n.right.insert(newNode)
	} else {
		n.value = newNode.value
	}
	n.updateHeight()
	return n.rebalance()
}

func (n *Node) search(value int) *Node {
	if n == nil {
		return nil
	}

	for current := n; current != nil; {
		if value < current.value {
			current = current.left
		} else if value > current.value {
			current = current.right
		} else {
			return current
		}
	}

	return nil
}

// 삭제 될 경우 반환 값은 Nil
func (n *Node) delete(value int) *Node {
	if n == nil {
		return nil
	}
	if value < n.value {
		n = n.left.delete(value)
	} else if value > n.value {
		n = n.right.delete(value)
	} else {
		if n.right != nil {
			successor := n.right
			// 삭제할 노드의 자식 노드중 가장 작은 값을 찾는다.
			for successor.left != nil {
				successor = successor.left
			}
			n.value = successor.value
			n.right = n.right.delete(successor.value)
		} else if n.left != nil {
			n = n.left
		} else {
			n = nil
			return n
		}
	}

	if n != nil {
		n.updateHeight()
		n = n.rebalance()
	}
	return n
}
