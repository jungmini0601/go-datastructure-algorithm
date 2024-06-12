package stack

import (
	"testing"

	. "github.com/franela/goblin"
)

func Test(t *testing.T) {
	g := Goblin(t)
	g.Describe("Stack", func() {
		g.It("should push and pop", func() {
			s := NewStack()
			s.Push(1)
			s.Push(2)
			g.Assert(s.Pop()).Equal(2)
			g.Assert(s.Pop()).Equal(1)
		})
	})
}
