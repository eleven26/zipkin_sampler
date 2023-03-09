package contract

import (
	"io"
	"time"
)

type Trace interface {
	IsRoot() bool
	TraceId() string
	StartTime() time.Time
	Append([]Span)
	Expired() bool
	Merge(Trace)
	Spans() []Span
}

type Sampler interface {
	ShouldReport(Trace) bool
}

type Span interface {
	IsRoot() bool
	TraceId() string
	ParentId() string
	Timestamp() int64
}

type Store interface {
	Add(Trace) Trace
	Remove(traceId string)
	Len() int
}

type Reporter interface {
	Report(Trace)
}

type Collector interface {
	Collect(io.ReadCloser) error
}
