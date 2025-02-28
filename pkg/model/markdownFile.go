package model

type MarkdownFile struct {
	ID      int64  `json:"id"`
	TITLE   string `json:"title"`
	SLUG    string `json:"slug"`
	CONTENT string `json:"content"`
}
