package sampler

import "zipkin_sampler/contract"

var _ contract.Reporter = &reporter{}

type reporter struct {
	endpoint string
}

func NewReporter(endpoint string) contract.Reporter {
	return &reporter{endpoint: endpoint}
}

func (r *reporter) Report(trace contract.Trace) error {
	panic("implement me")
}
