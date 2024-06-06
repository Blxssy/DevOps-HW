package calc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	assert.Equal(t, 8, Add(2, 6))
	assert.NotEqual(t, 8, Add(2, 4))

	assert.Equal(t, -4, Sub(2, 6))
	assert.NotEqual(t, -4, Add(2, 4))

	assert.Equal(t, 12, Mul(2, 6))
	assert.NotEqual(t, 12, Mul(2, 4))
}
