package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	Cache

	capacity int
	queue    List
	items    map[Key]*ListItem
}

func (l *lruCache) Set(key Key, value interface{}) bool {
	flag := false
	for k := range l.items {
		if k == key {
			flag = true
		}
	}

	if flag {
		l.items[key] = l.queue.PushFront(value)
	}

	if !flag {
		if l.capacity == l.queue.Len() {
			for k, v := range l.items {
				if v == l.queue.Back() {
					delete(l.items, k)
				}
			}
			l.queue.Remove(l.queue.Back())
		}

		l.items[key] = l.queue.PushFront(value)
	}

	return flag
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	flag := false

	for k := range l.items {
		if k == key {
			flag = true
		}
	}

	if flag {
		l.queue.MoveToFront(l.items[key])
		tmp := l.items[key]
		return tmp.Value, flag
	}

	return nil, flag
}

func (l *lruCache) Clear() {
	l.queue = NewList()
	l.items = make(map[Key]*ListItem, l.capacity)
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
