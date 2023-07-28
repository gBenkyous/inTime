package keisan

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTasizan(t *testing.T) {
	var expected int = 2
	var actual int = Tasizan(2, 1)
	assert.Equal(t, expected, actual)
}

func TestHikizan(t *testing.T) {
	var expected int = 1
	var actual int = Hikizan(2, 1)
	assert.Equal(t, expected, actual)
}

func TestKakezan(t *testing.T) {
	var expected int = 2
	var actual int = Kakezan(2, 1)
	assert.Equal(t, expected, actual)
}
