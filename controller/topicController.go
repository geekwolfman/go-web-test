package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-web-test/service"
)

func QueryPageInfo(c *gin.Context) *PageData {
	pageInfo, err := service.QueryPageInfo(c)
	if err != nil {
		return &PageData{
			Code: 1,
			Msg:  fmt.Sprintf("query page info error, error is %v", err),
		}
	}
	return &PageData{
		Code: 200,
		Msg:  "success",
		Data: pageInfo,
	}
}

func PostPage(c *gin.Context) *PageData {
	if err := service.PostPage(c); err != nil {
		return &PageData{
			Code: 1,
			Msg:  fmt.Sprintf("post page error, error is %v", err),
		}
	}
	return &PageData{
		Code: 200,
		Msg:  "success",
	}
}
