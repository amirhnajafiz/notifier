package cache

import (
	"time"
)

type Message struct {
	Topic   string
	Content string
	Date    time.Time
}

type Cache struct {
	buffer []Message
}

func (c *Cache) Put(message Message) {
	c.buffer = append(c.buffer, message)
}

func (c *Cache) Pull() []Message {
	return c.buffer
}

func (c *Cache) Mock() {
	c.buffer = c.buffer[:0]
}
