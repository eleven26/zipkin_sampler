package sampler

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"zipkin_sampler/contract"
	"zipkin_sampler/sampler/mocks"
)

func TestAppend(t *testing.T) {
	var startTimestamp int64 = 1678254537979162

	span1 := new(mocks.Span)
	span1.On("IsRoot").Return(true)
	span1.On("Timestamp").Return(startTimestamp)
	span1.On("TraceId").Return("traceId")

	span2 := new(mocks.Span)
	span2.On("IsRoot").Return(false)

	trace, err := NewTrace(time.Hour, []contract.Span{})
	assert.Nil(t, err)
	trace.Append([]contract.Span{span1, span2})

	assert.True(t, trace.IsRoot())
	assert.Equal(t, "traceId", trace.TraceId())
	assert.Equal(t, startTimestamp, trace.StartTime().UnixMicro())

	span1.AssertExpectations(t)
	span2.AssertExpectations(t)
}

func TestEmptySpans(t *testing.T) {
	var spans []contract.Span
	trace, err := NewTrace(time.Hour, spans)
	assert.Nil(t, err)
	assert.Zero(t, trace.IsRoot())
	assert.Equal(t, 1970, trace.StartTime().Year())
}

func TestExpired(t *testing.T) {
	span3 := new(mocks.Span)
	span3.On("IsRoot").Return(false)
	span3.On("TraceId").Return("traceId")
	trace, err := NewTrace(time.Nanosecond, []contract.Span{span3})
	assert.Nil(t, err)
	time.Sleep(time.Millisecond * 10)
	assert.True(t, trace.Expired())

	span3.AssertExpectations(t)
}

func TestMerge(t *testing.T) {
	var spans []contract.Span
	trace, err := NewTrace(time.Hour, spans)
	assert.Nil(t, err)

	span := new(mocks.Span)
	span.On("IsRoot").Return(false)
	span.On("TraceId").Return("traceId")

	spans = []contract.Span{span}
	trace1, err := NewTrace(time.Hour, spans)
	assert.Nil(t, err)

	trace.Merge(trace1)

	assert.Equal(t, 1, len(trace.(*Trace).spans))
	span.AssertExpectations(t)
}
