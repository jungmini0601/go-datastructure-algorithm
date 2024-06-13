package heap

import "fmt"

type Heap struct {
	array []int
}

func (h *Heap) Insert(key int) {
	h.array = append(h.array, key)
	h.heapifyUp(len(h.array) - 1)
}

func (h *Heap) heapifyUp(index int) {
	// 부모 노드의 값이 자식 노드의 값 보다 작을 경우 부모노드와 자식 노드의 값을 바꿔준다.
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

func (h *Heap) Extract() int {
	if len(h.array) == 0 {
		fmt.Println("Cannot Extract because heap length is 0")
		return -1
	}

	extracted := h.array[0]
	last := len(h.array) - 1
	h.array[0] = h.array[last]
	h.array = h.array[:last]

	h.heapifyDown(0)
	return extracted
}

func (h *Heap) heapifyDown(index int) {
	lastIndex := len(h.array) - 1
	leftIndex, rightIndex := left(index), right(index)
	childToCompare := 0

	for leftIndex <= lastIndex {
		// 자식 노드중 큰 값을 위로 올리기 위해서 childToCompare 값을 더 큰 노드의 인덱스로 설정한다.
		if leftIndex == lastIndex {
			childToCompare = leftIndex
		} else if h.array[leftIndex] > h.array[rightIndex] {
			childToCompare = leftIndex
		} else {
			childToCompare = rightIndex
		}
		// 부모노드의 값이 더 클 경우 값을 스왑하고 인덱스를 자식 노드의 인덱스로 증가시킨다.
		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			leftIndex, rightIndex = left(index), right(index)
		} else {
			return
		}
	}
}

// 주어진 인덱스의 부모 노드의 인덱스
func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}

func (h *Heap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}
