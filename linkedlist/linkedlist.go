package linkedlist

import (
	"fmt"
	"strconv"
)

// 데이터를 나타내는 구조체
type Node struct {
	data int
	next *Node
}

// 링크드 리스트
type LinkedList struct {
	head *Node
}

// 링크드 리스트 원소 추가
func (list *LinkedList) AddNode(data int) {
	newNode := &Node{data, nil}
	// 첫 번째 원소인 경우
	if list.head == nil {
		list.head = newNode
		return
	}
	// 원소가 1개 이상 존재하는 경우 마지막 노드가 current가 되도록 순회
	current := list.head
	for current.next != nil {
		current = current.next
	}
	// 마지막 노드의 next 참조에 새로운 노드를 연결한다.
	current.next = newNode
}

// 링크드 리스트 원소 제거
func (list *LinkedList) RemoveNode(data int) {
	// 링크드 리스트가 비어 있을 경우
	if list.head == nil {
		return
	}
	// 첫 번째 원소가 삭제 될 경우 head의 참조를 다음으로 넘기면 참조가 알아서 사라짐
	if list.head.data == data {
		list.head = list.head.next
		return
	}
	// 삭제할 노드의 이전 노드의 참조를 찾을 것이다.
	// [1-2-3-4-5] <- 3을 삭제할 경우 prev참조를 2에 가져다 놓는 것이 목표이다.
	prev := list.head
	for prev.next != nil {
		// prev 노드의 다음 노드가 삭제될 노드라면 다음 다음 노드를 연결한다.
		// 2 - 3 - 4     ->   2 - 4
		if prev.next.data == data {
			prev.next = prev.next.next
			return
		}
		prev = prev.next
	}
}

// 출력용 코드
func (list *LinkedList) PrintList() {
	current := list.head
	for current != nil {
		fmt.Print(current.data, " ")
		current = current.next
	}
	fmt.Println()
}

// 문자열 변환용 코드
func (list *LinkedList) toString() string {
	if list.head == nil {
		return "[]"
	}

	var ret = "["
	current := list.head

	for current != nil {
		ret += strconv.Itoa(current.data)
		current = current.next
		if current == nil {
			ret += "]"
		} else {
			ret += "-"
		}
	}

	return ret
}
