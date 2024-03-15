package format

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockAger int

func (m mockAger) Age() int { return int(m) }

type mockNamer string

func (m mockNamer) Name() string { return string(m) }

func TestFormatAge(t *testing.T) {
	p := mockAger(10)
	formatted := Age(p)
	assert.Equal(t, "10 years old", formatted)
}

func TestFormatName(t *testing.T) {
	p := mockNamer("Bob")
	formatted := Name(p)
	assert.Equal(t, "M. Bob", formatted)
}
