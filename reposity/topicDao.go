package reposity

import "sync"

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
	return topicIndexMap[id]
}
