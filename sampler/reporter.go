package sampler

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"

	"zipkin_sampler/contract"
)

var _ contract.Reporter = &reporter{}

type reporter struct {
	endpoint string
	sampler  contract.Sampler
}

func NewReporter(endpoint string, sampler contract.Sampler) contract.Reporter {
	return &reporter{
		endpoint: endpoint,
		sampler:  sampler,
	}
}

func (r *reporter) Report(trace contract.Trace) {
	if !r.sampler.ShouldReport(trace) {
		return
	}

	r.report(trace)
}

func (r *reporter) report(trace contract.Trace) {
	spans := trace.Spans()
	if len(spans) == 0 {
		return
	}

	data, err := json.Marshal(spans)
	if err != nil {
		log.Error("marshal spans error.", err)
		return
	}

	resp, err := http.Post(r.endpoint, "application/json", bytes.NewReader(data))
	if err != nil {
		log.Error("post spans error.", err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		log.Error("post spans error.", resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		if err == nil {
			log.Error("response body:", string(body))
		}

		return
	}
}
