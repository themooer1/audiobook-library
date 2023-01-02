package library

import (
	"fmt"
)

type AudioBookChapter struct {
	Title string
	Index int
	Url   string
}

func (c *AudioBookChapter) String() string {
	return fmt.Sprintf("AudioBookChapter{title: %s, index: %d, path: %s}", c.Title, c.Index, c.Url)
}
