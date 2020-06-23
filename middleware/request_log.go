package middleware

import (
	"bytes"
	"bastion/utils"
	"github.com/georgehao/log"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"strings"
	"time"
)

func RequestInLog(c *gin.Context) {
	traceContext := utils.NewTrace()
	if traceId := c.Request.Header.Get("com-header-rid"); traceId != "" {
		traceContext.TraceId = traceId
	}
	if spanId := c.Request.Header.Get("com-header-spanid"); spanId != "" {
		traceContext.SpanId = spanId
	}

	c.Set("startExecTime", time.Now())
	c.Set("trace", traceContext)

	bodyBytes, _ := ioutil.ReadAll(c.Request.Body)
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes)) // Write body back

	log.RequestLogInfow("_request_in", "",
		"uri", c.Request.RequestURI,
		"method", c.Request.Method,
		"args", c.Request.PostForm,
		"body", string(bodyBytes),
		"from", c.ClientIP(),
		"traceId", traceContext.TraceId,
		"spanId", traceContext.SpanId,
	)
}

func RequestOutLog(c *gin.Context) {
	// after request
	endExecTime := time.Now()
	response, _ := c.Get("response")
	st, _ := c.Get("startExecTime")

	startExecTime, _ := st.(time.Time)

	trace, _ := c.Get("trace")
	traceContext, _ := trace.(*utils.TraceContext)
	SpanId := ""
	if traceContext != nil {
		SpanId = traceContext.SpanId
	}

	log.RequestLogInfow("_request_out", "",
		"uri", c.Request.RequestURI,
		"method", c.Request.Method,
		"args", c.Request.PostForm,
		"from", c.ClientIP(),
		"proc_time", endExecTime.Sub(startExecTime).Seconds(),
		"response", response,
		"spanId", SpanId,
	)
}

func RequestLog() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 过滤掉监控检查的请求日志
		if strings.Contains(c.Request.URL.String(), "metrics") {
			c.Next()
			return
		}
		
		RequestInLog(c)
		defer RequestOutLog(c)
		c.Next()
	}
}
