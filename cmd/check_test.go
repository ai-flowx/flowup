package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatchToolchain(t *testing.T) {
	local := []string{
		"other 1.0.0",
		"shai 1.0.0",
		"shgpt1 test",
		"shup 2.0.0",
	}

	remote := []string{
		"shai 1.1.0",
		"shgpt1 1.0.0",
		"shgpt2 1.0.0",
		"shup 1.0.0",
	}

	buf, err := matchToolchain(local, remote)
	assert.Equal(t, nil, err)
	assert.Equal(t, len(remote), len(buf))
}
