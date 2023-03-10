package sampler

import (
	"encoding/json"

	"github.com/eleven26/zipkin_sampler/contract"
)

var _ contract.Span = &Span{}

type Span map[string]json.RawMessage

func NewSpans(data []byte) ([]contract.Span, error) {
	var spans []Span

	err := json.Unmarshal(data, &spans)
	if err != nil {
		return nil, err
	}

	var result []contract.Span
	for _, span := range spans {
		copySpan := span
		result = append(result, &copySpan)
	}

	return result, nil
}

func (s *Span) IsRoot() bool {
	return s.ParentId() == ""
}

func (s *Span) TraceId() string {
	var v string
	_ = json.Unmarshal((*s)["traceId"], &v)
	return v
}

func (s *Span) ParentId() string {
	var v string
	_ = json.Unmarshal((*s)["parentId"], &v)
	return v
}

func (s *Span) Timestamp() int64 {
	var v int64
	_ = json.Unmarshal((*s)["timestamp"], &v)
	return v
}
