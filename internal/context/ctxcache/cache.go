package ctxcache

import (
	"sync"
)

type Cacher interface {
	Get(key any) any
	GetOrCache(key any, fn func(key any) (any, error)) (any, error)
	Set(key any, value any)
	Release()
}

type cache struct {
	rw   sync.RWMutex
	data map[any]any
}

func newCahe() *cache {
	return &cache{
		data: make(map[any]any, 10),
	}
}

func (c *cache) Get(key any) any {
	c.rw.RLock()
	defer c.rw.RUnlock()
	val, ok := c.data[key]
	if !ok {
		return nil
	}
	return val
}

func (c *cache) GetOrCache(key any, fn func(key any) (any, error)) (any, error) {
	c.rw.Lock()
	defer c.rw.Unlock()
	val, ok := c.data[key]
	if !ok {
		var err error
		if val, err = fn(key); err != nil {
			return nil, err
		}
		c.data[key] = val
	}
	return val, nil
}

func (c *cache) Set(key any, value any) {
	c.data[key] = value
}

func (c *cache) Release() {
	c.rw.Lock()
	defer c.rw.Unlock()
	clear(c.data)
}

type cacheBlock struct {
	data sync.Map
}

func newCacheBlock() *cacheBlock {
	return &cacheBlock{
		data: sync.Map{},
	}
}

func (c *cacheBlock) Get(key any) Cacher {
	val, ok := c.data.Load(key)
	if !ok {
		val, _ = c.data.LoadOrStore(key, newCahe())
	}
	return val.(Cacher)
}

func (c *cacheBlock) Release() {
	c.data.Range(func(key, value any) bool {
		cache := value.(Cacher)
		cache.Release()
		return true
	})
	c.data.Clear()
}
