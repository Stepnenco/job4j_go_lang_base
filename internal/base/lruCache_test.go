package base

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Put(t *testing.T) {
	t.Run("Add to empty cache", func(t *testing.T) {
		cache := NewLruCache(3)
		cache.Put("key1", "value1")

		assert.NotNil(t, cache.Head, "Head should not be nil")
		assert.NotNil(t, cache.Tail, "Tail should not be nil")
		assert.Equal(t, "key1", cache.Head.Key, "Head should have key 'key1'")
		assert.Equal(t, "key1", cache.Tail.Key, "Tail should have key 'key1'")
		assert.Equal(t, "value1", cache.Head.Value, "Head value should be 'value1'")
		assert.Equal(t, "value1", cache.Tail.Value, "Tail value should be 'value1'")
		assert.Nil(t, cache.Head.Prev, "Head.Prev should be nil")
		assert.Nil(t, cache.Head.Next, "Head.Next should be nil for single node")
		assert.Nil(t, cache.Tail.Prev, "Tail.Prev should be nil for single node")
		assert.Nil(t, cache.Tail.Next, "Tail.Next should be nil for single node")
	})

	t.Run("Add multiple elements", func(t *testing.T) {
		cache := NewLruCache(3)
		cache.Put("a", "1")
		cache.Put("b", "2")
		cache.Put("c", "3")

		assert.Equal(t, "c", cache.Head.Key, "Head should be 'c' (most recent)")
		assert.Equal(t, "a", cache.Tail.Key, "Tail should be 'a' (oldest)")
		assert.Equal(t, "b", cache.Head.Next.Key, "Second node should be 'b'")
		assert.Equal(t, "c", cache.Head.Next.Prev.Key, "b.Prev should point to 'c'")
		assert.Equal(t, "a", cache.Head.Next.Next.Key, "Third node should be 'a'")
		assert.Equal(t, "b", cache.Tail.Prev.Key, "Tail.Prev should be 'b'")

		assert.Equal(t, "3", cache.Head.Value, "Head value should be '3'")
		assert.Equal(t, "2", cache.Head.Next.Value, "Second node value should be '2'")
		assert.Equal(t, "1", cache.Tail.Value, "Tail value should be '1'")
	})

	t.Run("Exceed size limit", func(t *testing.T) {
		cache := NewLruCache(2)
		cache.Put("a", "1")
		cache.Put("b", "2")
		cache.Put("c", "3")

		assert.Equal(t, "c", cache.Head.Key, "Head should be 'c' after overflow")
		assert.Equal(t, "b", cache.Tail.Key, "Tail should be 'b' after overflow")
		assert.Equal(t, "c", cache.Tail.Prev.Key, "Tail.Prev should be 'c'")
		assert.Equal(t, "b", cache.Head.Next.Key, "Head.Next should be 'b'")
		assert.Nil(t, cache.Tail.Next, "Tail.Next should be nil")

	})
}

func Test_Get(t *testing.T) {
	t.Run("Get existing element", func(t *testing.T) {
		cache := NewLruCache(3)
		cache.Put("a", "1")
		cache.Put("b", "2")
		cache.Put("c", "3")

		val := cache.Get("b")
		assert.NotNil(t, val, "Get should return non-nil for existing key")
		assert.Equal(t, "2", *val, "Get should return correct value")

		assert.Equal(t, "b", cache.Head.Key, "'b' should be at Head after Get")
		assert.Equal(t, "c", cache.Head.Next.Key, "Second node should be 'c'")
		assert.Equal(t, "a", cache.Tail.Key, "'a' should be at Tail")
	})

	t.Run("Get non-existing element", func(t *testing.T) {
		cache := NewLruCache(3)
		cache.Put("a", "1")

		val := cache.Get("b")
		assert.Nil(t, val, "Get should return nil for non-existing key")

		// Структура не должна измениться
		assert.Equal(t, "a", cache.Head.Key, "Head should remain 'a'")
		assert.Equal(t, "a", cache.Tail.Key, "Tail should remain 'a'")
	})

	t.Run("Get from empty cache", func(t *testing.T) {
		cache := NewLruCache(3)

		val := cache.Get("any")
		assert.Nil(t, val, "Get should return nil for empty cache")
		assert.Nil(t, cache.Head, "Head should be nil for empty cache")
		assert.Nil(t, cache.Tail, "Tail should be nil for empty cache")
	})

}
