package mocks

import (
	"github.com/stretchr/testify/mock"
	"zipkin_sampler/contract"
)

var _ contract.Span = &Span{}

type Span struct {
	mock.Mock
}

func (s *Span) TraceId() string {
	args := s.Called()
	return args.String(0)
}

func (s *Span) IsRoot() bool {
	args := s.Called()
	return args.Bool(0)
}

func (s *Span) ParentId() string {
	args := s.Called()
	return args.String(0)
}

func (s *Span) Timestamp() int64 {
	args := s.Called()
	return args.Get(0).(int64)
}
