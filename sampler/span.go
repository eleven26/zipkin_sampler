package sampler

import (
	"encoding/json"

	"github.com/spf13/cast"

	"github.com/eleven26/zipkin_sampler/contract"
)

var _ contract.Span = &Span{}

type Span map[string]interface{}

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

func (s *Span) TraceId() string {
	return (*s)["traceId"].(string)
}

func (s *Span) IsRoot() bool {
	v, ok := (*s)["parentId"]
	if !ok {
		return true
	}

	if v == "" {
		return true
	}

	return v == nil
}

func (s *Span) ParentId() string {
	v, ok := (*s)["parentId"]
	if !ok {
		return ""
	}

	if v == nil {
		return ""
	}

	return v.(string)
}

func (s *Span) Timestamp() int64 {
	return cast.ToInt64((*s)["timestamp"])
}
