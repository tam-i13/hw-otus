package hw04lrucache

type Key string

type KeyValue struct {
	ValueStruct interface{}
	KeyStruct   Key
}

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

	if _, ok := l.items[key]; ok {
		flag = true
	}

	if !flag {
		if l.capacity == l.queue.Len() {
			backItem := l.queue.Back()
			delete(l.items, backItem.Value.(KeyValue).KeyStruct)
			l.queue.Remove(backItem)
		}
	}

	l.items[key] = l.queue.PushFront(KeyValue{
		KeyStruct:   key,
		ValueStruct: value,
	})
	println(l.items[key])
	return flag
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	flag := false

	if _, ok := l.items[key]; ok {
		flag = true
	}

	if flag {
		l.queue.MoveToFront(l.items[key])
		tmp := l.items[key]
		return tmp.Value.(KeyValue).ValueStruct, flag
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
