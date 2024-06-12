package doublylinkedlist

import "strconv"

// 이중 연결 리스트 노드
type Node struct {
	data int   // 데이터
	prev *Node // 이전 원소 참조
	next *Node // 앞 원소 참조
}

type DoublyLinkedList struct {
	head *Node // 앞 부분
	tail *Node // 꼬리 부분
}

// 새로운 링크드 리스트 생성
func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

// 맨 뒤에 원소 추가
func (list *DoublyLinkedList) AppendBack(data int) {
	// 새로운 노드 생성
	newNode := &Node{data: data, prev: nil, next: nil}
	// 리스트에 원소가 없을 경우
	if list.head == nil {
		list.head = newNode
		list.tail = newNode
		return
	}
	// 리스트에 원소가 있을 경우
	newNode.prev = list.tail
	list.tail.next = newNode
	list.tail = newNode
}

// 맨 앞에 원소 추가
func (list *DoublyLinkedList) AppendFront(data int) {
	// 새로운 노드 생성
	newNode := &Node{data: data, prev: nil, next: nil}
	// 리스트에 원소가 없을 경우
	if list.head == nil {
		list.head = newNode
		list.tail = newNode
		return
	}
	// 리스트에 원소가 있을 경우
	newNode.next = list.head
	list.head.prev = newNode
	list.head = newNode
}

// 맨 앞 원소 삭재
func (list *DoublyLinkedList) RemoveFront() {
	// 리스트에 원소가 없을 경우
	if list.head == nil {
		return
	}
	// 리스트에 원소가 하나일 경우
	if list.head == list.tail {
		list.head = nil
		list.tail = nil
		return
	}
	// 리스트에 원소가 여러개일 경우
	list.head = list.head.next
	list.head.prev = nil
}

// 맨 뒤 원소 삭제
func (list *DoublyLinkedList) RemoveBack() {
	// 리스트에 원소가 없을 경우
	if list.head == nil {
		return
	}
	// 리스트에 원소가 하나일 경우
	if list.head == list.tail {
		list.head = nil
		list.tail = nil
		return
	}

	// 리스트에 원소가 여러 개일 경우
	list.tail = list.tail.prev
	list.tail.next = nil
}

// 특정 값을 가지는 원소 삭제
func (list *DoublyLinkedList) Remove(data int) {
	// 리스트에 원소가 없을 경우
	if list.head == nil {
		return
	}
	// 리스트에 원소가 하나일 경우
	if list.head == list.tail {
		if list.head.data == data {
			list.head = nil
			list.tail = nil
			return
		}
		return
	}
	// 리스트에 원소가 여러개일 경우
	current := list.head
	for current != nil {
		if current.data == data {
			// 삭제할 노드가 끝 원소인 경우
			if current == list.tail {
				list.tail = current.prev
				list.tail.next = nil
				return
			} else if current == list.head {
				// 삭제 할 원소가 맨 앞 원소인 경우
				list.head = current.next
				list.head.prev = nil
				return
			} else {
				// 삭제할 노드가 중간 원소인 경우
				current.prev.next = current.next
				current.next.prev = current.prev
				return
			}
		}
		current = current.next
	}
}

func (list *DoublyLinkedList) ToString() string {
	if list.head == nil {
		return "[]"
	}

	var ret = "["
	current := list.head

	for current != nil {
		ret += strconv.Itoa(current.data)
		if current.next == nil {
			ret += "]"
		} else {
			ret += "-"
		}
		current = current.next
	}

	return ret
}
