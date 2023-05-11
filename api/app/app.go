package app

import (
	"context"
	"log"
	"net"

	"github.com/Arkosh744/auth-service-api/internal/closer"
	"github.com/Arkosh744/auth-service-api/internal/config"
	desc "github.com/Arkosh744/auth-service-api/pkg/user_v1"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type App struct {
	serviceProvider *serviceProvider
	grpcServer      *grpc.Server
	log             *zap.SugaredLogger
}

func NewApp(ctx context.Context) (*App, error) {
	app := &App{}

	err := app.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return app, nil
}

func (app *App) Run() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	err := app.RunGrpcServer()
	if err != nil {
		return err
	}

	return nil
}

func (app *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		config.Init,
		app.initLogger,
		app.initServiceProvider,
		app.initGrpcServer,
	}

	for _, init := range inits {
		if err := init(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (app *App) initLogger(_ context.Context) error {
	zapLog, err := config.SelectLogger()
	if err != nil {
		log.Fatalf("failed to get logger: %s", err.Error())
	}

	app.log = zapLog.Sugar()

	return nil
}

func (app *App) initServiceProvider(_ context.Context) error {
	app.serviceProvider = newServiceProvider(app.log)

	return nil
}

func (app *App) initGrpcServer(ctx context.Context) error {
	app.grpcServer = grpc.NewServer()
	reflection.Register(app.grpcServer)

	desc.RegisterUserServer(app.grpcServer, app.serviceProvider.GetUserImpl(ctx))

	return nil
}

func (app *App) RunGrpcServer() error {
	list, err := net.Listen("tcp", ":"+app.serviceProvider.GetGRPCConfig().GetPort())
	if err != nil {
		return err
	}

	err = app.grpcServer.Serve(list)
	if err != nil {
		return err
	}

	return nil
}
