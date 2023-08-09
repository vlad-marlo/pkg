package logger

import (
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zapcore"
	"testing"
)

func TestNew(t *testing.T) {
	log, err := New(nil)
	assert.NoError(t, err)
	assert.NotNil(t, log)
}

func TestPriorities(t *testing.T) {
	tt := []struct {
		name string
		lvl  zapcore.Level
		low  assert.BoolAssertionFunc
		high assert.BoolAssertionFunc
	}{
		{"debug", zapcore.DebugLevel, assert.True, assert.False},
		{"info", zapcore.InfoLevel, assert.True, assert.False},
		{"warn", zapcore.WarnLevel, assert.True, assert.False},
		{"error", zapcore.ErrorLevel, assert.False, assert.True},
		{"dpanic", zapcore.DPanicLevel, assert.False, assert.True},
		{"panic", zapcore.PanicLevel, assert.False, assert.True},
		{"fatal", zapcore.FatalLevel, assert.False, assert.True},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			high := highPriorityLevelEnablerFunc(tc.lvl)
			low := lowPriorityLevelEnablerFunc(tc.lvl)
			tc.high(t, high)
			tc.low(t, low)
		})
	}
}
