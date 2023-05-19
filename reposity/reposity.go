package reposity

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
)

var (
	nextTopicId   int64
	topicIndexMap map[int64]*Topic
	nextPostId    int64
	postIndexMap  map[int64][]*Post
)

func InitData(filePath string) error {
	if err := initTopicIndexMap(filePath); err != nil {
		log.Panic("init topic index map error!", err)
	}
	if err := initPostIndexMap(filePath); err != nil {
		log.Panic("init post index map error!", err)
	}
	return nil
}

func initTopicIndexMap(filePath string) error {
	open, err := os.Open(filePath + "topic")
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(open)
	topicTmpMap := make(map[int64]*Topic)
	var nextTmpTopicId int64
	for scanner.Scan() {
		text := scanner.Text()
		var topic Topic
		if err := json.Unmarshal([]byte(text), &topic); err != nil {
			return err
		}
		topicTmpMap[topic.Id] = &topic
		nextTmpTopicId = int64(math.Max(float64(nextTmpTopicId), float64(topic.Id)))
	}
	topicIndexMap = topicTmpMap
	nextTopicId = nextTmpTopicId + 1
	s, _ := json.Marshal(topicIndexMap)
	fmt.Println(string(s))
	return nil
}

func initPostIndexMap(filePath string) error {
	open, err := os.Open(filePath + "post")
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(open)
	var nextTmpPostId int64
	postTmpMap := make(map[int64][]*Post)
	for scanner.Scan() {
		text := scanner.Text()
		var postList []*Post
		if err := json.Unmarshal([]byte(text), &postList); err != nil {
			return err
		}
		postTmpMap[postList[0].TopicId] = postList
		for _, post := range postList {
			nextTmpPostId = int64(math.Max(float64(nextTmpPostId), float64(post.Id)))
		}
	}
	postIndexMap = postTmpMap
	nextPostId = nextTmpPostId + 1
	s, _ := json.Marshal(topicIndexMap)
	fmt.Println(string(s))
	return nil
}
