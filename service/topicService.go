package service

import (
	"github.com/gin-gonic/gin"
)

func QueryPageInfo(c *gin.Context) (info *PageInfo, err error) {
	queryPageInfoFlow := &QueryPageInfoFlow{
		c: c,
	}
	return queryPageInfoFlow.Do()
}
