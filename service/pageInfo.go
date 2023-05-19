package service

import "go-web-test/reposity"

type PageInfo struct {
	Topic    *reposity.Topic  `json:"topic"`
	PostList []*reposity.Post `json:"post_list"`
}
