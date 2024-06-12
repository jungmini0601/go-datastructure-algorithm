package queue

import (
	"testing"

	. "github.com/franela/goblin"
)

func Test(t *testing.T) {
	g := Goblin(t)
	g.Describe("Stack", func() {
		g.It("should push and pop", func() {
			q := NewQueue()
			q.Push(1)
			q.Push(2)
			g.Assert(q.Pop()).Equal(1)
			g.Assert(q.Pop()).Equal(2)
		})
	})
}
