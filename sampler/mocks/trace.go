package mocks

import (
	"time"

	"github.com/stretchr/testify/mock"

	"zipkin_sampler/contract"
)

var _ contract.Trace = &Trace{}

type Trace struct {
	mock.Mock
}

func (t *Trace) Merge(trace contract.Trace) {
	t.Called(trace)
}

func (t *Trace) TraceId() string {
	args := t.Called()
	return args.String(0)
}

func (t *Trace) IsRoot() bool {
	args := t.Called()
	return args.Bool(0)
}

func (t *Trace) StartTime() time.Time {
	args := t.Called()
	return args.Get(0).(time.Time)
}

func (t *Trace) Append(traces []contract.Span) {
	t.Called(traces)
}

func (t *Trace) Expired() bool {
	args := t.Called()
	return args.Bool(0)
}
