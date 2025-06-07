package pages

import "time"

type _About struct {
	Title   string
	Content string
	Time    time.Time
}

func NewAbout(title, content string, time time.Time) *_About {
	return &_About{
		Title:   title,
		Content: content,
		Time:    time,
	}
}
