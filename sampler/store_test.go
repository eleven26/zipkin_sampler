package sampler

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"zipkin_sampler/sampler/mocks"
)

func TestStoreAdd(t *testing.T) {
	s := NewStore()
	trace := new(mocks.Trace)
	trace.On("TraceId").Return("1")

	s.Add(trace)

	trace.AssertExpectations(t)
	assert.Len(t, reflect.ValueOf(s).Interface().(*store).traces, 1)

	trace1 := new(mocks.Trace)
	trace1.On("TraceId").Return("1")
	trace.On("Merge", trace1)

	s.Add(trace1)
	trace.AssertExpectations(t)
	trace1.AssertExpectations(t)

	assert.Len(t, reflect.ValueOf(s).Interface().(*store).traces, 1)

	trace2 := new(mocks.Trace)
	trace2.On("TraceId").Return("2")

	s.Add(trace2)
	trace2.AssertExpectations(t)
	assert.Len(t, reflect.ValueOf(s).Interface().(*store).traces, 2)
}

func TestStoreRemove(t *testing.T) {
	s := NewStore()
	trace := new(mocks.Trace)
	trace.On("TraceId").Return("1")

	s.Add(trace)
	trace.AssertExpectations(t)
	assert.Len(t, reflect.ValueOf(s).Interface().(*store).traces, 1)

	s.Remove("1")
	assert.Len(t, reflect.ValueOf(s).Interface().(*store).traces, 0)
}
