package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/trace"
)

type BaseResponseSchema struct {
	Success  bool        `json:"success"`
	TraceId  string      `json:"traceId"`
	Error    interface{} `json:"error"`
	Metadata interface{} `json:"meta"`
	Result   interface{} `json:"result"`
}

type BaseErrorResponseSchema struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
}

type BaseErrorValidationSchema struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func restResponse(httpCode int, responseSchema interface{}, metadata interface{}, errorSchema BaseErrorResponseSchema, c *gin.Context) *gin.Context {
	var traceId = trace.SpanFromContext(c.Request.Context()).SpanContext().TraceID().String()
	var isSuccess = true
	var errData interface{}

	if errorSchema.Message != "" {
		isSuccess = false
		errData = errorSchema
	}

	c.JSON(httpCode, BaseResponseSchema{
		Success:  isSuccess,
		TraceId:  traceId,
		Error:    errData,
		Metadata: metadata,
		Result:   &responseSchema,
	})

	return c
}

func HttpSuccess(responseSchema interface{}, metadata interface{}, c *gin.Context) {
	var errorSchema = BaseErrorResponseSchema{}
	restResponse(http.StatusOK, &responseSchema, metadata, errorSchema, c)
}

func HttpErrorResponse(httpCode int, errorSchema BaseErrorResponseSchema, metadata interface{}, c *gin.Context) {
	restResponse(httpCode, nil, metadata, errorSchema, c)
}
