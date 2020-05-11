package scores

import "github.com/pkg/errors"

// ErrNotFound is returned if an entity was not found.
var ErrNotFound = errors.New("not found")

// ErrorUnauthorized is returned if the requestor of an operation is unauthorized.
var ErrorUnauthorized = errors.New("unauthorized")

// ErrorValidation is returned if the request validation has failed.
var ErrorValidation = errors.New("validation")
