package logger

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLogger_AddTarget(t *testing.T) {

	var target Target

	target = &PrintLogTarget{}
	logger := Logger{}

	logger.AddTarget(&target)
	logger.AddTarget(&target)

	assert.Equal(t, len(logger.Targets), 2)
}
