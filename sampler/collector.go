package sampler

import (
	"io"
	"sync"
	"time"

	"github.com/eleven26/zipkin_sampler/contract"
)

type collector struct {
	store    contract.Store
	reporter contract.Reporter
	duration time.Duration
}

var (
	c  contract.Collector // singleton
	mu sync.Mutex
)

func NewCollector(duration time.Duration, reporter contract.Reporter) contract.Collector {
	if c != nil {
		return c
	}

	mu.Lock()
	defer mu.Unlock()

	// double-checking
	if c != nil {
		return c
	}

	c = &collector{
		store:    NewStore(time.Minute),
		reporter: reporter,
		duration: duration,
	}

	return c
}

func (c *collector) Collect(rc io.ReadCloser) error {
	body, err := io.ReadAll(rc)
	defer rc.Close()
	if err != nil {
		return err
	}

	spans, err := NewSpans(body)
	if err != nil {
		return err
	}

	trace := NewTrace(c.duration, spans)
	if err != nil {
		return err
	}

	trace = c.store.Add(trace)

	if !trace.IsRoot() {
		return nil
	}

	// 启动协程，异步上报
	go c.reporter.Report(trace)
	c.store.Remove(trace.TraceId())

	return nil
}
