package library

import (
	"fmt"
)

type AudioBook struct {
	Title       string
	Author      string
	Description string
	Chapters    []AudioBookChapter
}

func (b *AudioBook) String() string {
	return fmt.Sprintf("AudioBook{title: %s, author: %s}", b.Title, b.Author)
}
