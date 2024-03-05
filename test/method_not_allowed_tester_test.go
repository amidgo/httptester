package test_test

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/amidgo/tester"
	"github.com/amidgo/tester/httptester"
	"github.com/amidgo/tester/httptester/mock"
	"github.com/stretchr/testify/assert"
)

func Test_MethodNotAllowedTester_Name(t *testing.T) {
	tester.RunNamedTesters(t,
		&MethodNotAllowedNameCase{
			Methods:      []string{http.MethodConnect, http.MethodGet, http.MethodDelete},
			ExpectedName: "test allowed methods: CONNECT,GET,DELETE",
		},
		&MethodNotAllowedNameCase{
			Methods:      []string{http.MethodConnect, http.MethodGet, http.MethodDelete, http.MethodPost, http.MethodPut},
			ExpectedName: "test allowed methods: CONNECT,GET,DELETE,POST,PUT",
		},
		&MethodNotAllowedNameCase{
			Methods:      []string{},
			ExpectedName: "test allowed methods: ",
		},
	)
}

type MethodNotAllowedNameCase struct {
	Methods      []string
	ExpectedName string
}

func (c *MethodNotAllowedNameCase) Name() string {
	return fmt.Sprintf("methods %s", strings.Join(c.Methods, ", "))
}

func (c *MethodNotAllowedNameCase) Test(t *testing.T) {
	methodNotAllowedTester := httptester.NewMethodNotAllowedTester(nil, c.Methods...)

	assert.Equal(t, c.ExpectedName, methodNotAllowedTester.Name())
}

func Test_MethodNotAllowedTester_Test(t *testing.T) {
	tester.RunNamedTesters(t,
		&MethodNotAllowedTestCase{
			Methods: []string{
				http.MethodGet,
				http.MethodHead,
				http.MethodPost,
				http.MethodPut,
				http.MethodPatch,
				http.MethodDelete,
				http.MethodConnect,
				http.MethodOptions,
				http.MethodTrace,
			},
		},
		&MethodNotAllowedTestCase{
			Methods: []string{
				http.MethodGet,
				http.MethodHead,
				http.MethodPut,
				http.MethodDelete,
				http.MethodOptions,
				http.MethodTrace,
			},
		},
		&MethodNotAllowedTestCase{
			Methods: []string{},
		},
	)
}

type MethodNotAllowedTestCase struct {
	Methods []string
}

func (c *MethodNotAllowedTestCase) Name() string {
	return fmt.Sprintf("methods %s", strings.Join(c.Methods, ", "))
}

func (c *MethodNotAllowedTestCase) Test(t *testing.T) {
	handler := mock.NewHttpMethodNotAllowedHandler()
	tester := httptester.NewMethodNotAllowedTester(handler, c.Methods...)
	tester.Test(t)
}
