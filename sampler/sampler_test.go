package sampler

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/eleven26/zipkin_sampler/sampler/mocks"
)

func TestTimebaseSampler(t *testing.T) {
	traces := new(mocks.Trace)
	traces.On("IsRoot").Return(false)

	sampler := NewTimeBaseSampler(time.Millisecond * 100)
	assert.False(t, sampler.ShouldReport(traces))
	traces.AssertExpectations(t)

	traces = new(mocks.Trace)
	traces.On("IsRoot").Return(true)
	traces.On("StartTime").Return(time.Now().Add(time.Millisecond * -200))

	sampler = NewTimeBaseSampler(time.Millisecond * 100)
	assert.True(t, sampler.ShouldReport(traces))
	traces.AssertExpectations(t)
}
