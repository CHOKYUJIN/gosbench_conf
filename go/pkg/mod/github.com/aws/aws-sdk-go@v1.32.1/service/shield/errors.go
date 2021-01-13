// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package shield

import (
	"github.com/aws/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeAccessDeniedException for service response error code
	// "AccessDeniedException".
	//
	// Exception that indicates the specified AttackId does not exist, or the requester
	// does not have the appropriate permissions to access the AttackId.
	ErrCodeAccessDeniedException = "AccessDeniedException"

	// ErrCodeAccessDeniedForDependencyException for service response error code
	// "AccessDeniedForDependencyException".
	//
	// In order to grant the necessary access to the DDoS Response Team (DRT), the
	// user submitting the request must have the iam:PassRole permission. This error
	// indicates the user did not have the appropriate permissions. For more information,
	// see Granting a User Permissions to Pass a Role to an AWS Service (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use_passrole.html).
	ErrCodeAccessDeniedForDependencyException = "AccessDeniedForDependencyException"

	// ErrCodeInternalErrorException for service response error code
	// "InternalErrorException".
	//
	// Exception that indicates that a problem occurred with the service infrastructure.
	// You can retry the request.
	ErrCodeInternalErrorException = "InternalErrorException"

	// ErrCodeInvalidOperationException for service response error code
	// "InvalidOperationException".
	//
	// Exception that indicates that the operation would not cause any change to
	// occur.
	ErrCodeInvalidOperationException = "InvalidOperationException"

	// ErrCodeInvalidPaginationTokenException for service response error code
	// "InvalidPaginationTokenException".
	//
	// Exception that indicates that the NextToken specified in the request is invalid.
	// Submit the request using the NextToken value that was returned in the response.
	ErrCodeInvalidPaginationTokenException = "InvalidPaginationTokenException"

	// ErrCodeInvalidParameterException for service response error code
	// "InvalidParameterException".
	//
	// Exception that indicates that the parameters passed to the API are invalid.
	ErrCodeInvalidParameterException = "InvalidParameterException"

	// ErrCodeInvalidResourceException for service response error code
	// "InvalidResourceException".
	//
	// Exception that indicates that the resource is invalid. You might not have
	// access to the resource, or the resource might not exist.
	ErrCodeInvalidResourceException = "InvalidResourceException"

	// ErrCodeLimitsExceededException for service response error code
	// "LimitsExceededException".
	//
	// Exception that indicates that the operation would exceed a limit.
	//
	// Type is the type of limit that would be exceeded.
	//
	// Limit is the threshold that would be exceeded.
	ErrCodeLimitsExceededException = "LimitsExceededException"

	// ErrCodeLockedSubscriptionException for service response error code
	// "LockedSubscriptionException".
	//
	// You are trying to update a subscription that has not yet completed the 1-year
	// commitment. You can change the AutoRenew parameter during the last 30 days
	// of your subscription. This exception indicates that you are attempting to
	// change AutoRenew prior to that period.
	ErrCodeLockedSubscriptionException = "LockedSubscriptionException"

	// ErrCodeNoAssociatedRoleException for service response error code
	// "NoAssociatedRoleException".
	//
	// The ARN of the role that you specifed does not exist.
	ErrCodeNoAssociatedRoleException = "NoAssociatedRoleException"

	// ErrCodeOptimisticLockException for service response error code
	// "OptimisticLockException".
	//
	// Exception that indicates that the resource state has been modified by another
	// client. Retrieve the resource and then retry your request.
	ErrCodeOptimisticLockException = "OptimisticLockException"

	// ErrCodeResourceAlreadyExistsException for service response error code
	// "ResourceAlreadyExistsException".
	//
	// Exception indicating the specified resource already exists.
	ErrCodeResourceAlreadyExistsException = "ResourceAlreadyExistsException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// Exception indicating the specified resource does not exist.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"AccessDeniedException":              newErrorAccessDeniedException,
	"AccessDeniedForDependencyException": newErrorAccessDeniedForDependencyException,
	"InternalErrorException":             newErrorInternalErrorException,
	"InvalidOperationException":          newErrorInvalidOperationException,
	"InvalidPaginationTokenException":    newErrorInvalidPaginationTokenException,
	"InvalidParameterException":          newErrorInvalidParameterException,
	"InvalidResourceException":           newErrorInvalidResourceException,
	"LimitsExceededException":            newErrorLimitsExceededException,
	"LockedSubscriptionException":        newErrorLockedSubscriptionException,
	"NoAssociatedRoleException":          newErrorNoAssociatedRoleException,
	"OptimisticLockException":            newErrorOptimisticLockException,
	"ResourceAlreadyExistsException":     newErrorResourceAlreadyExistsException,
	"ResourceNotFoundException":          newErrorResourceNotFoundException,
}
