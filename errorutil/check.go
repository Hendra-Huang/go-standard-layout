package errorutil

import "github.com/tokopedia/vehicle-insurance/backend/log"

// CheckAndLog checks the error and log if the error exist
func CheckAndLog(err error) {
	if err != nil {
		log.Errors(err)
	}
}

// CheckWithErrorHandler checks the error and run errorHandler if error exist
func CheckWithErrorHandler(err error, errorHandler func(error)) {
	if err != nil {
		errorHandler(err)
	}
}

// IsError returns true if there is an error
func IsError(err error) bool {
	if err != nil {
		return true
	}
	return false
}
