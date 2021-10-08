package ec

import (
	"encoding/json"
)

// VisaError ...
type VisaError struct {
	Code    VisaErrorCode
	Message string
}

func (e VisaError) Error() string {
	message, _ := json.Marshal(e)
	return string(message)
}

// VisaErrorCode ...
type VisaErrorCode string

// GetMessage ...
func (e VisaErrorCode) GetMessage() string {
	switch e {
	case ErrCodeSystemError, ErrCodeDataBaseError, ErrCodeTokenGeneratorError:
		return "系统错误"
	case ErrCodeAccountExistence:
		return "账号已被使用"
	case ErrCodeAccountNotExistence:
		return "账号不存在"
	case ErrCodeAccountFormatError:
		return "账号格式错误"
	case ErrorCodeWrongAccountOrPassword:
		return "密码错误"
	case ErrCodeAccountInputNULL:
		return "账号输入为空"
	case ErrCodeIllegalOperation:
		return "非法操作"
	case ErrorCodeWrongVerificationCode:
		return "验证码错误"
	case ErrorCodeSendingVerificationCodeFailed:
		return "验证码发送失败"
	case ErrorCodeSendingVerificationCodeFrequencyLimit:
		return "手机号超过频率限制"
	case ErrorCodeVerificationCodeExpired:
		return "验证码过期"
	case ErrorCodeAccountNotActivated:
		return "账号未激活"
	case ErrorCodeAccountHasActivated:
		return "账号已激活"
	case ErrorCodeWrongAccountOrInvitationCode:
		return "账号或邀请码错误"
	case ErrorCodeWrongNoPassword:
		return "密码为空"
	case ErrorCodeAccountLocked:
		return "账号被锁定"
	case ErrorCodeEmailExistence:
		return "邮箱已被使用"
	case ErrorCodeEmailFormatError:
		return "邮箱格式有误"
	case ErrorCodeNotReview:
		return "信息审核中"
	case ErrorCodeReviewPassed:
		return "信息已审核过"
	case ErrorCodeReservedEmailNotMatch:
		return "预留邮箱不匹配"
	case ErrorCodeReservedPhoneNotMatch:
		return "预留手机不匹配"
	case ErrorCodeTheNumberOfRequestsOutOfCount:
		return "今日单个资源访问次数已达上限"
	case ErrorCodeTheNumberOfRequestsOutOfTotalCount:
		return "今日资源访问总次数已达上限"
	case ErrorCodeResourceCreateTemplatesLimit:
		return "创建搜索模板已达上限"
	case ErrorCodeResourceRecommendLimit:
		return "查看智能推荐权限不足"
	case ErrorCodeResourceFollowEntitiesCountLimit:
		return "我的关注已达上限"
	case ErrorCodeResourceFilterLimit:
		return "筛选条件权限不足"
	case ErrorCodeResourceTableRowLimit:
		return "查看更多分页数据权限不足"
	case ErrorCodeResourceAccessLimit:
		return "权限不足"
	case ErrorCodeResourceDeleteUserLimit:
		return "删除用户权限不足"
	case ErrorCodeResourceInviteUserLimit:
		return "邀请用户权限不足"
	}

	return "系统异常"
}

// GetError ...
func (e VisaErrorCode) GetError() VisaError {
	return VisaError{
		Code:    e,
		Message: e.GetMessage(),
	}
}

// CustomErrorMessage ...
func (e VisaErrorCode) CustomErrorMessage(message string) VisaError {
	return VisaError{
		Code:    e,
		Message: message,
	}
}

