package binarysearchtree

import (
	"testing"

	. "github.com/franela/goblin"
)

func Test(t *testing.T) {
	g := Goblin(t)

	g.Describe("Binary Search Tree", func() {
		g.It("Insert Test", func() {
			bst := NewBST()
			bst.Insert(10)
			bst.Insert(5)
			bst.Insert(15)
			bst.Insert(1)
			bst.Insert(7)

			root := bst.root
			five := root.left
			fifteen := root.right
			one := five.left
			seven := five.right

			g.Assert(root.value).Equal(10)
			g.Assert(five.value).Equal(5)
			g.Assert(fifteen.value).Equal(15)
			g.Assert(one.value).Equal(1)
			g.Assert(seven.value).Equal(7)
		})

		g.It("Delete, Search Test", func() {
			bst := NewBST()
			bst.Insert(10)
			bst.Insert(5)
			bst.Insert(15)
			bst.Insert(1)
			bst.Insert(7)

			bst.Delete(5)
			five := bst.Search(5)
			g.Assert(five).IsNil()
		})
	})
}
