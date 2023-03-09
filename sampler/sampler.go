package sampler

import (
	"time"

	"github.com/eleven26/zipkin_sampler/contract"
)

var _ contract.Sampler = &TimeBaseSampler{}

type TimeBaseSampler struct {
	t time.Duration
}

func NewTimeBaseSampler(t time.Duration) contract.Sampler {
	return &TimeBaseSampler{t: t}
}

func (t TimeBaseSampler) ShouldReport(trace contract.Trace) bool {
	if !trace.IsRoot() {
		return false
	}

	return time.Since(trace.StartTime()) > t.t
}
