package errors

import "fmt"

var (
	INVALID_HEADER     = fmt.Errorf("Invalid header.")
	UNAUTHORIZED       = fmt.Errorf("Unauthorized.")
	VALIDATE_ERROR     = fmt.Errorf("Validate field error.")
	APPLICATION_LOCKED = fmt.Errorf("Application has been locked.")
	CREDIT_INSUFFIENCT = fmt.Errorf("Credit has insuffienct.")
	FILE_REQUIRED      = fmt.Errorf("File required.")
	FILE_TOO_LARGE     = fmt.Errorf("File too large.")
	FILE_NOT_SUPPORTED = fmt.Errorf("File not support (PNG/JPEG only).")
	SLIP_NOT_EXIST     = fmt.Errorf("Slip not exist.")
	SERVER_ERROR       = fmt.Errorf("Server error.")
	SERVICE_ERROR      = fmt.Errorf("Service error.")
	SERVICE_TIMEOUT    = fmt.Errorf("Service timeout.")
	UNKNOWN_ERROR      = fmt.Errorf("Unknown error.")
)
