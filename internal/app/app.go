package app

import (
	"log"
	"net"

	"protected-storage-server/internal/db"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"protected-storage-server/internal/app/interceptors"
	"protected-storage-server/internal/config"
	"protected-storage-server/internal/grpcserver"
	"protected-storage-server/internal/repositories/datarepository"
	"protected-storage-server/internal/repositories/userrepository"
	"protected-storage-server/internal/security"
	"protected-storage-server/internal/service/dataservice"
	"protected-storage-server/internal/service/userservice"
	"protected-storage-server/proto"
)

// GRPCApp запускает GRPC приложение.
type GRPCApp struct {
	GRPCServer *grpc.Server
}

// NewGRPC конструктор GRPCApp
func NewGRPC(cfg config.Config) (*GRPCApp, error) {
	log.Println("creating server")

	db, err := db.InitDB(cfg.DataBaseAddress)
	if err != nil {
		return nil, err
	}

	userRepository := userrepository.New(db)
	rawDataRepository := datarepository.New(db)

	userService := userservice.New(userRepository)

	jwtManager := security.NewJWTManager(cfg.Key, cfg.TokenDuration)
	cipherManager, err := security.NewCipherManager(cfg.Key)
	if err != nil {
		return nil, err
	}

	storageService := dataservice.New(rawDataRepository, cipherManager)
	if err != nil {
		return nil, err
	}

	serverImpl := grpcserver.NewServer(userService, storageService, jwtManager)

	authInterceptor := interceptors.NewAuthInterceptor(jwtManager)
	logInterceptor := interceptors.NewLogInterceptor()

	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			logInterceptor.Unary(),
			authInterceptor.Unary(),
		),
	)

	proto.RegisterGrpcServiceServer(s, serverImpl)
	reflection.Register(s)

	return &GRPCApp{
		GRPCServer: s,
	}, nil

}

// Run запуск сервера
func (app *GRPCApp) Run(cfg config.Config) error {
	listen, err := net.Listen("tcp", cfg.Host)
	if err != nil {
		return err
	}

	log.Println("Start GRPc-server")
	return app.GRPCServer.Serve(listen)
}
