package mocks

import (
	"github.com/stretchr/testify/mock"
	"zipkin_sampler/contract"
)

var _ contract.Sampler = &TimeBaseSampler{}

type TimeBaseSampler struct {
	mock.Mock
}

func (s *TimeBaseSampler) ShouldReport(trace contract.Trace) bool {
	args := s.Called(trace)
	return args.Bool(0)
}
