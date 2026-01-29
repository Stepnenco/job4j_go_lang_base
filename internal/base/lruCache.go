package base

type Node struct {
	Key   string
	Value string
	Prev  *Node
	Next  *Node
}

type LruCache struct {
	size int
	Head *Node
	Tail *Node
}

func NewLruCache(size int) *LruCache {
	return &LruCache{
		size: size,
	}
}

func (l *LruCache) Put(key string, value string) {
	prevNode := l.Head
	var foundNode *Node

	for prevNode != nil {
		if prevNode.Key == key {
			foundNode = prevNode
			break
		}
		prevNode = prevNode.Next
	}

	if foundNode != nil {
		foundNode.Value = value

		if foundNode != l.Head {

			if foundNode.Prev != nil {
				foundNode.Prev.Next = foundNode.Next
			}
			if foundNode.Next != nil {
				foundNode.Next.Prev = foundNode.Prev
			}

			if foundNode == l.Tail {
				l.Tail = foundNode.Prev
			}

			foundNode.Prev = nil
			foundNode.Next = l.Head
			if l.Head != nil {
				l.Head.Prev = foundNode
			}
			l.Head = foundNode
		}
		return
	}

	newNode := &Node{
		Key:   key,
		Value: value,
		Next:  l.Head,
	}

	if l.Head != nil {
		l.Head.Prev = newNode
	}
	l.Head = newNode

	if l.Tail == nil {
		l.Tail = newNode
	}

	count := 0
	current := l.Head
	for current != nil {
		count++
		if count > l.size {
			if l.Tail != nil && l.Tail.Prev != nil {
				l.Tail.Prev.Next = nil
				l.Tail = l.Tail.Prev
			} else {
				l.Head = nil
				l.Tail = nil
			}
			break
		}
		current = current.Next
	}
}

func (l *LruCache) Get(key string) *string {

	current := l.Head
	var foundNode *Node

	for current != nil {
		if current.Key == key {
			foundNode = current
			break
		}
		current = current.Next
	}

	if foundNode == nil {
		return nil
	}

	if foundNode != l.Head {
		if foundNode.Prev != nil {
			foundNode.Prev.Next = foundNode.Next
		}
		if foundNode.Next != nil {
			foundNode.Next.Prev = foundNode.Prev
		}

		if foundNode == l.Tail {
			l.Tail = foundNode.Prev
		}

		foundNode.Prev = nil
		foundNode.Next = l.Head
		if l.Head != nil {
			l.Head.Prev = foundNode
		}
		l.Head = foundNode
	}

	return &foundNode.Value
}
