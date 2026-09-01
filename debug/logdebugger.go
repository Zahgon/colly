package debug

import (
	"io"
	"log"
	"time"
)

type LogDebugger struct {
	Output io.Writer

	Prefix string

	Flag    int
	logger  *log.Logger
	counter int32
	start   time.Time
}

func (l *LogDebugger) Init() error { _ = "STUB: not implemented"; return nil }

func (l *LogDebugger) Event(e *Event) { _ = "STUB: not implemented"; return }
