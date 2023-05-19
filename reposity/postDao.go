package reposity

import "sync"

type PostDao struct {
}

var (
	postDao  *PostDao
	postOnce sync.Once
)

func PostDaoInstance() *PostDao {
	postOnce.Do(
		func() {
			postDao = &PostDao{}
		})
	return postDao
}

func (*PostDao) QueryPostListByTopicId(id int64) []*Post {
	return postIndexMap[id]
}
