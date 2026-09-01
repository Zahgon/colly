package colly

import (
	"io"
	"net/http"
	"net/url"
)

type Request struct {
	URL *url.URL

	Headers *http.Header

	Host string

	Ctx *Context

	Depth int

	Method string

	Body io.Reader

	ResponseCharacterEncoding string

	ID        uint32
	collector *Collector
	abort     bool
	baseURL   *url.URL

	ProxyURL string
}

type serializableRequest struct {
	URL     string
	Method  string
	Depth   int
	Body    []byte
	ID      uint32
	Ctx     map[string]interface{}
	Headers http.Header
	Host    string
}

func (r *Request) New(method, URL string, body io.Reader) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Request) Abort() { _ = "STUB: not implemented"; return }

func (r *Request) IsAbort() bool { _ = "STUB: not implemented"; return false }

func (r *Request) AbsoluteURL(u string) string { _ = "STUB: not implemented"; return "" }

func (r *Request) Visit(URL string) error { _ = "STUB: not implemented"; return nil }

func (r *Request) HasVisited(URL string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *Request) Post(URL string, requestData map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Request) PostRaw(URL string, requestData []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Request) PostMultipart(URL string, requestData map[string][]byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Request) Retry() error { _ = "STUB: not implemented"; return nil }

func (r *Request) Do() error { _ = "STUB: not implemented"; return nil }

func (r *Request) Marshal() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
