package contract

import "time"

type Trace interface {
	IsRoot() bool
	TraceId() string
	StartTime() time.Time
	Append([]Span)
	Expired() bool
	Merge(Trace)
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
	Add(Trace)
	Remove(traceId string)
}

type Reporter interface {
	Report(Trace) error
}
