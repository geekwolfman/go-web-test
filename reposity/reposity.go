package reposity

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
)

var (
	topicIndexMap map[int64]*Topic
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
	for scanner.Scan() {
		text := scanner.Text()
		var topic Topic
		if err := json.Unmarshal([]byte(text), &topic); err != nil {
			return err
		}
		topicTmpMap[topic.Id] = &topic
	}
	topicIndexMap = topicTmpMap
	return nil
}

func initPostIndexMap(filePath string) error {
	open, err := os.Open(filePath + "post")
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(open)
	postTmpMap := make(map[int64][]*Post)
	for scanner.Scan() {
		text := scanner.Text()
		var postList []*Post
		if err := json.Unmarshal([]byte(text), &postList); err != nil {
			return err
		}
		postTmpMap[postList[0].TopicId] = postList
	}
	postIndexMap = postTmpMap
	return nil
}
