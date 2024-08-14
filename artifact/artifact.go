package artifact

import (
	"context"

	"github.com/cligpt/shup/config"
)

type Artifact interface {
	Init(context.Context) error
	Deinit(context.Context) error
	Query(context.Context) ([]string, error)
	Get(context.Context, string, string) ([]string, error)
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

func (a *artifact) Init(_ context.Context) error {
	return nil
}

func (a *artifact) Deinit(_ context.Context) error {
	return nil
}

func (a *artifact) Query(_ context.Context) ([]string, error) {
	// TBD: FIXME
	return []string{}, nil
}

func (a *artifact) Get(_ context.Context, channel, version string) ([]string, error) {
	// TBD: FIXME
	return []string{}, nil
}
