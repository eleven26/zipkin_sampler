package sampler

import (
	"time"

	"github.com/eleven26/zipkin_sampler/contract"
)

var _ contract.Trace = &Trace{}

type Trace struct {
	traceId        string
	spans          []contract.Span
	isRoot         bool
	startTimestamp int64
	expiredAt      time.Time
}

func NewTrace(duration time.Duration, spans []contract.Span) contract.Trace {
	trace := Trace{expiredAt: time.Now().Add(duration)}
	trace.Append(spans)

	return &trace
}

func (t *Trace) IsRoot() bool {
	return t.isRoot
}

func (t *Trace) TraceId() string {
	return t.traceId
}

func (t *Trace) StartTime() time.Time {
	return time.Unix(0, t.startTimestamp*int64(time.Microsecond))
}

func (t *Trace) Append(spans []contract.Span) {
	if len(spans) > 0 {
		t.traceId = spans[0].TraceId()
	}

	for _, span := range spans {
		if span.IsRoot() {
			t.startTimestamp = span.Timestamp()
			t.isRoot = true
		}
	}

	t.spans = append(t.spans, spans...)
}

func (t *Trace) Expired() bool {
	return time.Now().After(t.expiredAt)
}

func (t *Trace) Merge(trace contract.Trace) {
	t.Append(trace.(*Trace).spans)
}

func (t *Trace) Spans() []contract.Span {
	return t.spans
}
