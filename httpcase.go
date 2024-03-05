package httptester

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

type HttpCase struct {
	CaseName    string
	Method      string
	Target      string
	QueryValues map[string]string
	Header      http.Header

	Handler http.Handler

	Input          []byte
	ExpectedOutput []byte
	ExpectedHeader http.Header

	ExpectedStatusCode int
}

func (c *HttpCase) Name() string {
	return c.CaseName
}

func (c *HttpCase) Test(t *testing.T) {
	req := c.Request()

	rec := httptest.NewRecorder()
	c.Handler.ServeHTTP(rec, req)

	c.assertStatusCode(t, rec)

	if c.ExpectedOutput != nil {
		c.assertOutput(t, rec)
	}

	if c.ExpectedHeader != nil {
		c.assertHeader(t, rec)
	}
}

func (c *HttpCase) Request() *http.Request {
	req := httptest.NewRequest(c.Method, c.target()+c.query(), bytes.NewReader(c.Input))
	setRequestHeader(req, c.Header)

	return req
}

func (c *HttpCase) target() string {
	if c.Target == "" {
		return "/any/target?"
	}

	return c.Target
}

func (c *HttpCase) query() string {
	if c.QueryValues == nil {
		return ""
	}

	urlValues := make(url.Values)
	for key, value := range c.QueryValues {
		urlValues.Set(key, value)
	}

	return urlValues.Encode()
}

func setRequestHeader(req *http.Request, header http.Header) {
	if header == nil {
		return
	}

	for key := range header {
		for _, value := range header.Values(key) {
			req.Header.Add(key, value)
		}
	}
}

func (c *HttpCase) assertStatusCode(t *testing.T, rec *httptest.ResponseRecorder) {
	statusCode := rec.Result().StatusCode

	assert.Equal(t, c.ExpectedStatusCode, statusCode, "wrong status code")
}

func (c *HttpCase) assertOutput(t *testing.T, rec *httptest.ResponseRecorder) {
	buf := bytes.Buffer{}
	buf.ReadFrom(rec.Result().Body)

	output := buf.Bytes()

	assert.Equal(t, c.ExpectedOutput, output, "output not equal")
}

func (c *HttpCase) assertHeader(t *testing.T, rec *httptest.ResponseRecorder) {
	if len(rec.Header()) != len(c.ExpectedHeader) {
		t.Fatalf("header length not equal, expected lenth %d, actual %d", len(c.ExpectedHeader), len(rec.Header()))
	}

	for key := range rec.Header() {
		actual := rec.Header().Values(key)
		expected := c.ExpectedHeader.Values(key)

		assert.Equal(t, expected, actual, "header values by key %s not equal")
	}
}
