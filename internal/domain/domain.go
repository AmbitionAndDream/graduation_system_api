package domain

//所有外部数据获取，可以自己创建实例

//OpenAPI 返回格式
type OpenAPIResponse struct {
	ResponseMetaData OpenAPIRespMetaData `json:"responseMetaData"`
	Result           interface{}         `json:"result,omitempty"`
}

type OpenAPIRespMetaData struct {
	Kind   string       `json:"kind,omitempty"`
	Action string       `json:"action,omitempty"`
	Error  *OpenAPIError `json:"error,omitempty"`
}

type OpenAPIError struct {
	Code    int `json:"code"`
	Message string `json:"message"`
}

type TokenRole struct {
	IsAdmin bool
}