package module

import "net/http"

//数据请求的类型
type Request struct {
	httpReq *http.Request
	depth uint32
}

//用于创建一个新的请求实例
func NewRequest(httpReq *http.Request,depth uint32) *Request {
	return &Request{httpReq:httpReq,depth:depth}
}

//用于获取http的请求
func (req *Request) HTTPReq() *http.Request {
	return req.httpReq
}

//用于返回一个请求的深度
func (req *Request) Depth() uint32 {
	return req.depth
}

//数据响应的类型
type Response struct {
	//响应
	httpResp *http.Response
	//深度
	depth uint32
}

//同样也要新建个方法来新建
func NewResponse(httpResp *http.Response,depth uint32) *Response {
	return &Response{httpResp:httpResp,depth:depth}
}

//用于获取到HTTP的响应
func (resp *Response) HTTPResp() *http.Response {
	return resp.httpResp
}

//用于获取响应的深度
func (resp *Response) Depth() uint32 {
	return resp.depth
}


//条目的类型
type Item map[string]interface{}

//数据的接口类型
type Data interface {
	Valid() bool
}

func (req *Request) Valid() bool {
	return req.httpReq != nil && req.httpReq.URL != nil
}

func (req *Response) Valid() bool {
	return resp.httpResp != nil && resp.httpResp.Body != nil
}

func (item Item) Valid() bool {
	return item != nil
}



