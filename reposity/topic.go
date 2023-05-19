package reposity

import "time"

type Topic struct {
	Id         int64     `json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	CreateTime time.Time `json:"create_time"`
}
