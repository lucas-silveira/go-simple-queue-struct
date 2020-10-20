package main

import "sync"

// Item ...
type Item interface{}

// Queue ...
type Queue struct {
	items []Item
	mutex sync.Mutex
}

// Add ...
func (queue *Queue) Add(item Item) {
	queue.mutex.Lock()
	defer queue.mutex.Unlock()

	queue.items = append(queue.items, item)
}

// Remove ...
func (queue *Queue) Remove() Item {
	queue.mutex.Lock()
	defer queue.mutex.Unlock()

	if len(queue.items) == 0 {
		return nil
	}

	lastItem := queue.items[0]
	queue.items = queue.items[1:]

	return lastItem
}

// IsEmpty ...
func (queue *Queue) IsEmpty() bool {
	queue.mutex.Lock()
	defer queue.mutex.Unlock()

	return len(queue.items) == 0
}

// Reset ...
func (queue *Queue) Reset() {
	queue.mutex.Lock()
	defer queue.mutex.Unlock()

	queue.items = nil
}

// Dump ...
func (queue *Queue) Dump() []Item {
	queue.mutex.Lock()
	defer queue.mutex.Unlock()

	copiedQueue := make([]Item, len(queue.items))
	copy(copiedQueue, queue.items)

	return copiedQueue
}

// Peek ...
func (queue *Queue) Peek() Item {
	queue.mutex.Lock()
	defer queue.mutex.Unlock()

	if len(queue.items) == 0 {
		return nil
	}

	return queue.items[len(queue.items)-1]
}
