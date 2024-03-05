package httptester

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/amidgo/tester"
)

var methods = []string{
	http.MethodGet,
	http.MethodHead,
	http.MethodPost,
	http.MethodPut,
	http.MethodPatch,
	http.MethodDelete,
	http.MethodConnect,
	http.MethodOptions,
	http.MethodTrace,
}

type MethodNotAllowedTester struct {
	AllowedMethods []string
	Handler        http.Handler
}

func NewMethodNotAllowedTester(handler http.Handler, allowedMethods ...string) tester.NamedTester {
	return &MethodNotAllowedTester{
		Handler:        handler,
		AllowedMethods: allowedMethods,
	}
}

func (t *MethodNotAllowedTester) Name() string {
	return fmt.Sprintf("test allowed methods: %s", strings.Join(t.AllowedMethods, ","))
}

func (mt *MethodNotAllowedTester) Test(t *testing.T) {
	allowedMethods := mt.allowedMethods()

	var tester tester.NamedTesterContainer

	for _, method := range methods {
		if _, ok := allowedMethods[method]; ok {
			continue
		}

		tester.AddNamedTester(
			&HttpCase{
				CaseName:           fmt.Sprintf("test method not allowed for %s method", method),
				Method:             method,
				Handler:            mt.Handler,
				ExpectedStatusCode: http.StatusMethodNotAllowed,
			},
		)
	}

	tester.Test(t)
}

func (mt *MethodNotAllowedTester) allowedMethods() map[string]struct{} {
	allowedMethods := make(map[string]struct{}, len(mt.AllowedMethods))

	for _, method := range mt.AllowedMethods {
		allowedMethods[method] = struct{}{}
	}

	return allowedMethods
}
