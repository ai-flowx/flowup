package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHostInfo(t *testing.T) {
	_, err := hostInfo()
	assert.Equal(t, nil, err)
}
