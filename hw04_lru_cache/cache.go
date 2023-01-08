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
	if l.queue.Len() == 0 {
		l.items[key] = l.queue.PushFront(value)
	} else {
		for k, _ := range l.items {
			if k == key {
				flag = true
			}
		}

		if flag {
			l.items[key].Value = value
			l.queue.MoveToFront(l.items[key])
		}

		if !flag {
			if l.capacity >= l.queue.Len() {
				l.queue.Remove(l.queue.Back())
			}

			l.items[key] = l.queue.PushFront(value)
		}
	}
	return flag

}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	flag := false

	for k, _ := range l.items {
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
