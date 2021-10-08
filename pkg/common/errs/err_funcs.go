package errs

import (
	"github.com/pkg/errors"
	ec "qx/pkg/common/error_code"
)

var (
	InternalError,
	InternalErrorf,
	WithInternalError,
	WithInternalErrorf = generateErrorFunctions(ec.ErrCodeInternalServerError.GetError())

	BadParameterError,
	BadParameterErrorf,
	WithBadParameterError,
	WithBadParameterErrorf = generateErrorFunctions(ec.ErrCodeBadParameterError.GetError())

	UnauthorizedError,
	UnauthorizedErrorf,
	WithUnauthorizedError,
	WithUnauthorizedErrorf = generateErrorFunctions(ec.ErrCodeUnauthorizedError.GetError())

	FrequencyLimitError,
	FrequencyLimitErrorf,
	WithFrequencyLimitError,
	WithFrequencyLimitErrorf = generateErrorFunctions(ec.ErrCodeFrequencyLimitError.GetError())

	ResourceLimitError,
	ResourceLimitErrorf,
	WithResourceLimitError,
	WithResourceLimitErrorf = generateErrorFunctions(ec.ErrCodeResourceLimitError.GetError())

	ExcelFormatError,
	ExcelFormatErrorf,
	WithExcelFormatError,
	WithExcelFormatErrorf = generateErrorFunctions(ec.ErrCodeExcelFormatError.GetError())

	ExcelIsEmptyError,
	ExcelIsEmptyErrorf,
	WithExcelIsEmptyError,
	WithExcelIsEmptyErrorf = generateErrorFunctions(ec.ErrCodeExcelIsEmptyError.GetError())

	ExcelRowNumOutRangeError,
	ExcelRowNumOutRangeErrorf,
	WithExcelRowNumOutRangeError,
	WithExcelRowNumOutRangeErrorf = generateErrorFunctions(ec.ErrCodeExcelRowNumOutRangeError.GetError())

	EntitiesImportError,
	EntitiesImportErrorf,
	WithEntitiesImportError,
	WithEntitiesImportErrorf = generateErrorFunctions(ec.ErrCodeEntitiesImportError.GetError())

	ResourceNotFoundError,
	ResourceNotFoundErrorf,
	WithResourceNotFoundError,
	WithResourceNotFoundErrorf = generateErrorFunctions(ec.ErrCodeResourceNotFoundError.GetError())
)

var Is = errors.Is
