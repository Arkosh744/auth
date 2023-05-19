package app

import (
	"context"
	"github.com/Arkosh744/auth-service-api/internal/interceptor"
	"github.com/Arkosh744/auth-service-api/internal/log"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"net"
	"net/http"
	"sync"

	"github.com/Arkosh744/auth-service-api/internal/closer"
	"github.com/Arkosh744/auth-service-api/internal/config"
	desc "github.com/Arkosh744/auth-service-api/pkg/user_v1"
	_ "github.com/Arkosh744/auth-service-api/statik"
	"github.com/rakyll/statik/fs"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type App struct {
	serviceProvider *serviceProvider
	grpcServer      *grpc.Server
	httpServer      *http.Server
	swaggerServer   *http.Server
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

	wg := sync.WaitGroup{}
	wg.Add(3)

	go func() {
		defer wg.Done()

		err := app.RunGrpcServer()
		if err != nil {
			log.Fatalf("failed to run grpc server: %v", err)
		}
	}()

	go func() {
		defer wg.Done()

		err := app.RunHTTPServer()
		if err != nil {
			log.Fatalf("failed to run http server: %v", err)
		}
	}()

	go func() {
		defer wg.Done()

		err := app.RunSwaggerServer()
		if err != nil {
			log.Fatalf("failed to run http server: %v", err)
		}
	}()

	wg.Wait()

	return nil
}

func (app *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		config.Init,
		log.InitLogger,
		app.initServiceProvider,
		app.initGrpcServer,
		app.initHTTPServer,
		app.initSwaggerServer,
	}

	for _, init := range inits {
		if err := init(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (app *App) initServiceProvider(_ context.Context) error {
	app.serviceProvider = newServiceProvider()

	return nil
}

func (app *App) initGrpcServer(ctx context.Context) error {
	app.grpcServer = grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.ValidateInterceptor),
	)
	reflection.Register(app.grpcServer)

	desc.RegisterUserServer(app.grpcServer, app.serviceProvider.GetUserImpl(ctx))

	return nil
}

func (app *App) initHTTPServer(ctx context.Context) error {
	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	err := desc.RegisterUserHandlerFromEndpoint(ctx, mux, app.serviceProvider.GetGRPCConfig().GetHost(), opts)
	if err != nil {
		return err
	}

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Authorization"},
		AllowCredentials: true,
	})

	app.httpServer = &http.Server{
		Addr:    app.serviceProvider.GetHTTPConfig().GetHost(),
		Handler: corsMiddleware.Handler(mux),
	}

	return nil
}

func (app *App) initSwaggerServer(_ context.Context) error {
	statikFs, err := fs.New()
	if err != nil {
		return err
	}

	mux := http.NewServeMux()

	mux.Handle("/", http.StripPrefix("/", http.FileServer(statikFs)))
	mux.HandleFunc("/swagger.json", serveSwaggerFile("/swagger.json"))

	app.swaggerServer = &http.Server{
		Addr:    app.serviceProvider.GetSwaggerConfig().GetHost(),
		Handler: mux,
	}

	return nil
}

func (app *App) RunSwaggerServer() error {
	log.Infof("Swagger server is running on %s", app.serviceProvider.GetSwaggerConfig().GetHost())

	err := app.swaggerServer.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}

func serveSwaggerFile(path string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		statikFs, err := fs.New()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		file, err := statikFs.Open(path)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()

		content, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(content)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (app *App) RunGrpcServer() error {
	log.Infof("GRPC server listening on port %s", app.serviceProvider.GetGRPCConfig().GetHost())

	list, err := net.Listen("tcp", app.serviceProvider.GetGRPCConfig().GetHost())
	if err != nil {
		return err
	}

	err = app.grpcServer.Serve(list)
	if err != nil {
		return err
	}

	return nil
}

func (app *App) RunHTTPServer() error {
	log.Infof("HTTP server listening on port %s", app.serviceProvider.GetHTTPConfig().GetHost())

	err := app.httpServer.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
