package mocks

import (
	"github.com/stretchr/testify/mock"

	"github.com/eleven26/zipkin_sampler/contract"
)

var _ contract.Reporter = &Reporter{}

type Reporter struct {
	mock.Mock
}

func (r *Reporter) Report(trace contract.Trace) {
	r.Called(trace)
}
