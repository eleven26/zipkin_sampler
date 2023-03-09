package sampler

import (
	"sync"
	"time"

	"zipkin_sampler/contract"
)

var _ contract.Store = &store{}

type store struct {
	mu     sync.Mutex
	traces map[string]contract.Trace
}

func NewStore(interval time.Duration) contract.Store {
	s := &store{
		traces: make(map[string]contract.Trace),
	}

	go s.clean(interval)

	return s
}

func (s *store) Add(trace contract.Trace) contract.Trace {
	s.mu.Lock()
	defer s.mu.Unlock()

	v, ok := s.traces[trace.TraceId()]
	if ok {
		v.Merge(trace)
	} else {
		s.traces[trace.TraceId()] = trace
	}

	return s.traces[trace.TraceId()]
}

func (s *store) Remove(traceId string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.traces, traceId)
}

func (s *store) Len() int {
	return len(s.traces)
}

func (s *store) clean(interval time.Duration) {
	ticker := time.NewTicker(interval)

	for range ticker.C {
		s.mu.Lock()

		for traceId, trace := range s.traces {
			if trace.Expired() {
				delete(s.traces, traceId)
			}
		}

		s.mu.Unlock()
	}
}
