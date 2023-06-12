package main

import (
	"sync/atomic"
	"time"
)

type Book struct {
	ID        uint32        `json:"id"`
	Title     string        `json:"title"`
	Author    string        `json:"author"`
	Status    ReadingStatus `json:"status"`
	Rating    int           `json:"rating"`
	CreatedAt time.Time     `json:"created_at"`
}

func CreateBook(title string, author string) *Book {
	b := &Book{
		ID:        atomic.AddUint32(&ID, 1),
		Title:     title,
		Author:    author,
		Status:    ToRead,
		Rating:    1,
		CreatedAt: time.Now(),
	}
	cache.AddBook(*b)
	return b
}

func UpdateStatus(id uint32, status ReadingStatus) {
	b := cache.GetBookByID(id)
	b.Status = status
	cache.UpdateBook(b)
}
