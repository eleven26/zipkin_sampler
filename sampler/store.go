package sampler

import (
	"sync"

	"zipkin_sampler/contract"
)

var _ contract.Store = &store{}

type store struct {
	mu     sync.RWMutex
	traces map[string]contract.Trace
}

func NewStore() contract.Store {
	return &store{
		traces: make(map[string]contract.Trace),
	}
}

func (s *store) Add(trace contract.Trace) {
	s.mu.Lock()
	defer s.mu.Unlock()

	v, ok := s.traces[trace.TraceId()]
	if ok {
		v.Merge(trace)
	} else {
		s.traces[trace.TraceId()] = trace
	}
}

func (s *store) Remove(traceId string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.traces, traceId)
}
