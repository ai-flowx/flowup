package artifact

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/pkg/errors"

	"github.com/ai-flowx/flowup/config"
)

const (
	apiFetch = "/zd-devops-nj-release-generic/aiops-codegpt/cli"
	apiQuery = "/api/storage/zd-devops-nj-release-generic/aiops-codegpt/cli"
)

var (
	filterName   = []string{"shdrive"}
	filterSuffix = []string{".deb", ".exe", ".msi", ".yml", ".zip"}
)

type Artifact interface {
	Init(context.Context) error
	Deinit(context.Context) error
	Fetch(context.Context, string, string, string, string) error
	Query(context.Context, string, string) ([]string, error)
}

type Config struct {
	Config config.Config
}

type artifact struct {
	cfg *Config
}

func New(_ context.Context, cfg *Config) Artifact {
	return &artifact{
		cfg: cfg,
	}
}

func DefaultConfig() *Config {
	return &Config{}
}

func (a *artifact) Init(_ context.Context) error {
	return nil
}

func (a *artifact) Deinit(_ context.Context) error {
	return nil
}

func (a *artifact) Fetch(_ context.Context, channel, version, name, install string) error {
	helper := func(data []byte, install string) error {
		file, err := os.OpenFile(install, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, config.BinPerm)
		if err != nil {
			return err
		}
		defer func(file *os.File) {
			_ = file.Close()
		}(file)
		if _, err := file.Write(data); err != nil {
			return err
		}
		return nil
	}

	url := a.cfg.Config.Spec.Artifact.Url
	user := a.cfg.Config.Spec.Artifact.User
	pass := a.cfg.Config.Spec.Artifact.Pass

	url = url + apiFetch + "/" + channel

	if channel == config.ChannelRelease {
		if version != "" {
			url += "/" + version + "/" + name
		}
	} else if channel == config.ChannelNightly {
		url += name
	} else {
		return errors.New("invalid channel")
	}

	body, err := a.get(url, user, pass)
	if err != nil {
		return err
	}

	if err := helper(body, install); err != nil {
		return err
	}

	return nil
}

func (a *artifact) Query(_ context.Context, channel, version string) ([]string, error) {
	var body []byte
	var buf map[string]interface{}
	var err error

	url := a.cfg.Config.Spec.Artifact.Url
	user := a.cfg.Config.Spec.Artifact.User
	pass := a.cfg.Config.Spec.Artifact.Pass

	url = url + apiQuery + "/" + channel
	if channel == config.ChannelRelease && version != "" {
		url += "/" + version
	}

	body, err = a.get(url, user, pass)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &buf); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal")
	}

	ret, err := a.filter(buf)
	if err != nil {
		return nil, errors.Wrap(err, "failed to filter")
	}

	return ret, nil
}

func (a *artifact) get(_url, user, pass string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, _url, http.NoBody)
	if err != nil {
		return nil, errors.Wrap(err, "failed to request")
	}

	if user != "" && pass != "" {
		req.SetBasicAuth(user, pass)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to send")
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("invalid status")
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("failed to read")
	}

	return data, nil
}

func (a *artifact) filter(data map[string]interface{}) ([]string, error) {
	matchName := func(data string) bool {
		found := false
		for _, item := range filterName {
			if item == data {
				found = true
				break
			}
		}
		return found
	}

	matchSuffix := func(data string) bool {
		found := false
		for _, item := range filterSuffix {
			if strings.HasSuffix(data, item) {
				found = true
				break
			}
		}
		return found
	}

	var buf []string

	for _, item := range data["children"].([]interface{}) {
		b := item.(map[string]interface{})
		s := strings.TrimPrefix(b["uri"].(string), "/")
		if matchName(s) || matchSuffix(s) {
			continue
		}
		buf = append(buf, s)
	}

	return buf, nil
}
