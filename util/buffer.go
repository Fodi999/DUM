//util/buffer.go
package util

import (
    "errors"
    "sync"
)

type CircularBuffer struct {
    buffer []string
    size   int
    start  int
    end    int
    count  int
    mutex  sync.Mutex
}

func NewCircularBuffer(size int) *CircularBuffer {
    return &CircularBuffer{
        buffer: make([]string, size),
        size:   size,
    }
}

func (cb *CircularBuffer) Add(item string) {
    cb.mutex.Lock()
    defer cb.mutex.Unlock()

    cb.buffer[cb.end] = item
    cb.end = (cb.end + 1) % cb.size

    if cb.count == cb.size {
        cb.start = (cb.start + 1) % cb.size
    } else {
        cb.count++
    }
}

func (cb *CircularBuffer) Get() (string, error) {
    cb.mutex.Lock()
    defer cb.mutex.Unlock()

    if cb.count == 0 {
        return "", errors.New("buffer is empty")
    }

    item := cb.buffer[cb.start]
    cb.start = (cb.start + 1) % cb.size
    cb.count--

    return item, nil
}

func (cb *CircularBuffer) GetAll() []string {
    cb.mutex.Lock()
    defer cb.mutex.Unlock()

    var items []string
    index := cb.start
    for i := 0; i < cb.count; i++ {
        items = append(items, cb.buffer[index])
        index = (index + 1) % cb.size
    }
    return items
}

func (cb *CircularBuffer) IsFull() bool {
    cb.mutex.Lock()
    defer cb.mutex.Unlock()
    return cb.count == cb.size
}

func (cb *CircularBuffer) IsEmpty() bool {
    cb.mutex.Lock()
    defer cb.mutex.Unlock()
    return cb.count == 0
}

