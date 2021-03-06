package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-funcards/grpc-server"
	"github.com/go-funcards/logger"
	"github.com/go-funcards/mongodb"
	"github.com/go-funcards/tag-service/internal/config"
	"github.com/go-funcards/tag-service/internal/tag"
	"github.com/go-funcards/tag-service/internal/tag/db"
	"github.com/go-funcards/tag-service/proto/v1"
	"github.com/go-funcards/validate"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"os"
)

//go:generate sh genproto.sh

const envConfigFile = "CONFIG_FILE"

var (
	version    string
	buildDate  string
	buildTime  string
	configFile string
)

func init() {
	configFile = os.Getenv(envConfigFile)
	if configFile == "" {
		flag.StringVar(&configFile, "c", "config.yaml", "application config path")
		flag.Parse()
	}
}

func main() {
	ctx := context.Background()

	log := logger.GetLog().WithFields(logrus.Fields{
		"service": os.Args[0],
		"version": fmt.Sprintf("%s.%s.%s", version, buildDate, buildTime),
	})

	cfg := config.GetConfig(configFile, log)

	validate.Default.RegisterStructRules(cfg.Validation.Rules, []any{
		v1.CreateTagRequest{},
		v1.UpdateTagRequest{},
		v1.DeleteTagRequest{},
		v1.TagsRequest{},
	}...)

	mongoDB := mongodb.GetDB(ctx, cfg.MongoDB.URI, log)
	storage := db.NewStorage(ctx, mongoDB, log)

	register := func(server *grpc.Server) {
		v1.RegisterTagServer(server, tag.NewTagServer(storage))
	}

	lis, err := net.Listen("tcp", cfg.GRPC.Addr)
	if err != nil {
		log.WithField("error", err).Fatal("failed to create tcp listener")
	}

	log.Infof("bind application to addr: %s", lis.Addr().(*net.TCPAddr).String())

	grpcserver.Start(ctx, lis, register, log, grpc.ChainUnaryInterceptor(
		grpc_recovery.UnaryServerInterceptor(),
		grpc_logrus.UnaryServerInterceptor(log),
		mongodb.ErrorUnaryServerInterceptor(),
		validate.DefaultValidatorUnaryServerInterceptor(),
	))
}
