package proc

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	assert := assert.New(t)
	p := NewProc(11)
	assert.Equal(21, p.Add(10))
}

func TestMul(t *testing.T) {
	assert := assert.New(t)
	p := NewProc(11)
	assert.Equal(110, p.Mul(10))
}

func TestMul2(t *testing.T) {
	assert := assert.New(t)
	p := NewProc(11)
	assert.Equal(121, p.Mul(11))
}

func TestEnv(t *testing.T) {
	assert := assert.New(t)
	p, ok := os.LookupEnv("FOO")
	assert.True(ok, "FOO must be set")
	assert.Equal(p, "BAR")
}
