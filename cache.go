package lru

import (
	"container/list"
	"sync"
)

const (
	defaultSize = 1024
)

type entry struct {
	key   string
	value interface{}
}

type cache struct {
	mu     sync.RWMutex
	size   int
	values *list.List
	items  map[string]*list.Element
}

func New() *cache {
	return &cache{
		size:   defaultSize,
		values: new(list.List),
		items:  make(map[string]*list.Element, defaultSize),
	}
}

func (c *cache) Set(key string, value interface{}) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	if node, ok := c.items[key]; ok {
		c.values.MoveToFront(node)
		node.Value.(*entry).value = &entry{key: key, value: value}
		return false
	}

	e := c.values.PushFront(&entry{key, value})
	c.items[key] = e
	if c.values.Len() > c.size {
		back := c.values.Back()
		c.values.Remove(back)
		delete(c.items, back.Value.(*entry).key)
	}
	return true
}

func (c *cache) Get(key string) (interface{}, bool) {
	if node, ok := c.items[key]; ok {
		c.values.MoveToFront(node)
		return node.Value.(*entry).value, true
	}

	return nil, false
}

func (c *cache) Del(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if node, ok := c.items[key]; ok {
		delete(c.items, key)
		c.values.Remove(node)
	}
}
