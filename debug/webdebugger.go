package debug

import (
	"net/http"
	"sync"
	"time"
)

type WebDebugger struct {
	Address         string
	initialized     bool
	CurrentRequests map[uint32]requestInfo
	RequestLog      []requestInfo
	sync.Mutex
}

type requestInfo struct {
	URL            string
	Started        time.Time
	Duration       time.Duration
	ResponseStatus string
	ID             uint32
	CollectorID    uint32
}

func (w *WebDebugger) Init() error { _ = "STUB: not implemented"; return nil }

func (w *WebDebugger) Event(e *Event) { _ = "STUB: not implemented"; return }

func (w *WebDebugger) indexHandler(wr http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (w *WebDebugger) statusHandler(wr http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}
