package user

import (
	"fmt"
	upbv1 "mxshop/api/user/v1"
	"mxshop/app/user/srv/config"
	usercontroller "mxshop/app/user/srv/controller/user"
	db2 "mxshop/app/user/srv/data/v1/db"
	servicev1 "mxshop/app/user/srv/service/v1"
	"mxshop/gmicro/core/trace"
	"mxshop/gmicro/server/rpcserver"
	"mxshop/pkg/log"
)

func NewUserRPCServer(cfg *config.Config) (*rpcserver.Server, error) {
	trace.InitAgent(trace.Options{
		cfg.Telemetry.Name,
		cfg.Telemetry.Endpoint,
		cfg.Telemetry.Sampler,
		cfg.Telemetry.Batcher,
	})

	dataFactory, err := db2.GetDBFactoryOr(cfg.MySQLOptions)
	if err != nil {
		log.Fatal(err.Error())
	}

	userService := servicev1.NewUserService(dataFactory)
	userServer := usercontroller.NewUserServer(userService)
	rpcAddr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	grpcServer := rpcserver.NewServer(
		rpcserver.WithAddress(rpcAddr),
		rpcserver.WithMetrics(cfg.Server.EnableMetrics),
	)

	upbv1.RegisterUserServer(grpcServer.Server, userServer)

	return grpcServer, nil
}
