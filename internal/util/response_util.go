package util

import (
	"github.com/gin-gonic/gin"
	"graduation_system_api/internal/domain"
	"net/http"
)

//构建返回体
func newOpenAPIResponse(ctx *gin.Context) domain.OpenAPIResponse {
	kind := ctx.Param("kind")
	action := ctx.Param("action")
	response := domain.OpenAPIResponse{
		ResponseMetaData: domain.OpenAPIRespMetaData{
			Kind:   kind,
			Action: action,
		},
		Result: struct {
		}{},
	}
	return response
}

//构建成功请求的返回信息
func BuildSuccessResp(ctx *gin.Context, data interface{}) {
	resp := newOpenAPIResponse(ctx)
	if data != nil {
		resp.Result = data
	}
	ctx.JSON(http.StatusOK, resp)
}

//构建错误返回信息
func BuildFailedResp(ctx *gin.Context, code int, err error) {
	resp := newOpenAPIResponse(ctx)
	resp.ResponseMetaData.Error = &domain.OpenAPIError{
		Code:    code,
		Message: err.Error(),
	}
	//通过Error message以及code 来判断是否有错误
	ctx.JSON(http.StatusOK, resp)
}
