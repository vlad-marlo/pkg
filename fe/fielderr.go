package fe

import (
	"net/http"
)

// Error is custom error that is using in service to deliver error to controllers with prepared statuses and log fields.
type Error[T any] struct {
	// msg is error message that will be returned when Error() is called.
	msg string
	// Data stores http response if it must be returned back to user.
	data any
	// Code must be internal code from this pkg.
	code Code
	// Fields is additional field for zap logger.
	fields []T
	// parent is parent error
	parent error
}

// New creates new error with provided fields.
func New[T any](msg string, data any, code Code, fields ...T) *Error[T] {
	return &Error[T]{msg, data, code, fields, nil}
}

// Error return error message.
func (f *Error[T]) Error() string {
	if f == nil {
		return ""
	}
	return f.msg
}

// CodeHTTP returns http Code that is equal to custom one.
func (f *Error[T]) CodeHTTP() int {
	if f == nil {
		return http.StatusInternalServerError
	}
	if c, ok := httpCodes[f.code]; ok {
		return c
	}
	return http.StatusInternalServerError
}

// With create new error object that copies error fields instead of Fields
func (f *Error[T]) With(fields ...T) *Error[T] {
	if f == nil {
		return &Error[T]{fields: fields}
	}
	return &Error[T]{
		msg:    f.msg,
		data:   f.data,
		code:   f.code,
		fields: append(f.fields, fields...),
		parent: f,
	}
}

// WithData create new error object that copies error fields instead of data.
func (f *Error[T]) WithData(data any) *Error[T] {
	if f == nil {
		return &Error[T]{data: data}
	}
	return &Error[T]{
		msg:    f.msg,
		data:   data,
		code:   f.code,
		fields: f.fields,
		parent: f,
	}
}

// Data return data to return to user.
func (f *Error[T]) Data() any {
	if f == nil {
		return nil
	}
	return f.data
}

// Code return internal code.
func (f *Error[T]) Code() Code {
	if f == nil {
		return 0
	}
	return f.code
}

// Fields return zap fields.
func (f *Error[T]) Fields() []T {
	if f == nil {
		return nil
	}
	return f.fields
}

// Unwrap make available to use errors.Is with *Error.
func (f *Error[T]) Unwrap() error {
	if f == nil {
		return nil
	}
	return f.parent
}
