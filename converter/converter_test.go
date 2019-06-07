package converter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvert_10(t *testing.T) {
	assert.Equal(t, "X", Convert(10))
}

func TestConvert_8(t *testing.T) {
	assert.Equal(t, "VIII", Convert(8))
}

func TestConvert_15(t *testing.T) {
	assert.Equal(t, "XV", Convert(15))
}

func TestConvert_9(t *testing.T) {
	assert.Equal(t, "IX", Convert(9))
}
