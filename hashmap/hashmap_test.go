package hashmap

import (
	"testing"

	. "github.com/franela/goblin"
)

func Test(t *testing.T) {
	g := Goblin(t)
	g.Describe("HashMap", func() {
		g.It("Insert Test", func() {
			hm := NewHashMap()
			hm.Insert("key1", "value1")
			hm.Insert("key2", "value2")
			val1, _ := hm.Get("key1")
			val2, _ := hm.Get("key2")
			g.Assert(val1).Equal("value1")
			g.Assert(val2).Equal("value2")
		})

		g.It("Delete Test", func() {
			hm := NewHashMap()
			hm.Insert("key1", "value1")
			hm.Insert("key2", "value2")
			hm.Delete("key1")
			_, err := hm.Get("key1")
			g.Assert(err).IsFalse()
		})
	})
}
