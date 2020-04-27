package lru

import "container/list"

type Cache struct {
	maxBytes int64 // 允许使用的最大内存
	nbytes   int64 // 当前已使用的内存
	ll       *list.List
	cache    map[string]*list.Element
	// optional and executed when an entry is purged.
	OnEvicted func(key string, value Value) // 某条记录被移除时的回调函数
}

type entry struct {
	key   string
	value Value
}

// Value use Len to count how many bytes it takes
type Value interface {
	Len() int
}

func New(maxBytes int64, OnEvicted func(string, value)) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		ll:        list.New(),
		OnEvicted: OnEvicted,
	}
}
