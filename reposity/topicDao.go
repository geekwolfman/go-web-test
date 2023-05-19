package reposity

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

var (
	lock sync.RWMutex
)

type TopicDao struct {
}

var (
	topicDao  *TopicDao
	topicOnce sync.Once
)

func TopicDaoInstance() *TopicDao {
	topicOnce.Do(
		func() {
			topicDao = &TopicDao{}
		})
	return topicDao
}

func (*TopicDao) QueryTopicById(id int64) *Topic {
	lock.RLock()
	defer lock.RUnlock()
	return topicIndexMap[id]
}

func (*TopicDao) AppendTopicList(topicList []Topic) error {
	lock.Lock()
	defer lock.Unlock()
	topicTmpList := make([]string, 0, len(topicList))
	for i := range topicList {
		topicList[i].Id = nextTopicId
		nextTopicId++
		topicList[i].CreateTime = time.Now()
		s, err := json.Marshal(topicList[i])
		if err == nil {
			topicTmpList = append(topicTmpList, string(s))
		}
	}
	str := strings.Join(topicTmpList, "\n")
	if len(topicTmpList) == 1 {
		str += "\n"
	}
	// 写文件
	if err := appendStringToTopic("data/", str); err != nil {
		return err
	}
	// 写内存
	for i := range topicList {
		topicIndexMap[topicList[i].Id] = &topicList[i]
	}
	return nil
}

func appendStringToTopic(filePath string, str string) error {
	topicOpen, err := os.OpenFile(filePath+"topic", os.O_APPEND, 0)
	defer topicOpen.Close()
	if err != nil {
		log.Panic(err)
	}
	topicWriter := bufio.NewWriter(topicOpen)
	defer topicWriter.Flush()
	if _, err := topicWriter.WriteString(str); err != nil {
		return err
	}
	return nil
}
