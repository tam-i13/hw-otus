package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	List map[*ListItem]*ListItem
}

func (l *list) len() int {
	return len(l.List)
}

func (l *list) Front() *ListItem {
	tmpNode := l.List[0]
	flag := false
	for flag == false {
		if tmpNode.Prev == nil {
			return &tmpNode
		} else {
			tmpNode = *tmpNode.Prev
		}
	}
}

func (l *list) Back() *ListItem {
	tmpNode := l.List[0]
	for false == false {
		if tmpNode.Next == nil {
			return tmpNode
		} else {
			tmpNode = tmpNode.Next
		}
	}
}

func (l *list) PushFront(v interface{}) *ListItem {
	front := l.Front()
	newNode := ListItem{
		Value: v,
		Prev:  nil,
		Next:  front,
	}

	l.List[&newNode] = &newNode
	front.Prev = &newNode

	return &newNode
}

func (l *list) PushBack(v interface{}) *ListItem {
	back := l.Back()
	newNode := ListItem{
		Value: v,
		Prev:  back,
		Next:  nil,
	}
	l.List = append(l.List, &newNode)
	back.Next = &newNode
	return &newNode
}

func (l *list) Remove(i *ListItem) {

	if i.Next != nil && i.Prev != nil {
		i.Next.Prev = i.Prev
		i.Prev.Next = i.Next
	}
	// удаление
}

func (l *list) MoveToFront(i *ListItem) {
	return len(l.List)
}

func NewList() List {
	return new(list)
}
