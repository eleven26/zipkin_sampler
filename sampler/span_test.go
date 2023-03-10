package sampler

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSpans(t *testing.T) {
	bs1 := ([]byte)(`[{"timestamp": 123}]`)
	spans, err := NewSpans(bs1)
	assert.Nil(t, err)
	assert.Len(t, spans, 1)
	assert.Equal(t, int64(123), spans[0].Timestamp())
	assert.Equal(t, "", spans[0].ParentId())
	assert.True(t, spans[0].IsRoot())

	bs2 := ([]byte)(`[{"parentId": "123"}]`)
	spans, err = NewSpans(bs2)
	assert.Nil(t, err)
	assert.Len(t, spans, 1)
	assert.False(t, spans[0].IsRoot())

	bs3 := ([]byte)(`[{"timestamp": 123}, {"timestamp": 456}]`)
	spans, err = NewSpans(bs3)
	assert.Nil(t, err)
	assert.Len(t, spans, 2)

	bs4 := ([]byte)(`just a test.`)
	_, err = NewSpans(bs4)
	assert.NotNil(t, err)
}

func TestNewSpans2(t *testing.T) {
	bs := ([]byte)(`[{"parentId": "123"}, {"parentId": "456"}]`)
	spans, err := NewSpans(bs)
	assert.Nil(t, err)
	assert.Len(t, spans, 2)
	assert.False(t, spans[0].IsRoot())

	assert.Equal(t, "123", spans[0].ParentId())
	assert.Equal(t, "456", spans[1].ParentId())
}

func TestSpan(t *testing.T) {
	span, err := mapToSpan(map[string]any{"timestamp": 123})
	assert.Nil(t, err)
	assert.True(t, span.IsRoot())
	assert.Equal(t, "", span.ParentId())
	assert.Equal(t, int64(123), span.Timestamp())

	span, err = mapToSpan(map[string]any{"parentId": "123", "timestamp": 123})
	assert.Nil(t, err)
	assert.False(t, span.IsRoot())
	assert.Equal(t, "123", span.ParentId())
	assert.Equal(t, int64(123), span.Timestamp())

	span, err = mapToSpan(map[string]any{"parentId": nil, "timestamp": 123})
	assert.Nil(t, err)
	assert.True(t, span.IsRoot())
	assert.Equal(t, "", span.ParentId())
	assert.Equal(t, int64(123), span.Timestamp())

	span, err = mapToSpan(map[string]any{"parentId": "", "timestamp": 123, "traceId": "123456"})
	assert.Nil(t, err)
	assert.True(t, span.IsRoot())
	assert.Equal(t, "", span.ParentId())
	assert.Equal(t, "123456", span.TraceId())
}

func mapToSpan(m map[string]any) (Span, error) {
	bs, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}

	var span Span
	err = json.Unmarshal(bs, &span)

	return span, err
}
