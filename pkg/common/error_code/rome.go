package ec

import "encoding/json"

// RomeError ...
type RomeError struct {
	Code    RomeErrorCode
	Message string
}

// Error ...
func (e RomeError) Error() string {
	message, _ := json.Marshal(e)
	return string(message)
}

// RomeErrorCode ...
type RomeErrorCode string

// RomeErrorCode ...
const (
	// 账号未登录
	ErrCodeNotLoggedIn RomeErrorCode = "111000"
	// 账号在其他地方登录
	ErrCodeAccountLoginInOtherPlaces RomeErrorCode = "111001"
)

// GetMessage ...
func (e RomeErrorCode) GetMessage() string {
	switch e {
	case ErrCodeNotLoggedIn:
		return "未登录"
	case ErrCodeAccountLoginInOtherPlaces:
		return "账号在其他地方登录"
	}

	return "未知错误"
}

// GetError ...
func (e RomeErrorCode) GetError() RomeError {
	return RomeError{
		Code:    e,
		Message: e.GetMessage(),
	}
}
