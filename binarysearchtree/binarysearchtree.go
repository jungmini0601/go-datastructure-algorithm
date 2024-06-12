package binarysearchtree

// 이진 탐색트리 노드
type Node struct {
	value int
	left  *Node
	right *Node
}

type BST struct {
	root *Node
}

func NewBST() *BST {
	return &BST{root: nil}
}

func (bst *BST) Search(value int) *Node {
	return search(bst.root, value)
}

func search(n *Node, value int) *Node {
	if n == nil || n.value == value {
		return n
	}

	if value < n.value {
		return search(n.left, value)
	}

	return search(n.right, value)
}

func (bst *BST) Insert(value int) {
	n := &Node{value: value, left: nil, right: nil}
	if bst.root == nil {
		bst.root = n
		return
	} else {
		insert(bst.root, n)
	}
}

func insert(n *Node, newNode *Node) {
	// 새로운 노드가 왼쪽 자식 노드에 들어가는 경우
	if newNode.value < n.value {
		if n.left == nil {
			n.left = newNode
		} else {
			insert(n.left, newNode)
		}
	} else { // 새로운 노드가 오른쪽 자식 노드에 들어가는 경우
		if n.right == nil {
			n.right = newNode
		} else {
			insert(n.right, newNode)
		}
	}
}

func (bst *BST) Delete(value int) {
	bst.root = delete(bst.root, value)
}

func delete(n *Node, value int) *Node {
	if n == nil {
		return nil
	}

	if value < n.value {
		n.left = delete(n.left, value)
		return n
	} else if value > n.value {
		n.right = delete(n.right, value)
		return n
	} else { // 지워져야 하는 노드
		if n.left == nil && n.right == nil { // 리프 노드이면 참조를 널로 바꿔버림
			n = nil
			return n
		} else if n.left == nil { // 우측 자식노드만 있는 경우 부모로 올려버림
			n = n.right
			return n
		} else if n.right == nil { // 좌측 자식노드만 있는 경우 부모로 올려버림
			n = n.left
			return n
		} else { // 자식 노드가 2개인 경우 (좌측 자식 노드 중 가장 큰 값, 우측 자식 노드 중 가장 작은 값)
			minNode := findMinNode(n.right) // 우측 자식 노드 중 가장 작은 값(리프 노드임)
			n.value = minNode.value
			n.right = delete(n.right, minNode.value) // 리프노드의 삭제 결과는 nil
			return n
		}
	}
}

// 자식노드중 최소값을 찾기위한 함수
func findMinNode(n *Node) *Node {
	if n.left == nil {
		return n
	}
	return findMinNode(n.left)
}
