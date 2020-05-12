package sync

import "sync"

// Counter type is a struct to hold counter informations
type Counter struct {
	sync.Mutex
	value int
}

// Value return the values of a counter
func (c *Counter) Value() int {
	return c.value
}

// Inc increment the counter by 1
func (c *Counter) Inc() {
	c.Lock()
	defer c.Unlock()
	c.value++
}

// NewCounter Initialize one counter (cause have mutex)
func NewCounter() *Counter {
	return &Counter{}
}
