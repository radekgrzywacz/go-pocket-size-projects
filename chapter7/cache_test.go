package cache_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"learngo-pockets/genericcache"
	"sync"
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	c := cache.New[int, string](3, time.Second)

	c.Upsert(5, "asd")
	v, found := c.Read(5)
	assert.True(t, found)
	assert.Equal(t, v, "asd")

	c.Upsert(1, "one")
	v, found = c.Read(1)
	assert.True(t, found)
	assert.Equal(t, v, "one")
	assert.IsType(t, "1", v)

}

func TestCache_Parallel_goroutines(t *testing.T) {
	c := cache.New[int, string](3, time.Second)

	const parallelTasks = 10
	wg := sync.WaitGroup{}
	wg.Add(parallelTasks)

	for i := 0; i < parallelTasks; i++ {
		go func(j int) {
			defer wg.Done()
			c.Upsert(4, fmt.Sprint(j))
		}(i)
	}

	wg.Wait()
}

func TestCache_Parallel(t *testing.T) {
	c := cache.New[int, string](3, time.Second)

	t.Run("write six", func(t *testing.T) {
		t.Parallel()
		c.Upsert(6, "six")
	})

	t.Run("write kuus", func(t *testing.T) {
		t.Parallel()
		c.Upsert(6, "kuus")
	})
}

func TestCache_TTL(t *testing.T) {
	t.Parallel()

	c := cache.New[string, string](3, time.Microsecond*100)
	c.Upsert("Norwegian", "Blue")

	got, found := c.Read("Norwegian")
	assert.True(t, found)
	assert.Equal(t, "Blue", got)

	time.Sleep(time.Millisecond * 200)

	got, found = c.Read("Norwegian")
	assert.False(t, found)
	assert.Equal(t, "", got)
}

func TestCache_MaxSize(t *testing.T) {
	t.Parallel()

	c := cache.New[int, int](3, time.Minute)
	c.Upsert(1, 1)
	c.Upsert(2, 2)
	c.Upsert(3, 3)

	got, found := c.Read(1)
	assert.True(t, found)
	assert.Equal(t, 1, got)

	c.Upsert(1, 10)
	c.Upsert(4, 4)

	got, found = c.Read(2)
	assert.False(t, found)
	assert.Equal(t, 0, got)
}
