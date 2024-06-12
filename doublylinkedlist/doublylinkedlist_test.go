package doublylinkedlist

import (
	"testing"

	. "github.com/franela/goblin"
)

func Test(t *testing.T) {
	g := Goblin(t)
	g.Describe("DoublyLinkedList", func() {
		g.It("Empty list", func() {
			list := NewDoublyLinkedList()
			g.Assert(list.ToString()).Equal("[]")
		})

		g.It("Append Back 1 2 3 4 5", func() {
			list := NewDoublyLinkedList()
			for i := 1; i <= 5; i++ {
				list.AppendBack(i)
			}

			g.Assert(list.ToString()).Equal("[1-2-3-4-5]")
		})

		g.It("Append Front 1 2 3 4 5", func() {
			list := NewDoublyLinkedList()
			for i := 1; i <= 5; i++ {
				list.AppendFront(i)
			}

			g.Assert(list.ToString()).Equal("[5-4-3-2-1]")
		})

		g.It("Append Front 1 2 3 4 5 RemoveBack", func() {
			list := NewDoublyLinkedList()
			for i := 1; i <= 5; i++ {
				list.AppendFront(i)
			}

			list.RemoveBack()
			g.Assert(list.ToString()).Equal("[5-4-3-2]")
		})

		g.It("Append Front 1 2 3 4 5 RemoveFront", func() {
			list := NewDoublyLinkedList()
			for i := 1; i <= 5; i++ {
				list.AppendFront(i)
			}

			list.RemoveFront()
			g.Assert(list.ToString()).Equal("[4-3-2-1]")
		})

		g.It("Remove Test - 맨 뒤 원소", func() {
			list := NewDoublyLinkedList()
			for i := 1; i <= 5; i++ {
				list.AppendFront(i)
			}

			list.Remove(1)
			g.Assert(list.ToString()).Equal("[5-4-3-2]")
		})

		g.It("Remove Test - 맨 앞 원소", func() {
			list := NewDoublyLinkedList()
			for i := 1; i <= 5; i++ {
				list.AppendFront(i)
			}

			list.Remove(5)
			g.Assert(list.ToString()).Equal("[4-3-2-1]")
		})

		g.It("Remove Test - 삭제할 원소가 중간에 있을 경우", func() {
			list := NewDoublyLinkedList()
			for i := 1; i <= 5; i++ {
				list.AppendFront(i)
			}

			list.Remove(3)
			g.Assert(list.ToString()).Equal("[5-4-2-1]")
		})
	})
}
