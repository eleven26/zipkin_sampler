package mocks

import (
	"github.com/stretchr/testify/mock"

	"zipkin_sampler/contract"
)

var _ contract.Reporter = &Reporter{}

type Reporter struct {
	mock.Mock
}

func (r *Reporter) Report(trace contract.Trace) {
	r.Called(trace)
}
