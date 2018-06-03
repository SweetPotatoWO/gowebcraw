package errors

import (
	"strings"
	"bytes"
	"fmt"
)

type ErrorType string


//一个抓取的接口
type CrawlerError interface {
	Type() ErrorType
	Error() string
}

//爬虫错误的实现类型
type myCrawlerError struct {
	errType ErrorType
	errMsg string
	fullErrMsg string
}

const (
	//下载器错误
	ERROR_TYPE_DOWNLOADER ErrorType = "downloader error"
	//分析器错误
	ERROR_TYPE_ANALYZER ErrorType = "analyzer error"
	//条目处理管道错误
	ERROR_TYPE_PIPELINE ErrorType = "pipeline error"
	//调度器错误
	ERROR_TYPE_SCHENULER ErrorType = "scheduler error"
)

//创建一个新的错误类型
func NewCrawlerError(errorType ErrorType,errMsg string) CrawlerError {
	return &myCrawlerError{
		errType:errorType,
		errMsg:strings.TrimSpace(errMsg),
	}
}

//实现接口
func (e *myCrawlerError) Type() ErrorType  {
	return  e.errType
}

func (e *myCrawlerError) Error() string  {
	if e.fullErrMsg == "" {
		e.genFullErrMsg()
	}
	return e.fullErrMsg
}
// 结构体自带的方法
func (e *myCrawlerError) genFullErrMsg() {
	var buffer bytes.Buffer
	buffer.WriteString("crawler error: ")
	if e.errType != "" {
		buffer.WriteString(string(e.errType))
		buffer.WriteString(":")
	}
	buffer.WriteString(e.errMsg)
	e.fullErrMsg = fmt.Sprint("%s",buffer.String())
	return
}









