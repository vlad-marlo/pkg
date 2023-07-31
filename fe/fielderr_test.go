package fe

import (
	"crypto/rand"
	"fmt"
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
	testError = &Error[string]{
		msg:    "some msg",
		data:   nil,
		code:   CodeInternalServerError,
		fields: []string{},
		parent: nil,
	}
)

func TestNewError(t *testing.T) {
	msg, data := "err", map[string]any{"xd": "ds"}
	err1 := New[int](msg, data, CodeInternalServerError)
	err2 := New[int](msg, data, CodeInternalServerError)
	assert.NotErrorIs(t, err1, err2)
	assert.Equal(t, err1, err2)
	wrappedErr1 := fmt.Errorf("err: %w", err1)
	assert.ErrorIs(t, wrappedErr1, err1)
}

//goland:noinspection GoNilness
func TestFieldError_Error(t *testing.T) {
	for i := 0; i < 1000; i++ {
		msg := randomString()
		var err error = &Error[int]{
			data: nil,
			msg:  msg,
		}
		assert.Equal(t, err.Error(), msg)
	}
	var err *Error[int]
	assert.Equal(t, "", err.Error())
}

func TestFieldError_Fields(t *testing.T) {
	fields := map[string]any{}
	err := &Error[int]{data: fields}
	assert.Equal(t, fields, err.Data())
	assert.Equal(t, nil, (*Error[int])(nil).Data())
}

func TestError_CodeHTTP(t *testing.T) {
	for k, v := range httpCodes {
		assert.Equal(t, v, (&Error[int]{code: k}).CodeHTTP())
	}
	assert.Equal(t, http.StatusInternalServerError, (&Error[int]{code: 123}).CodeHTTP())
	assert.Equal(t, http.StatusInternalServerError, (*Error[int])(nil).CodeHTTP())
}

func TestErrorIs(t *testing.T) {
	msg := "msg"
	data := map[string]any{
		"xd":  nil,
		"bad": 12331,
	}
	code := CodeInternalServerError

	err := New[string](msg, data, code)
	newErr := err.With("aboba", "other field")
	assert.ErrorIs(t, error(newErr), error(err))
	assert.NotEqual(t, error(err), error(newErr))
	assert.Equal(t, (error)(nil), (*Error[string])(nil).Unwrap())
}

//goland:noinspection GoNilness
func TestError_Fields(t *testing.T) {
	fields := []int{
		1,
		2,
	}
	err := &Error[int]{fields: fields}
	assert.Equal(t, err.Fields(), err.fields)
	err = nil
	assert.Equal(t, ([]int)(nil), err.Fields())
}

func TestError_With(t *testing.T) {
	fields := []string{"one"}
	err := (*Error[string])(nil).With(fields...)
	assert.Equal(t, &Error[string]{fields: fields}, err)
}

func TestError_Code(t *testing.T) {
	err := &Error[string]{code: CodeInternalServerError}
	assert.Equal(t, Code(0), (*Error[string])(nil).Code())
	assert.Equal(t, err.code, err.Code())
}

func TestError_WithData(t *testing.T) {
	tt := []struct {
		name string
		err  *Error[string]
		want *Error[string]
		data any
	}{
		{"nil error", nil, &Error[string]{data: someData}, someData},
		{"non nil error", testError, &Error[string]{
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
