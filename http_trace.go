package colly

import (
	"net/http"
	"net/http/httptrace"
	"time"
)

type HTTPTrace struct {
	start, connect    time.Time
	ConnectDuration   time.Duration
	FirstByteDuration time.Duration
}

func (ht *HTTPTrace) trace() *httptrace.ClientTrace { _ = "STUB: not implemented"; return nil }

func (ht *HTTPTrace) WithTrace(req *http.Request) *http.Request {
	_ = "STUB: not implemented"
	return nil
}
