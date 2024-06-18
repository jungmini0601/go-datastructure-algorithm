package avltree

import (
	"testing"

	. "github.com/franela/goblin"
)

func Test(t *testing.T) {
	g := Goblin(t)

	g.Describe("AVL Tree", func() {
		g.It("Insert Test", func() {
			tree := new(AvlTree)
			list := []int{10, 20, 30, 40, 50, 25}

			for _, v := range list {
				tree.Insert(NewAVLTree(v))
			}

			g.Assert(tree.root.value).Equal(30)
			g.Assert(tree.root.left.value).Equal(20)
			g.Assert(tree.root.right.value).Equal(40)
			g.Assert(tree.root.right.right.value).Equal(50)
			g.Assert(tree.root.left.left.value).Equal(10)
			g.Assert(tree.root.left.right.value).Equal(25)

		})
	})
}
