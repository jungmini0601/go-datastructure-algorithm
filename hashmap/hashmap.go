package hashmap

const MAP_SIZE = 50

// HashMap에 사용되는 노드
type Node struct {
	key   string
	value string
	next  *Node
}

// HashMap은 연결리스트의 배열로 이루어져 있음.
type HashMap struct {
	data []*Node
}

func NewHashMap() *HashMap {
	return &HashMap{data: make([]*Node, MAP_SIZE)}
}

// Jenkins hash function(뭔지 잘 모름 코드 가져온거)
func hash(key string) (hash uint32) {
	hash = 0
	for _, ch := range key {
		hash += uint32(ch)
		hash += hash << 10
		hash ^= hash >> 6
	}

	hash += hash << 3
	hash ^= hash >> 11
	hash += hash << 15

	return hash
}

func getIndex(key string) (index int) {
	return int(hash(key)) % MAP_SIZE
}

func (h *HashMap) Insert(key string, value string) {
	index := getIndex(key)

	if h.data[index] == nil {
		h.data[index] = &Node{key: key, value: value}
	} else {
		starting_node := h.data[index]
		for starting_node.next != nil {
			// 이미 존재하는 키가 있다면 값을 바꿔준다.
			if starting_node.key == key {
				starting_node.value = value
				return
			}
			starting_node = starting_node.next
		}
		// 해시충돌만 있고 끝 값에 도달한 경우
		starting_node.next = &Node{key: key, value: value}
	}
}

func (h *HashMap) Get(key string) (string, bool) {
	index := getIndex(key)
	if h.data[index] != nil {
		// 키가 연결리스트 어딘가에 존재한다.
		starting_node := h.data[index]
		// 첫 번째 원소가 키 값인 경우
		if starting_node.key == key {
			return starting_node.value, true
		}

		for starting_node.next != nil {
			if starting_node.key == key {
				return starting_node.value, true
			}
			starting_node = starting_node.next
		}
	}

	// 키 값이 없는 경우
	return "", false
}

func (h *HashMap) Delete(key string) {
	index := getIndex(key)
	if h.data[index] != nil {
		// 키가 연결리스트 어딘가에 존재한다.
		starting_node := h.data[index]
		// 첫 번째 원소가 키 값인 경우
		if starting_node.key == key {
			h.data[index] = starting_node.next
			return
		}

		for starting_node.next != nil {
			if starting_node.next.key == key {
				starting_node.next = starting_node.next.next
				return
			}
			starting_node = starting_node.next
		}
	}
}
