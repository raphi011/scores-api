package scores

import "github.com/pkg/errors"

// ErrorNotFound is returned if an entity was not found
var ErrorNotFound = errors.New("not found")

// ErrorUnauthorized is returned if the requestor of an operation is unauthorized
var ErrorUnauthorized = errors.New("unauthorized")
