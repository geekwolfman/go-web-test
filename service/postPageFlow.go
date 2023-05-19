package service

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go-web-test/reposity"
)

type PostPageFlow struct {
	c     *gin.Context
	topic reposity.Topic
}

func (f *PostPageFlow) Do() error {
	if err := f.check(); err != nil {
		return err
	}
	if err := f.prepare(); err != nil {
		return err
	}
	if err := f.pack(); err != nil {
		return err
	}
	return nil
}

func (f *PostPageFlow) check() error {
	return nil
}

func (f *PostPageFlow) prepare() error {
	raw, err := f.c.GetRawData()
	if err != nil {
		return err
	}
	var topic reposity.Topic
	if err := json.Unmarshal(raw, &topic); err != nil {
		return err
	}
	f.topic = topic
	return nil
}

func (f *PostPageFlow) pack() error {
	return reposity.TopicDaoInstance().AppendTopicList([]reposity.Topic{
		f.topic,
	})
}
