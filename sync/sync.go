package sync

import "sync"

type Counter struct {
	RWMutex sync.RWMutex
	Number  int
}

func (c *Counter) Inc() {
	c.RWMutex.Lock()
	defer c.RWMutex.Unlock()

	c.Number += 1
}

func (c *Counter) Value() int {
	c.RWMutex.RLock()
	num := c.Number
	c.RWMutex.RUnlock()
	return num
}
