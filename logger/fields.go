package logger

import "go.uber.org/fx"

const (
	VersionField     = "version"
	ApplicationField = "application"
	BuildTimeField   = "build-time"
	ZapFieldsGroup   = `group:"zap-fields"`
)

// AsZapField annotates the given constructor to state that
// it provides a zap field to the "routes" group.
func AsZapField(f any) any {
	return fx.Annotate(f, fx.ResultTags(ZapFieldsGroup))
}

// NewForFx annotates the New constructor to state that it uses
// zap fields by "zap-field" group.
func NewForFx() any {
	return fx.Annotate(New, fx.ParamTags(ZapFieldsGroup))
}
