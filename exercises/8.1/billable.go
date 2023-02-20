package billable

import "sync"

// Your implementation of Channel goes here!!
type Channel[T any] struct {
	mutex           *sync.RWMutex
	ch              chan T
	sends, receives int
}

func NewChannel[T any](cap int) *Channel[T] {
	return &Channel[T]{
		mutex: &sync.RWMutex{},
		ch:    make(chan T, cap),
	}
}

func (c *Channel[T]) Sends() int {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	return c.sends
}

func (c *Channel[T]) Receives() int {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	return c.receives
}

func (c *Channel[T]) Send(v T) {
	c.ch <- v

	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.sends++
}

func (c *Channel[T]) Receive() T {
	v := <-c.ch

	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.receives++
	return v
}
