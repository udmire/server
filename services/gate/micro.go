package gate

import (
	"context"
	"crypto/tls"
	"os"
	"sync"

	grpc_client "github.com/asim/go-micro/plugins/client/grpc/v3"
	grpc_server "github.com/asim/go-micro/plugins/server/grpc/v3"
	"github.com/asim/go-micro/plugins/transport/tcp/v3"
	"github.com/asim/go-micro/plugins/wrapper/monitoring/prometheus/v3"
	ratelimit "github.com/asim/go-micro/plugins/wrapper/ratelimiter/ratelimit/v3"
	"github.com/asim/go-micro/v3"
	micro_logger "github.com/asim/go-micro/v3/logger"
	"github.com/asim/go-micro/v3/server"
	"github.com/asim/go-micro/v3/store"
	"github.com/asim/go-micro/v3/transport"
	"github.com/east-eden/server/logger"
	juju_ratelimit "github.com/juju/ratelimit"
	"github.com/rs/zerolog/log"
	cli "github.com/urfave/cli/v2"

	// micro plugins
	_ "github.com/asim/go-micro/plugins/broker/nsq/v3"
	_ "github.com/asim/go-micro/plugins/registry/consul/v3"
)

type MicroService struct {
	srv   micro.Service
	store store.Store
	g     *Gate
	sync.RWMutex
}

func NewMicroService(ctx *cli.Context, g *Gate) *MicroService {
	// cert
	certPath := ctx.String("cert_path_release")
	keyPath := ctx.String("key_path_release")

	if ctx.Bool("debug") {
		certPath = ctx.String("cert_path_debug")
		keyPath = ctx.String("key_path_debug")
	}

	tlsConf := &tls.Config{InsecureSkipVerify: true}
	cert, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("load certificates failed")
	}
	tlsConf.Certificates = []tls.Certificate{cert}

	err = micro_logger.Init(micro_logger.WithOutput(logger.Logger))
	if err != nil {
		log.Fatal().Err(err).Msg("micro_logger init failed")
	}

	s := &MicroService{
		g: g,
	}

	bucket := juju_ratelimit.NewBucket(ctx.Duration("rate_limit_interval"), int64(ctx.Int("rate_limit_capacity")))
	s.srv = micro.NewService(
		micro.Server(
			grpc_server.NewServer(
				server.WrapHandler(ratelimit.NewHandlerWrapper(bucket, false)),
			),
		),

		micro.Client(
			grpc_client.NewClient(),
		),

		micro.Name("gate"),
		micro.WrapHandler(prometheus.NewHandlerWrapper()),

		micro.Transport(tcp.NewTransport(
			transport.TLSConfig(tlsConf),
		)),
	)

	// set environment
	os.Setenv("MICRO_SERVER_ID", ctx.String("gate_id"))

	if ctx.Bool("debug") {
		os.Setenv("MICRO_REGISTRY", ctx.String("registry_debug"))
		// os.Setenv("MICRO_REGISTRY_ADDRESS", ctx.String("registry_address_debug"))
		os.Setenv("MICRO_BROKER", ctx.String("broker_debug"))
		os.Setenv("MICRO_BROKER_ADDRESS", ctx.String("broker_address_debug"))
	} else {
		os.Setenv("MICRO_REGISTRY", ctx.String("registry_release"))
		os.Setenv("MICRO_REGISTRY_ADDRESS", ctx.String("registry_address_release"))
		os.Setenv("MICRO_BROKER", ctx.String("broker_release"))
		os.Setenv("MICRO_BROKER_ADDRESS", ctx.String("broker_address_release"))
	}

	s.srv.Init()

	return s
}

func (s *MicroService) Run(ctx context.Context) error {

	// Run service
	if err := s.srv.Run(); err != nil {
		return err
	}

	return nil
}

func (s *MicroService) StoreWrite(key string, value string) error {
	return s.store.Write(&store.Record{
		Key:   key,
		Value: []byte(value),
	})
}
