package drive

import (
	"context"
	"math"
	"strconv"
	"time"

	"github.com/hashicorp/go-hclog"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/cligpt/shup/config"
	rpc "github.com/cligpt/shup/drive/rpc"
)

const (
	upTiemout = "10s"
)

type Drive interface {
	Init(context.Context) error
	Deinit(context.Context) error
	Run(context.Context, string, string, string) (string, string, string, string, error)
}

type Config struct {
	Logger hclog.Logger
	Config config.Config
}

type drive struct {
	cfg    *Config
	client rpc.RpcProtoClient
	conn   *grpc.ClientConn
}

func New(_ context.Context, cfg *Config) Drive {
	return &drive{
		cfg: cfg,
	}
}

func DefaultConfig() *Config {
	return &Config{}
}

func (d *drive) Init(ctx context.Context) error {
	if err := d.initConn(ctx); err != nil {
		return errors.Wrap(err, "failed to init conn")
	}

	return nil
}

func (d *drive) Deinit(ctx context.Context) error {
	_ = d.deinitConn(ctx)

	return nil
}

func (d *drive) Run(ctx context.Context, name, arch, os string) (version, url, user, pass string, err error) {
	ret, err := d.sendQuery(ctx, name, arch, os)
	if err != nil {
		return version, url, user, pass, errors.Wrap(err, "failed to query")
	}

	return ret.GetVersion(), ret.GetUrl(), ret.GetUser(), ret.GetPass(), nil
}

func (d *drive) initConn(_ context.Context) error {
	var err error

	host := d.cfg.Config.Spec.Drive.Host
	port := d.cfg.Config.Spec.Drive.Port

	d.conn, err = grpc.NewClient(host+":"+strconv.Itoa(port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(math.MaxInt32), grpc.MaxCallSendMsgSize(math.MaxInt32)))
	if err != nil {
		return errors.Wrap(err, "failed to dial")
	}

	d.client = rpc.NewRpcProtoClient(d.conn)

	return nil
}

func (d *drive) deinitConn(_ context.Context) error {
	return d.conn.Close()
}

func (d *drive) sendQuery(ctx context.Context, name, arch, os string) (*rpc.QueryReply, error) {
	ctx, cancel := context.WithTimeout(ctx, d.setTimeout(upTiemout))
	defer cancel()

	reply, err := d.client.SendQuery(ctx, &rpc.QueryRequest{
		Name: name,
		Arch: arch,
		Os:   os,
	})

	if err != nil {
		return nil, errors.Wrap(err, "failed to send")
	}

	return reply, nil
}

func (d *drive) setTimeout(timeout string) time.Duration {
	duration, _ := time.ParseDuration(timeout)

	return duration
}
