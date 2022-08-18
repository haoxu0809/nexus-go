package log

import (
	"io"
	"net/http"
	"net/url"

	"go.uber.org/zap/zapcore"
)

type context struct {
	request  *request
	response *response
}

type request struct {
	Method string
	Host   string
	Path   string
	Proto  string
	Header http.Header
	Form   url.Values
	Body   any
}

type response struct {
	Status     string
	StatusCode int
	Proto      string
	Header     http.Header
	Body       string
}

func (c *context) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	err := encoder.AddObject("request", c.request)
	if err != nil {
		return err
	}
	err = encoder.AddObject("response", c.response)
	if err != nil {
		return err
	}
	return nil
}

func (r *request) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	encoder.AddString("method", r.Method)
	encoder.AddString("host", r.Host)
	encoder.AddString("path", r.Path)
	encoder.AddString("proto", r.Proto)
	err := encoder.AddReflected("header", r.Header)
	if err != nil {
		return err
	}
	err = encoder.AddReflected("form", r.Form)
	if err != nil {
		return err
	}
	err = encoder.AddReflected("body", r.Body)
	if err != nil {
		return err
	}
	return nil
}

func (r *response) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	encoder.AddString("status", r.Status)
	encoder.AddInt("status_code", r.StatusCode)
	encoder.AddString("proto", r.Proto)
	encoder.AddString("body", r.Body)
	err := encoder.AddReflected("header", r.Header)
	if err != nil {
		return err
	}
	return nil
}

func Context(body any, meta *http.Response) *context {
	bytes, _ := io.ReadAll(meta.Body)
	return &context{
		request: &request{
			Method: meta.Request.Method,
			Host:   meta.Request.URL.Host,
			Path:   meta.Request.URL.Path,
			Proto:  meta.Request.Proto,
			Header: meta.Request.Header,
			Form:   meta.Request.Form,
			Body:   body,
		},
		response: &response{
			Status:     meta.Status,
			StatusCode: meta.StatusCode,
			Proto:      meta.Proto,
			Header:     meta.Header,
			Body:       string(bytes),
		},
	}
}
