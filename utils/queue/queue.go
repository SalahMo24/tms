package queue

import (
	"errors"
	"sync"
)

// InMemoryQueue implements a simple in-memory queue
type InMemoryQueue struct {
	mu    sync.Mutex
	queue []queueItem
}

type queueItem struct {
	id   string
	data []byte
}

func NewInMemoryQueue() *InMemoryQueue {
	return &InMemoryQueue{
		queue: make([]queueItem, 0),
	}
}

func (q *InMemoryQueue) EnqueueTransaction(id string, data []byte) error {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.queue = append(q.queue, queueItem{id: id, data: data})
	return nil
}

func (q *InMemoryQueue) DequeueTransaction() (string, []byte, error) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.queue) == 0 {
		return "", nil, errors.New("queue is empty")
	}

	item := q.queue[0]
	q.queue = q.queue[1:]
	return item.id, item.data, nil
}
