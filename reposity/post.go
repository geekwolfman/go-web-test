package reposity

import "time"

type Post struct {
	Id         int64     `json:"id"`
	TopicId    int64     `json:"topic_id"`
	Content    string    `json:"content"`
	CreateTime time.Time `json:"create_time"`
}
