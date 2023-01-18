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
	List          map[*ListItem]*ListItem
	FrontPosition *ListItem
	BackPosition  *ListItem
}

func (l *list) Len() int {
	return len(l.List)
}

func (l *list) Front() *ListItem {
	if l.Len() == 0 {
		return nil
	}
	return l.FrontPosition
}

func (l *list) Back() *ListItem {
	if l.Len() == 0 {
		return nil
	}
	return l.BackPosition
}

func (l *list) PushFront(v interface{}) *ListItem {
	newNode := ListItem{
		Value: v,
		Prev:  nil,
		Next:  nil,
	}

	if l.Len() != 0 {
		front := l.Front()
		newNode.Next = front
		l.List[&newNode] = &newNode
		front.Prev = &newNode
	} else {
		l.List[&newNode] = &newNode
		l.BackPosition = &newNode
	}

	l.FrontPosition = &newNode

	return &newNode
}

func (l *list) PushBack(v interface{}) *ListItem {
	newNode := ListItem{
		Value: v,
		Prev:  nil,
		Next:  nil,
	}

	if l.Len() != 0 {
		back := l.Back()
		newNode.Prev = back
		l.List[&newNode] = &newNode
		back.Next = &newNode
	} else {
		l.List[&newNode] = &newNode
		l.FrontPosition = &newNode
	}

	l.BackPosition = &newNode

	return &newNode
}

func (l *list) Remove(i *ListItem) {
	if i.Next != nil && i.Prev != nil {
		i.Next.Prev = i.Prev
		i.Prev.Next = i.Next
	}
	if i.Next == nil && i.Prev != nil {
		i.Prev.Next = i.Next
		l.BackPosition = i.Prev
	}
	if i.Next != nil && i.Prev == nil {
		i.Next.Prev = i.Prev
		l.FrontPosition = i.Next
	}

	delete(l.List, i)
}

func (l *list) MoveToFront(i *ListItem) {
	if l.Back() != nil {
		l.Remove(i)
		l.PushFront(i.Value)
	}
}

func NewList() List {
	return &list{
		List:          make(map[*ListItem]*ListItem),
		FrontPosition: nil,
		BackPosition:  nil,
	}
}
