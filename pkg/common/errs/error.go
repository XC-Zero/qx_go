package errs

import "errors"

func ErrListToError(errList []error) error {
	if len(errList) > 0 {
		errMsg := ""
		for _, err := range errList {
			if err != nil {
				errMsg += err.Error() + ";\n"
			}
		}
		return errors.New(errMsg)
	}

	return nil
}
