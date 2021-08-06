package ismetrics

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetIntervalString(t *testing.T) {
	endTime := time.Date(2020, time.November, 21, 0, 20, 0, 0, time.UTC)
	duration := time.Duration(20) * time.Minute

	assert.Equal(t, "2020-11-21T00:00:00+00:00/2020-11-21T00:20:00+00:00", getIntervalString(endTime, duration))
}
