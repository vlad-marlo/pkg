package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func highPriorityLevelEnablerFunc(lvl zapcore.Level) bool {
	return lvl >= zapcore.ErrorLevel
}

func lowPriorityLevelEnablerFunc(lvl zapcore.Level) bool {
	return lvl < zapcore.ErrorLevel
}

// New creates logger.
//
// If all is good, developer can get access to logger by zap.L().
func New(fields []zap.Field) (*zap.Logger, error) {
	highPriority := zap.LevelEnablerFunc(highPriorityLevelEnablerFunc)
	lowPriority := zap.LevelEnablerFunc(lowPriorityLevelEnablerFunc)

	consoleDebugging := zapcore.Lock(os.Stdout)
	consoleErrors := zapcore.Lock(os.Stderr)
	jsonEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())

	core := zapcore.NewTee(
		zapcore.NewCore(jsonEncoder, consoleErrors, highPriority),
		zapcore.NewCore(jsonEncoder, consoleDebugging, lowPriority),
	)
	l := zap.New(core).With(fields...)
	zap.ReplaceGlobals(l)
	return l, nil
}
