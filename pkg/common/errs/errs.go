package errs

import (
	"fmt"

	"github.com/twitchtv/twirp"
	ec "qx/pkg/common/error_code"
)

// RimeErr implement twirp.Error
type RimeErr struct {
	twirpErrorCode twirp.ErrorCode
	DisplayErr     *ec.DisplayError
	LogErr         error
	metaMap        map[string]string
}

// 仅起到类型检查作用
func checkImplement() twirp.Error {
	return &RimeErr{}
}

func (t *RimeErr) ErrorID() string {
	if t == nil || t.DisplayErr == nil {
		return ""
	}

	return t.DisplayErr.ErrorID
}

func (t *RimeErr) Code() twirp.ErrorCode {
	if t == nil || t.twirpErrorCode == "" {
		return twirp.Internal
	}

	return t.twirpErrorCode
}

func (t *RimeErr) Msg() string {
	if t == nil || t.DisplayErr == nil {
		return ""
	}

	return t.DisplayErr.Error()
}

func (t *RimeErr) WithMeta(key string, val string) twirp.Error {
	if t == nil {
		return nil
	}
	if t.metaMap == nil {
		return &RimeErr{
			twirpErrorCode: t.twirpErrorCode,
			DisplayErr:     t.DisplayErr,
			LogErr:         t.LogErr,
			metaMap:        map[string]string{key: val},
		}
	}

	newMetaMap := make(map[string]string)
	for k, v := range t.metaMap {
		newMetaMap[k] = v
	}
	newMetaMap[key] = val

	return &RimeErr{
		twirpErrorCode: t.twirpErrorCode,
		DisplayErr:     t.DisplayErr,
		LogErr:         t.LogErr,
		metaMap:        newMetaMap,
	}
}

func (t *RimeErr) Meta(key string) string {
	if t == nil || t.metaMap == nil {
		return ""
	}

	return t.metaMap[key]
}

func (t *RimeErr) MetaMap() map[string]string {
	if t == nil {
		return nil
	}

	return t.metaMap
}

func (t *RimeErr) Error() string {
	if t == nil || t.LogErr == nil {
		return ""
	}

	return fmt.Sprintf("%+v", t.LogErr)
}