const (
	// ErrCodeSystemError 系统错误
	ErrCodeSystemError VisaErrorCode = "100000"
	// ErrCodeDataBaseError 数据库错误
	ErrCodeDataBaseError VisaErrorCode = "100001"
	// ErrCodeTokenGeneratorError  token 生成器错误
	ErrCodeTokenGeneratorError VisaErrorCode = "100002"

	// ErrCodeIllegalOperation 非法操作 => 无效令牌
	ErrCodeIllegalOperation VisaErrorCode = "101000"

	// ErrCodeAccountExistence 账号存在
	ErrCodeAccountExistence VisaErrorCode = "109100"
	// ErrCodeAccountNotExistence 账号不存在
	ErrCodeAccountNotExistence VisaErrorCode = "109101"
	// ErrCodeAccountFormatError 账号格式错误
	ErrCodeAccountFormatError VisaErrorCode = "109102"
	// ErrCodeAccountInputNULL 账号输入为空
	ErrCodeAccountInputNULL VisaErrorCode = "109103"
	// ErrorCodeWrongAccountOrPassword 账号或密码错误
	ErrorCodeWrongAccountOrPassword VisaErrorCode = "109104"
	// ErrorCodeAccountNotActivated 账号未激活
	ErrorCodeAccountNotActivated VisaErrorCode = "109105"
	// ErrorCodeAccountHasActivated 账号已激活
	ErrorCodeAccountHasActivated VisaErrorCode = "109106"
	// ErrorCodeWrongAccountOrInvitationCode 账号或邀请码错误
	ErrorCodeWrongAccountOrInvitationCode VisaErrorCode = "109107"
	// ErrorCodeWrongNoPassword 密码为空
	ErrorCodeWrongNoPassword VisaErrorCode = "109108"
	// ErrorCodeAccountLocked 账号被锁定
	ErrorCodeAccountLocked VisaErrorCode = "109109"
	// ErrorCodeEmailExistence 邮箱已被使用
	ErrorCodeEmailExistence VisaErrorCode = "109110"
	// ErrorCodeEmailFormatError 邮箱格式有误
	ErrorCodeEmailFormatError VisaErrorCode = "109111"

	// ErrorCodeWrongVerificationCode 验证码错误
	ErrorCodeWrongVerificationCode VisaErrorCode = "109201"
	// ErrorCodeSendingVerificationCodeFailed 验证码发送失败
	ErrorCodeSendingVerificationCodeFailed VisaErrorCode = "109202"
	// ErrorCodeSendingVerificationCodeFrequencyLimit 手机号接收超过频率限制
	ErrorCodeSendingVerificationCodeFrequencyLimit VisaErrorCode = "109203"
	// ErrorCodeVerificationCodeExpired  验证码过期
	ErrorCodeVerificationCodeExpired VisaErrorCode = "109204"
	// ErrorCodeReservedEmailNotMatch 预留邮箱不匹配
	ErrorCodeReservedEmailNotMatch VisaErrorCode = "109205"
	// ErrorCodeReservedPhoneNotMatch 预留手机不匹配
	ErrorCodeReservedPhoneNotMatch VisaErrorCode = "109206"

	// ErrorCodeNotReview 信息审核中
	ErrorCodeNotReview VisaErrorCode = "109301"
	// ErrorCodeReviewPassed 信息已审核
	ErrorCodeReviewPassed VisaErrorCode = "109302"

	// ErrorCodeTheNumberOfRequestsOutOfCount 今日单个资源访问次数已达上限
	ErrorCodeTheNumberOfRequestsOutOfCount VisaErrorCode = "109401"
	// ErrorCodeTheNumberOfRequestsOutOfTotalCount 今日资源访问总次数已达上限
	ErrorCodeTheNumberOfRequestsOutOfTotalCount VisaErrorCode = "109402"
	// ErrorCodeResourceCreateTemplatesLimit 创建搜索模板已达上限
	ErrorCodeResourceCreateTemplatesLimit VisaErrorCode = "109403"
	// ErrorCodeResourceRecommendLimit 查看智能推荐权限不足
	ErrorCodeResourceRecommendLimit VisaErrorCode = "109404"
	// ErrorCodeResourceFollowEntitiesCountLimit 我的关注已达上限
	ErrorCodeResourceFollowEntitiesCountLimit VisaErrorCode = "109405"
	// ErrorCodeResourceFilterLimit 条件筛选权限不足
	ErrorCodeResourceFilterLimit VisaErrorCode = "109406"
	// ErrorCodeResourceTableRowLimit 查看更多分页数据权限不足
	ErrorCodeResourceTableRowLimit VisaErrorCode = "109407"
	// ErrorCodeResourceAccessLimit 权限不足
	ErrorCodeResourceAccessLimit VisaErrorCode = "109408"
	// ErrorCodeResourceDeleteUserLimit 删除用户权限不足
	ErrorCodeResourceDeleteUserLimit VisaErrorCode = "109409"
	// ErrorCodeResourceInviteUserLimit 邀请用户权限不足
	ErrorCodeResourceInviteUserLimit VisaErrorCode = "109410"
)
