package main

import (
	"encoding/json"
	"sync"
)

type BookCache struct {
	Books map[uint32]Book
	mu    sync.Mutex
}

func (c *BookCache) AddBook(b Book) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Books[b.ID] = b
}

func (c *BookCache) GetBookByID(id uint32) Book {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.Books[id]
}

func (c *BookCache) DeleteBookById(id uint32) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.Books, id)
}

func (c *BookCache) UpdateBook(b Book) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Books[b.ID] = b
}

func (c *BookCache) ToJSON() []byte {
	c.mu.Lock()
	defer c.mu.Unlock()
	var books []Book
	for _, v := range c.Books {
		books = append(books, v)
	}
	json, err := json.Marshal(books)
	if err != nil {
		panic(err)
	}
	return json
}
