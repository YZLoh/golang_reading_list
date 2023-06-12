package main

import "strings"

type ReadingStatus int

var ID uint32 = 0

const (
	Reading ReadingStatus = iota
	ToRead
	Finished
	All
)

func (r ReadingStatus) String() string {
	return [...]string{"Reading", "To Read", "Finished"}[r]
}

func getStars(num int) string {
	if num == 0 {
		return ""
	}
	return strings.Repeat("☆", num)
}
