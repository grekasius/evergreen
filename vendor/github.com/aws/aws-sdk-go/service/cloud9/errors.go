// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloud9

import (
	"github.com/aws/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeBadRequestException for service response error code
	// "BadRequestException".
	//
	// The target request is invalid.
	ErrCodeBadRequestException = "BadRequestException"

	// ErrCodeConflictException for service response error code
	// "ConflictException".
	//
	// A conflict occurred.
	ErrCodeConflictException = "ConflictException"

	// ErrCodeForbiddenException for service response error code
	// "ForbiddenException".
	//
	// An access permissions issue occurred.
	ErrCodeForbiddenException = "ForbiddenException"

	// ErrCodeInternalServerErrorException for service response error code
	// "InternalServerErrorException".
	//
	// An internal server error occurred.
	ErrCodeInternalServerErrorException = "InternalServerErrorException"

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceededException".
	//
	// A service limit was exceeded.
	ErrCodeLimitExceededException = "LimitExceededException"

	// ErrCodeNotFoundException for service response error code
	// "NotFoundException".
	//
	// The target resource cannot be found.
	ErrCodeNotFoundException = "NotFoundException"

	// ErrCodeTooManyRequestsException for service response error code
	// "TooManyRequestsException".
	//
	// Too many service requests were made over the given time period.
	ErrCodeTooManyRequestsException = "TooManyRequestsException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"BadRequestException":          newErrorBadRequestException,
	"ConflictException":            newErrorConflictException,
	"ForbiddenException":           newErrorForbiddenException,
	"InternalServerErrorException": newErrorInternalServerErrorException,
	"LimitExceededException":       newErrorLimitExceededException,
	"NotFoundException":            newErrorNotFoundException,
	"TooManyRequestsException":     newErrorTooManyRequestsException,
}
