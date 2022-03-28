package cache_test

import (
	"math/rand"
	"testing"

	"github.com/deep2essence/cache"
)

const (
	cnt_put  = 100000
	cnt_get  = 0
	capacity = 1000
)

func TestNodeList(t *testing.T) {
	// t.Run("Test Put", func(t *testing.T) {
	// 	var c = cache.NodeList{}
	// 	for i := 0; i < 1000000; i++ {
	// 		k := rand.Int()

	// 		v := cache.Data([]byte("xxxx"))
	// 		c.Put(k, &v)
	// 	}

	// })

	t.Run("Test Get", func(t *testing.T) {
		var c = cache.NodeList{}
		for i := 0; i < cnt_put; i++ {
			k := rand.Int()

			v := cache.Data([]byte("xxxx"))
			c.Put(k, &v)
		}
		for i := 0; i < cnt_get; i++ {
			k := rand.Int()
			c.Get(k)

		}

	})
}

func TestCache(t *testing.T) {

	t.Run("Test Get", func(t *testing.T) {
		var c = cache.NewCache(capacity)
		for i := 0; i < cnt_put; i++ {
			k := rand.Int()

			v := cache.Data([]byte("xxxx"))
			c.Put(k, &v)
		}
		for i := 0; i < cnt_get; i++ {
			k := rand.Int()
			c.Get(k)

		}

	})
}

func TestHashMap(t *testing.T) {

	t.Run("Test Get", func(t *testing.T) {
		var c = make(map[int]cache.Data)
		for i := 0; i < cnt_put; i++ {
			k := rand.Int()

			v := cache.Data([]byte("xxxx"))
			c[k] = v
		}
		for i := 0; i < cnt_get; i++ {
			k := rand.Int()
			var _ = c[k]

		}

	})
}
