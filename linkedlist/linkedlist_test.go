package linkedlist

import (
	"testing"

	. "github.com/franela/goblin"
)

func Test(t *testing.T) {
	g := Goblin(t)
	g.Describe("LinkedList", func() {
		g.It("Add [1-2-3-4-5]", func() {
			list := &LinkedList{}
			for i := 1; i <= 5; i++ {
				list.AddNode(i)
			}

			list.PrintList()
			g.Assert(list.toString()).Equal("[1-2-3-4-5]")
		})

		g.It("Add [1-2-3-4-5] Remove [3]", func() {
			list := &LinkedList{}
			for i := 1; i <= 5; i++ {
				list.AddNode(i)
			}
			list.RemoveNode(3)
			g.Assert(list.toString()).Equal("[1-2-4-5]")
		})

		g.It("Empty []", func() {
			list := &LinkedList{}
			g.Assert(list.toString()).Equal("[]")
		})
	})
}
