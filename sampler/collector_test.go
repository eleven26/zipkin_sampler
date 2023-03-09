package sampler

import (
	"io"
	"reflect"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/eleven26/zipkin_sampler/contract"
	"github.com/eleven26/zipkin_sampler/sampler/mocks"
)

func TestNewCollector(t *testing.T) {
	r := new(mocks.Reporter)

	var collectors []contract.Collector
	var wg sync.WaitGroup
	wg.Add(10)
	var mu sync.Mutex
	for i := 0; i < 10; i++ {
		go func() {
			mu.Lock()
			collectors = append(collectors, NewCollector(time.Hour, r))
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()

	first := collectors[0]
	for _, c := range collectors {
		assert.Same(t, first, c)
	}
}

func TestCollectUnmarshalError(t *testing.T) {
	reader := strings.NewReader("test")
	readCloser := io.NopCloser(reader)

	r := new(mocks.Reporter)
	collector := NewCollector(time.Hour, r)

	err := collector.Collect(readCloser)
	assert.NotNil(t, err)
	assert.True(t, strings.Contains(err.Error(), "invalid character"))
}

func TestCollectAdd(t *testing.T) {
	defer func() {
		c = nil
	}()

	reader := strings.NewReader(`[{"timestamp": 123, "traceId": "123456", "parentId": "123"}]`)
	readCloser := io.NopCloser(reader)

	r := new(mocks.Reporter)
	coll := NewCollector(time.Hour, r)
	err := coll.Collect(readCloser)
	assert.Nil(t, err)

	assert.Equal(t, 1, reflect.ValueOf(c).Interface().(*collector).store.Len())
}

func TestCollectReport(t *testing.T) {
	defer func() {
		c = nil
	}()

	reader := strings.NewReader(`[{"timestamp": 123, "traceId": "123456"}]`)
	readCloser := io.NopCloser(reader)

	r := new(mocks.Reporter)
	r.On("Report", mock.Anything).Return(nil).Maybe()

	coll := NewCollector(time.Hour, r)
	err := coll.Collect(readCloser)
	time.Sleep(time.Millisecond * 10)
	assert.Nil(t, err)
	r.AssertExpectations(t)

	assert.Equal(t, 0, reflect.ValueOf(c).Interface().(*collector).store.Len())
}

func TestCollectMerge(t *testing.T) {
	defer func() {
		c = nil
	}()

	childReader := strings.NewReader(`[{"timestamp": 123, "traceId": "123456", "parentId": "123"}]`)
	childReadCloser := io.NopCloser(childReader)

	reader := strings.NewReader(`[{"timestamp": 123, "traceId": "123456"}]`)
	readCloser := io.NopCloser(reader)

	r := new(mocks.Reporter)
	r.On("Report", mock.Anything).Return(nil).Maybe()

	coll := NewCollector(time.Hour, r)
	err := coll.Collect(childReadCloser)
	assert.Nil(t, err)
	err = coll.Collect(readCloser)
	time.Sleep(time.Millisecond)
	assert.Nil(t, err)
	r.AssertExpectations(t)

	assert.Equal(t, 0, reflect.ValueOf(c).Interface().(*collector).store.Len())
}
