package sampler

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"zipkin_sampler/contract"
	"zipkin_sampler/sampler/mocks"
)

func TestReporter(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bs, err := io.ReadAll(r.Body)
		assert.Nil(t, err)
		assert.Equal(t, `[{}]`, string(bs))

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(([]byte)("ok."))
	}))
	defer ts.Close()

	sampler := new(mocks.TimeBaseSampler)
	sampler.On("ShouldReport", mock.Anything).Return(true)
	reporter := NewReporter(ts.URL, sampler)

	span := new(mocks.Span)
	span.On("IsRoot").Return(false)
	span.On("TraceId").Return("traceId")
	trace := NewTrace(time.Nanosecond, []contract.Span{span})

	reporter.Report(trace)

	sampler.AssertExpectations(t)
	span.AssertExpectations(t)
}

func TestReporterResponseError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write(([]byte)("ok."))
	}))
	defer ts.Close()

	sampler := new(mocks.TimeBaseSampler)
	sampler.On("ShouldReport", mock.Anything).Return(true)
	reporter := NewReporter(ts.URL, sampler)

	span := new(mocks.Span)
	span.On("IsRoot").Return(false)
	span.On("TraceId").Return("traceId")
	trace := NewTrace(time.Nanosecond, []contract.Span{span})

	reporter.Report(trace)

	sampler.AssertExpectations(t)
	span.AssertExpectations(t)
}
