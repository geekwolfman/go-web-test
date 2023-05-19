package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"go-web-test/reposity"
	"log"
	"os"
	"time"
)

func main() {
	// 初始化topic数据
	var topicList []*reposity.Topic
	for i := 0; i < 10; i++ {
		topicList = append(topicList, &reposity.Topic{
			Id:         int64(i),
			Title:      fmt.Sprintf("%d", i),
			Content:    fmt.Sprintf("%d", i),
			CreateTime: time.Now(),
		})
	}
	topicOpen, err := os.OpenFile("data/topic", os.O_RDWR|os.O_CREATE, 0)
	defer topicOpen.Close()
	if err != nil {
		log.Panic(err)
	}
	topicWriter := bufio.NewWriter(topicOpen)
	defer topicWriter.Flush()
	for _, topic := range topicList {
		if t, e := json.Marshal(topic); e == nil {
			_, _ = topicWriter.WriteString(string(t) + "\n")
		}
	}

	// 初始化post数据
	var postList [10][]*reposity.Post
	var id int64
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			postList[i] = append(postList[i], &reposity.Post{
				Id:         id,
				TopicId:    int64(i),
				Content:    fmt.Sprintf("%d-%d", i, j),
				CreateTime: time.Now(),
			})
			id++
		}
	}
	postOpen, err := os.OpenFile("data/post", os.O_RDWR|os.O_CREATE, 0)
	defer postOpen.Close()
	if err != nil {
		log.Panic(err)
	}
	postWriter := bufio.NewWriter(postOpen)
	defer postWriter.Flush()
	for _, post := range postList {
		if t, e := json.Marshal(post); e == nil {
			_, _ = postWriter.WriteString(string(t) + "\n")
		}
	}
	reposity.InitData("data/")
}
