package rest

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"

	"github.com/parnurzeal/gorequest"
)

type Request struct {
	c       *Client
	timeout time.Duration
	params  url.Values
	headers http.Header

	verb     string
	endpoint string
	err      error
	body     interface{}
}

func NewRequest(c *Client) *Request {
	r := &Request{
		c: c,
	}

	authMethod := 0

	for _, fn := range []func() bool{c.context.HasBasicAuth} {
		if fn() {
			authMethod++
		}
	}

	if authMethod > 1 {
		r.err = fmt.Errorf(
			"username/password or bearer token or secretID/secretKey may be set, but should use only one of them",
		)

		return r
	}

	switch {
	case c.context.HasTokenAuth():
		r.SetHeader("Authorization", fmt.Sprintf("Bearer %s", c.context.BearerToken))
		//case c.context.HasBasicAuth():
		//	r.SetHeader("Authorization", "Basic"+basicAuth(c.context.Username, c.context.Password))
	}

	switch {
	case len(c.context.AcceptContextTypes) > 0:
		r.SetHeader("Accept", c.context.AcceptContextTypes)
	case len(c.context.ContentType) > 0:
		r.SetHeader("Accept", c.context.ContentType+", */*")
	}

	return r
}

func (r *Request) SetHeader(key string, values ...string) *Request {
	if r.headers == nil {
		r.headers = http.Header{}
	}

	r.headers.Del(key)
	for _, value := range values {
		r.headers.Add(key, value)
	}

	return r
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func (r *Request) Verb(verb string) *Request {
	r.verb = verb
	return r
}

func (r *Request) Body(obj interface{}) *Request {
	if v := reflect.ValueOf(obj); v.Kind() == reflect.Struct {
		r.SetHeader("Content-Type", r.c.context.ContentType)
	}

	r.body = obj

	return r
}

func (r *Request) SetEndpoint(endpoint string) *Request {
	r.endpoint = endpoint
	return r
}

func (r *Request) URL() *url.URL {
	finalURL := &url.URL{}
	if r.c.base != nil {
		*finalURL = *r.c.base
	}

	finalURL.Path = r.endpoint
	query := url.Values{}
	for key, values := range r.params {
		for _, value := range values {
			query.Add(key, value)
		}
	}

	if r.timeout != 0 {
		query.Set("timeout", r.timeout.String())
	}

	finalURL.RawQuery = query.Encode()
	return finalURL
}

func (r *Request) Do() Result {
	client := r.c.Client
	client.Header = r.headers
	if r.c.context.HasBasicAuth() {
		client.SetBasicAuth(r.c.context.Username, r.c.context.Password)
	}
	resp, body, errs := client.CustomMethod(r.verb, r.URL().String()).Send(r.body).EndBytes()
	if err := combineErr(resp.StatusCode, body, errs); err != nil {
		return Result{
			response: &resp,
			err:      err,
			body:     body,
		}
	}

	return Result{
		response: &resp,
		body:     body,
	}
}

func (r *Request) Param(params url.Values) *Request {
	if r.err != nil {
		return r
	}

	r.params = params

	return r
}

func combineErr(statusCode int, body []byte, errs []error) error {
	var e, sep string
	if len(errs) > 0 {
		for _, err := range errs {
			e = sep + err.Error()
			sep = "\n"
		}
		return errors.New(e)
	}

	if statusCode >= 400 {
		return errors.New(string(body))
	}

	return nil
}

// NameMayNotBe specifies strings that cannot be used as names specified as
// path segments (like the REST API or etcd store).
var NameMayNotBe = []string{".", ".."}

// NameMayNotContain specifies substrings that cannot be used in names specified
// as path segments (like the REST API or etcd store).
var NameMayNotContain = []string{"/", "%"}

// IsValidPathSegmentName validates the name can be safely encoded as a path segment.
func IsValidPathSegmentName(name string) []string {
	for _, illegalName := range NameMayNotBe {
		if name == illegalName {
			return []string{fmt.Sprintf(`may not be '%s'`, illegalName)}
		}
	}

	var errs []string
	for _, illegalContext := range NameMayNotContain {
		if strings.Contains(name, illegalContext) {
			errs = append(errs, fmt.Sprintf(`may not contain '%s'`, illegalContext))
		}
	}
	return errs
}

// IsValidPathSegmentPrefix validates the name can be used as a prefix for a name
// which will be encoded as a path segment. It does not check for exact matches
// with disallowed names, since an arbitrary suffix might make the name valid.
func IsValidPathSegmentPrefix(name string) []string {
	var errs []string

	for _, illegalContent := range NameMayNotContain {
		if strings.Contains(name, illegalContent) {
			errs = append(errs, fmt.Sprintf(`may not contain '%s'`, illegalContent))
		}
	}

	return errs
}

// ValidatePathSegmentName validates the name can be safely encoded as a path segment.
func ValidatePathSegmentName(name string, prefix bool) []string {
	if prefix {
		return IsValidPathSegmentPrefix(name)
	}

	return IsValidPathSegmentName(name)
}

type Result struct {
	response *gorequest.Response
	err      error
	body     []byte
}

func (r Result) Response() *http.Response {
	return &**r.response
}

func (r Result) Raw() ([]byte, error) {
	return r.body, r.err
}

func (r Result) Error() error {
	return r.err
}

func (r Result) Into(v interface{}) error {
	if r.err != nil {
		return r.Error()
	}

	if err := Decode(r.body, &v); err != nil {
		return err
	}

	return nil
}

func Decode(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func Encode(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}
