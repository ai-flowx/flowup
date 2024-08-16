package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetchHost(t *testing.T) {
	_, err := fetchHost()
	assert.Equal(t, nil, err)
}

func TestFetchToolchain(t *testing.T) {
	buf, err := fetchToolchain("../test/install")
	assert.Equal(t, nil, err)
	assert.Equal(t, "shai 1.0.0", buf[0])
	assert.Equal(t, "shup 1.0.0", buf[1])
}
