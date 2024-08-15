package artifact

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	var buf map[string]interface{}

	a := artifact{
		cfg: &Config{},
	}

	data, _ := os.ReadFile("../test/artifact/folder.json")
	_ = json.Unmarshal(data, &buf)

	ret, err := a.filter(buf)
	assert.Equal(t, nil, err)
	assert.Equal(t, len(buf["children"].([]interface{})), len(ret))

	data, _ = os.ReadFile("../test/artifact/files.json")
	_ = json.Unmarshal(data, &buf)

	_, err = a.filter(buf)
	assert.Equal(t, nil, err)
}
