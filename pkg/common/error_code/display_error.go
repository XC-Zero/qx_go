package ec

import (
	"encoding/json"
)

// DisplayError ...
type DisplayError struct {
	Code    DisplayErrorCode
	Message string
	ErrorID string
}

func (d DisplayError) Error() string {
	message, _ := json.Marshal(d)
	return string(message)
}

// DisplayErrorCode ...
type DisplayErrorCode string

// const 命名格式 ErrCodeXXXError
const (
	// 一切系统内部的错误
	// 比如 sql 错误，Marshal 错误 等
	ErrCodeInternalServerError DisplayErrorCode = "201000"
	// Excel 格式与模板不匹配
	ErrCodeExcelFormatError DisplayErrorCode = "201001"
	// Excel 匹配条件不允许为空
	ErrCodeExcelIsEmptyError DisplayErrorCode = "201002"
	// Excel 记录数超过允许最大记录数
	ErrCodeExcelRowNumOutRangeError    DisplayErrorCode = "201003"
	ErrCodeRelevantEntitiesImportError DisplayErrorCode = "201004"
	ErrCodeEntitiesImportError         DisplayErrorCode = "201005"

	// 前端请求参数错误
	ErrCodeBadParameterError DisplayErrorCode = "201006"
	// 用户未授权
	ErrCodeUnauthorizedError DisplayErrorCode = "201007"
	// 访问次数达到上限
	ErrCodeFrequencyLimitError DisplayErrorCode = "201008"
	// 资源受限
	ErrCodeResourceLimitError DisplayErrorCode = "201009"
	// 资源未找到
	ErrCodeResourceNotFoundError DisplayErrorCode = "201010"
)

// GetMessage 根据 DisplayErrorCode 获取具体展示消息
func (d DisplayErrorCode) GetMessage() string {
	switch d {
	case ErrCodeInternalServerError:
		return "系统内部异常"
	case ErrCodeExcelFormatError:
		return "Excel 格式与模板不匹配"
	case ErrCodeExcelIsEmptyError:
		return "Excel 匹配条件不允许为空"
	case ErrCodeExcelRowNumOutRangeError:
		return "Excel 记录数超过允许最大记录数 %d"
	case ErrCodeEntitiesImportError:
		return "我们将为您人工导入，请耐心等待"
	case ErrCodeBadParameterError:
		return "错误的请求参数"
	case ErrCodeUnauthorizedError:
		return "未授权"
	case ErrCodeFrequencyLimitError:
		return "访问次数已达到上限"
	case ErrCodeResourceLimitError:
		return "资源受限"
	case ErrCodeResourceNotFoundError:
		return "资源未找到"
	}

	return "未知错误"
}

// GetError 根据 DisplayErrorCode 获取整个 DisplayError
func (d DisplayErrorCode) GetError() DisplayError {
	return DisplayError{
		Code:    d,
		Message: d.GetMessage(),
	}
}
