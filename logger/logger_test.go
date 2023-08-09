package logger

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	log, err := New(nil)
	assert.NoError(t, err)
	assert.NotNil(t, log)
}
