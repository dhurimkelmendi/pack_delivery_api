package api

// ErrorContext is the name of the context for the API error.
type ErrorContext string

// ErrorContextFn is a function that decorate the given errors in the error context of the function.
type ErrorContextFn func(*ResponseError, error) error

// Authentication error contexts
const (
	CtxAuthentication ErrorContext = "ctxAuthentication"
)

// Product error contexts
const (
	CtxCreateProduct   ErrorContext = "ctxCreateProduct"
	CtxChangePackSizes ErrorContext = "ctxChangePackSizes"
)

// Serializer error contexts
const ()
