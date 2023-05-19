package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-web-test/reposity"
	"strconv"
	"sync"
)

type QueryPageInfoFlow struct {
	topicId  int64
	c        *gin.Context
	pageInfo *PageInfo
	topic    *reposity.Topic
	postList []*reposity.Post
}

func (f *QueryPageInfoFlow) Do() (info *PageInfo, err error) {
	if err := f.checkParam(); err != nil {
		return nil, err
	}
	if err := f.prepareInfo(); err != nil {
		return nil, err
	}
	if err := f.packPageInfo(); err != nil {
		return nil, err
	}
	return f.pageInfo, nil
}

func (f *QueryPageInfoFlow) checkParam() error {
	topicId, err := strconv.ParseInt(f.c.Param("id"), 10, 64)
	if err != nil {
		return err
	}
	f.topicId = topicId
	if f.topicId < 0 {
		return fmt.Errorf("invalid topic id")
	}
	return nil
}

func (f *QueryPageInfoFlow) prepareInfo() error {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		f.topic = reposity.TopicDaoInstance().QueryTopicById(f.topicId)
		wg.Done()
	}()
	go func() {
		f.postList = reposity.PostDaoInstance().QueryPostListByTopicId(f.topicId)
		fmt.Println(f.postList)
		wg.Done()
	}()
	wg.Wait()
	return nil
}

func (f *QueryPageInfoFlow) packPageInfo() error {
	f.pageInfo = &PageInfo{
		Topic:    f.topic,
		PostList: f.postList,
	}
	return nil
}
