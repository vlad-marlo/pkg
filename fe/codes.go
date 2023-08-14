package fe

import (
	"net/http"
	"sync"
)

// Code is alias to integer.
//
// Code uses only with this package. You can set up returned code by error with this.
type Code int

// Setup codes with default values.
const (
	CodeInternalServerError Code = iota
	CodeContinue
	CodeSwitchingProtocols
	CodeProcessing
	CodeEarlyHints
	CodeOK
	CodeCreated
	CodeAccepted
	CodeNonAuthoritativeInfo
	CodeNoContent
	CodeResetContent
	CodePartialContent
	CodeMultiStatus
	CodeAlreadyReported
	CodeIMUsed
	CodeMultipleChoices
	CodeMovedPermanently
	CodeFound
	CodeSeeOther
	CodeNotModified
	CodeUseProxy
	CodeTemporaryRedirect
	CodePermanentRedirect
	CodeBadRequest
	CodeUnauthorized
	CodePaymentRequired
	CodeForbidden
	CodeNotFound
	CodeMethodNotAllowed
	CodeNotAcceptable
	CodeProxyAuthRequired
	CodeRequestTimeout
	CodeConflict
	CodeGone
	CodeLengthRequired
	CodePreconditionFailed
	CodeRequestEntityTooLarge
	CodeRequestURITooLong
	CodeUnsupportedMediaType
	CodeRequestedRangeNotSatisfiable
	CodeExpectationFailed
	CodeTeapot
	CodeMisdirectedRequest
	CodeUnprocessableEntity
	CodeLocked
	CodeFailedDependency
	CodeTooEarly
	CodeUpgradeRequired
	CodePreconditionRequired
	CodeTooManyRequests
	CodeRequestHeaderFieldsTooLarge
	CodeUnavailableForLegalReasons
	CodeNotImplemented
	CodeBadGateway
	CodeServiceUnavailable
	CodeGatewayTimeout
	CodeHTTPVersionNotSupported
	CodeVariantAlsoNegotiates
	CodeInsufficientStorage
	CodeLoopDetected
	CodeNotExtended
	CodeNetworkAuthenticationRequired
)

var (
	globalMutex sync.RWMutex
)

func ReplaceCodesTable(table map[Code]int) func() {
	globalMutex.Lock()
	old := httpCodes
	httpCodes = table
	globalMutex.Unlock()
	return func() {
		ReplaceCodesTable(old)
	}
}

func (c Code) HTTP() int {
	globalMutex.RLock()
	code, ok := httpCodes[c]
	globalMutex.RUnlock()
	if !ok {
		return http.StatusInternalServerError
	}
	return code
}

var httpCodes = map[Code]int{
	CodeContinue:                      http.StatusContinue,
	CodeSwitchingProtocols:            http.StatusSwitchingProtocols,
	CodeProcessing:                    http.StatusProcessing,
	CodeEarlyHints:                    http.StatusEarlyHints,
	CodeOK:                            http.StatusOK,
	CodeCreated:                       http.StatusCreated,
	CodeAccepted:                      http.StatusAccepted,
	CodeNonAuthoritativeInfo:          http.StatusNonAuthoritativeInfo,
	CodeNoContent:                     http.StatusNoContent,
	CodeResetContent:                  http.StatusResetContent,
	CodePartialContent:                http.StatusPartialContent,
	CodeMultiStatus:                   http.StatusMultiStatus,
	CodeAlreadyReported:               http.StatusAlreadyReported,
	CodeIMUsed:                        http.StatusIMUsed,
	CodeMultipleChoices:               http.StatusMultipleChoices,
	CodeMovedPermanently:              http.StatusMovedPermanently,
	CodeFound:                         http.StatusFound,
	CodeSeeOther:                      http.StatusSeeOther,
	CodeNotModified:                   http.StatusNotModified,
	CodeUseProxy:                      http.StatusUseProxy,
	CodeTemporaryRedirect:             http.StatusTemporaryRedirect,
	CodePermanentRedirect:             http.StatusPermanentRedirect,
	CodeBadRequest:                    http.StatusBadRequest,
	CodeUnauthorized:                  http.StatusUnauthorized,
	CodePaymentRequired:               http.StatusPaymentRequired,
	CodeForbidden:                     http.StatusForbidden,
	CodeNotFound:                      http.StatusNotFound,
	CodeMethodNotAllowed:              http.StatusMethodNotAllowed,
	CodeNotAcceptable:                 http.StatusNotAcceptable,
	CodeProxyAuthRequired:             http.StatusProxyAuthRequired,
	CodeRequestTimeout:                http.StatusRequestTimeout,
	CodeConflict:                      http.StatusConflict,
	CodeGone:                          http.StatusGone,
	CodeLengthRequired:                http.StatusLengthRequired,
	CodePreconditionFailed:            http.StatusPreconditionFailed,
	CodeRequestEntityTooLarge:         http.StatusRequestEntityTooLarge,
	CodeRequestURITooLong:             http.StatusRequestURITooLong,
	CodeUnsupportedMediaType:          http.StatusUnsupportedMediaType,
	CodeRequestedRangeNotSatisfiable:  http.StatusRequestedRangeNotSatisfiable,
	CodeExpectationFailed:             http.StatusExpectationFailed,
	CodeTeapot:                        http.StatusTeapot,
	CodeMisdirectedRequest:            http.StatusMisdirectedRequest,
	CodeUnprocessableEntity:           http.StatusUnprocessableEntity,
	CodeLocked:                        http.StatusLocked,
	CodeFailedDependency:              http.StatusFailedDependency,
	CodeTooEarly:                      http.StatusTooEarly,
	CodeUpgradeRequired:               http.StatusUpgradeRequired,
	CodePreconditionRequired:          http.StatusPreconditionRequired,
	CodeTooManyRequests:               http.StatusTooManyRequests,
	CodeRequestHeaderFieldsTooLarge:   http.StatusRequestHeaderFieldsTooLarge,
	CodeUnavailableForLegalReasons:    http.StatusUnavailableForLegalReasons,
	CodeInternalServerError:           http.StatusInternalServerError,
	CodeNotImplemented:                http.StatusNotImplemented,
	CodeBadGateway:                    http.StatusBadGateway,
	CodeServiceUnavailable:            http.StatusServiceUnavailable,
	CodeGatewayTimeout:                http.StatusGatewayTimeout,
	CodeHTTPVersionNotSupported:       http.StatusHTTPVersionNotSupported,
	CodeVariantAlsoNegotiates:         http.StatusVariantAlsoNegotiates,
	CodeInsufficientStorage:           http.StatusInsufficientStorage,
	CodeLoopDetected:                  http.StatusLoopDetected,
	CodeNotExtended:                   http.StatusNotExtended,
	CodeNetworkAuthenticationRequired: http.StatusNetworkAuthenticationRequired,
}
