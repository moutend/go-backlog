package types

import (
	"errors"
	"fmt"
)

var (
	ErrorInternal              = errors.New("internal error")
	ErrorLicense               = errors.New("license error")
	ErrorLicenseExpired        = errors.New("license expired error")
	ErrorAccessDenied          = errors.New("access denied error")
	ErrorUnauthorizedOperation = errors.New("unauthorized operation error")
	ErrorNoResource            = errors.New("no resource error")
	ErrorInvalidRequest        = errors.New("invalid request error")
	ErrorSpaceOverCapacity     = errors.New(" space over capacity error")
	ErrorResourceOverflow      = errors.New("resource overflow error")
	ErrorTooLargeFile          = errors.New("too large file error")
	ErrorAuthentication        = errors.New("authentication error")
	ErrorUnexpected            = errors.New("unexpected error")
)

type Error struct {
	Message  string `json:"message"`
	Code     int    `json:"code"`
	MoreInfo string `json:"moreInfo"`
}

type Errors struct {
	Errors []Error `json:"errors"`
}

func convertToError(code int) error {
	switch code {
	case 1:
		return ErrorInternal
	case 2:
		return ErrorLicense
	case 3:
		return ErrorLicenseExpired
	case 4:
		return ErrorAccessDenied
	case 5:
		return ErrorUnauthorizedOperation
	case 6:
		return ErrorNoResource
	case 7:
		return ErrorInvalidRequest
	case 8:
		return ErrorSpaceOverCapacity
	case 9:
		return ErrorResourceOverflow
	case 10:
		return ErrorTooLargeFile
	case 11:
		return ErrorAuthentication
	}

	return ErrorUnexpected
}

// Has returns true if the err matches one of the errors.
func (e Errors) Has(err error) bool {
	for i, _ := range e.Errors {
		if err == convertToError(e.Errors[i].Code) {
			return true
		}
	}

	return false
}

// Error implements error interface.
func (e Errors) Error() string {
	if len(e.Errors) == 0 {
		return ""
	}

	result := fmt.Sprintf(
		"%s: %s (%s)",
		convertToError(e.Errors[0].Code).Error(),
		e.Errors[0].Message,
		e.Errors[0].MoreInfo,
	)

	if len(e.Errors) == 1 {
		return result
	}
	for i := 1; i < len(e.Errors); i++ {
		result = fmt.Sprintf(
			"%s; %s: %s (%s)",
			result,
			convertToError(e.Errors[i].Code).Error(),
			e.Errors[i].Message,
			e.Errors[i].MoreInfo,
		)
	}

	return result
}
