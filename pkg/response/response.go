package response

import (
	"bastion/pkg/errno"
	"bastion/utils"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

type Response struct {
	Status    string      `json:"status"`
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	Error     interface{} `json:"error"`
	Result    interface{} `json:"result"`
	TraceId   interface{} `json:"_trace_id"`
	TimeStamp int64       `json:"_ts"`
}

type PageData struct {
	Page     int         `json:"page"`
	Pagesize int         `json:"pagesize"`
	Total    int         `json:"total"`
	Rows     interface{} `json:"rows"`
}

func getTraceId(c *gin.Context) string {
	trace, _ := c.Get("trace")
	traceContext, _ := trace.(*utils.TraceContext)
	traceId := ""
	if traceContext != nil {
		traceId = traceContext.TraceId
	}
	return traceId
}

// 正常
func Success(c *gin.Context, result interface{}, message ...string) {
	traceId := getTraceId(c)

	// 空结构 防止前端不严谨 导致前端同学错误
	if result == nil {
		result = struct{}{}
	}

	// 没有消息 就给个ok
	if len(message) == 0 {
		message = append(message, "ok")
	}

	resp := &Response{
		Status:    "ok",
		Code:      errno.OK.Code,
		Message:   strings.Join(message, ","),
		Result:    result,
		TraceId:   traceId,
		TimeStamp: time.Now().UnixNano(),
	}

	c.JSON(200, resp)
	response, _ := json.Marshal(resp)
	c.Set("response", string(response))
}

// 失败
func Fail(c *gin.Context, errno *errno.Errno, e error) {
	traceId := getTraceId(c)

	resp := &Response{
		Status:    "fail",
		Code:      errno.Code,
		Message:   errno.Message,
		Error:     e,
		Result:    struct{}{},
		TraceId:   traceId,
		TimeStamp: time.Now().UnixNano(),
	}

	c.JSON(200, resp)
	response, _ := json.Marshal(resp)
	c.Set("response", string(response))
	_ = c.AbortWithError(200, e)
}

// 非业务错误 会影响 http Status Code
func Error(c *gin.Context, httpStatusCode int, e error) {
	traceId := getTraceId(c)

	resp := &Response{
		Status:    "error",
		Code:      -1,
		Message:   e.Error(),
		Error:     e,
		Result:    struct{}{},
		TraceId:   traceId,
		TimeStamp: time.Now().UnixNano(),
	}

	c.JSON(httpStatusCode, resp)
	response, _ := json.Marshal(resp)
	c.Set("response", string(response))
	_ = c.AbortWithError(httpStatusCode, e)
}
