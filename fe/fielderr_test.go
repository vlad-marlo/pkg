package fe

import (
	"crypto/rand"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func randomString() string {
	b := make([]byte, 30)
	if _, err := rand.Read(b); err != nil {
		return "non_random_string"
	}
	return string(b)
}

var (
	someData  = ""
	testError = &Error{
		msg:    "some msg",
		data:   nil,
		code:   CodeInternalServerError,
		fields: []zap.Field{},
		parent: nil,
	}
)

func TestNewError(t *testing.T) {
	msg, data := "err", map[string]any{"xd": "ds"}
	err1 := New(msg, data, CodeInternalServerError)
	err2 := New(msg, data, CodeInternalServerError)
	assert.NotErrorIs(t, err1, err2)
	assert.Equal(t, err1, err2)
	wrappedErr1 := fmt.Errorf("err: %w", err1)
	assert.ErrorIs(t, wrappedErr1, err1)
}

//goland:noinspection GoNilness
func TestFieldError_Error(t *testing.T) {
	for i := 0; i < 1000; i++ {
		msg := randomString()
		var err error = &Error{
			data: nil,
			msg:  msg,
		}
		assert.Equal(t, err.Error(), msg)
	}
	var err *Error
	assert.Equal(t, "", err.Error())
}

func TestFieldError_Fields(t *testing.T) {
	fields := map[string]any{}
	var err = &Error{data: fields}
	assert.Equal(t, fields, err.Data())
	assert.Equal(t, nil, (*Error)(nil).Data())
}

func TestError_CodeHTTP(t *testing.T) {
	for k, v := range httpCodes {
		assert.Equal(t, v, (&Error{code: k}).CodeHTTP())
	}
	assert.Equal(t, http.StatusInternalServerError, (&Error{code: 123}).CodeHTTP())
	assert.Equal(t, http.StatusInternalServerError, (*Error)(nil).CodeHTTP())
}

func TestErrorIs(t *testing.T) {
	msg := "msg"
	data := map[string]any{
		"xd":  nil,
		"bad": 12331,
	}
	code := CodeInternalServerError

	err := New(msg, data, code)
	newErr := err.With(zap.String("string", "sdf"), zap.Error(nil))
	assert.ErrorIs(t, error(newErr), error(err))
	assert.NotEqual(t, error(err), error(newErr))
	assert.Equal(t, (error)(nil), (*Error)(nil).Unwrap())
}

//goland:noinspection GoNilness
func TestError_Fields(t *testing.T) {
	fields := []zap.Field{
		zap.String("xd", "xd"),
	}
	err := &Error{fields: fields}
	assert.Equal(t, err.Fields(), err.fields)
	err = nil
	assert.Equal(t, ([]zap.Field)(nil), err.Fields())
}

func TestError_With(t *testing.T) {
	fields := []zap.Field{zap.String("", "")}
	err := (*Error)(nil).With(fields...)
	assert.Equal(t, &Error{fields: fields}, err)
}

func TestError_Code(t *testing.T) {
	err := &Error{code: CodeInternalServerError}
	assert.Equal(t, Code(0), (*Error)(nil).Code())
	assert.Equal(t, err.code, err.Code())
}

func TestError_WithData(t *testing.T) {
	tt := []struct {
		name string
		err  *Error
		want *Error
		data any
	}{
		{"nil error", nil, &Error{data: someData}, someData},
		{"non nil error", testError, &Error{
			msg:    testError.msg,
			data:   someData,
			code:   testError.code,
			fields: testError.fields,
			parent: testError,
		}, someData},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.err.WithData(tc.data)
			assert.Equal(t, tc.want, got)
		})
	}
}
