package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatchToolchain(t *testing.T) {
	local := []string{
		"flowx 1.0.0",
		"flowup 2.0.0",
	}

	remote := []string{
		"flowx 1.1.0",
		"flowup 1.0.0",
	}

	buf, err := matchToolchain(local, remote)
	assert.Equal(t, nil, err)
	assert.Equal(t, len(remote), len(buf))
}
