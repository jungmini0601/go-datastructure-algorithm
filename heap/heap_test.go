package heap

import (
	"testing"

	. "github.com/franela/goblin"
)

func Test(t *testing.T) {
	g := Goblin(t)
	g.Describe("Test", func() {
		g.It("Heap Test", func() {
			maxHeap := &Heap{}
			nodes := []int{10, 20, 30, 5, 45, 87, 78, 12, 55}
			for _, v := range nodes {
				maxHeap.Insert(v)
			}

			g.Assert(maxHeap.Extract()).Equal(87)
			g.Assert(maxHeap.Extract()).Equal(78)
			g.Assert(maxHeap.Extract()).Equal(55)
			g.Assert(maxHeap.Extract()).Equal(45)
			g.Assert(maxHeap.Extract()).Equal(30)
			g.Assert(maxHeap.Extract()).Equal(20)
			g.Assert(maxHeap.Extract()).Equal(12)
			g.Assert(maxHeap.Extract()).Equal(10)
			g.Assert(maxHeap.Extract()).Equal(5)
		})
	})
}
