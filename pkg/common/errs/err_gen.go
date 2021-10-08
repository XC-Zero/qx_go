package errs

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/twitchtv/twirp"
	ec "qx/pkg/common/error_code"
)

type (
	errFunc      = func(errMsg string) error
	errfFunc     = func(errMsg string, v ...interface{}) error
	withErrFunc  = func(err error) error
	withErrfFunc = func(err error, errMsg string, v ...interface{}) error
)

func generateErrorFunc(displayError *ec.DisplayError) errFunc {
	return func(errMsg string) error {
		return &RimeErr{
			DisplayErr: &ec.DisplayError{
				Code:    displayError.Code,
				Message: displayError.Message,
				ErrorID: uuid.New().String(),
			},
			LogErr: errors.New(errMsg),
		}
	}
}

func generateErrorfFunc(displayError *ec.DisplayError) errfFunc {
	return func(errMsg string, v ...interface{}) error {
		return &RimeErr{
			DisplayErr: &ec.DisplayError{
				Code:    displayError.Code,
				Message: displayError.Message,
				ErrorID: uuid.New().String(),
			},
			LogErr: errors.Errorf(errMsg, v...),
		}
	}
}

func generateWithErrorFunc(displayError *ec.DisplayError) withErrFunc {
	return func(err error) error {
		if err == nil {
			return nil
		}

		switch e := err.(type) {
		case *RimeErr:
			return err
		case twirp.Error:
			if e != nil && e.Msg() != "" {
				var d ec.DisplayError
				err := json.Unmarshal([]byte(e.Msg()), &d)
				if err == nil && d.ErrorID != "" {
					return &RimeErr{
						DisplayErr: &ec.DisplayError{
							Code:    displayError.Code,
							Message: displayError.Message,
							ErrorID: d.ErrorID,
						},
						LogErr: errors.WithStack(err),
					}
				}
			}
		}

		return &RimeErr{
			DisplayErr: &ec.DisplayError{
				Code:    displayError.Code,
				Message: displayError.Message,
				ErrorID: uuid.New().String(),
			},
			LogErr: errors.WithStack(err),
		}
	}
}

func generateWithErrorfFunc(displayError *ec.DisplayError) withErrfFunc {
	return func(err error, errMsg string, v ...interface{}) error {
		if err == nil {
			return nil
		}

		switch e := err.(type) {
		case *RimeErr:
			return err
		case twirp.Error:
			if e != nil && e.Msg() != "" {
				var d ec.DisplayError
				err := json.Unmarshal([]byte(e.Msg()), &d)
				if err == nil && d.ErrorID != "" {
					return &RimeErr{
						DisplayErr: &ec.DisplayError{
							Code:    displayError.Code,
							Message: displayError.Message,
							ErrorID: d.ErrorID,
						},
						LogErr: errors.Wrapf(err, errMsg, v...),
					}
				}
			}

		}
		return &RimeErr{
			DisplayErr: &ec.DisplayError{
				Code:    displayError.Code,
				Message: displayError.Message,
				ErrorID: uuid.New().String(),
			},
			LogErr: errors.Wrapf(err, errMsg, v...),
		}
	}
}

func generateErrorFunctions(displayError ec.DisplayError) (f1 errFunc, f2 errfFunc, f3 withErrFunc, f4 withErrfFunc) {
	f1 = generateErrorFunc(&displayError)
	f2 = generateErrorfFunc(&displayError)
	f3 = generateWithErrorFunc(&displayError)
	f4 = generateWithErrorfFunc(&displayError)

	return
}
