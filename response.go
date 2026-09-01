package colly

import (
	"net/http"
)

type Response struct {
	StatusCode int

	Body []byte

	Ctx *Context

	Request *Request

	Headers *http.Header

	Trace *HTTPTrace
}

func (r *Response) Save(fileName string) error { _ = "STUB: not implemented"; return nil }

func (r *Response) FileName() string { _ = "STUB: not implemented"; return "" }

func (r *Response) fixCharset(detectCharset bool, defaultEncoding string) error {
	_ = "STUB: not implemented"
	return nil
}

func encodeBytes(b []byte, contentType string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
