package app

import (
	"context"
	"log"

	userV1 "github.com/Arkosh744/auth-grpc/internal/api/user_v1"
	"github.com/Arkosh744/auth-grpc/internal/client/pg"
	"github.com/Arkosh744/auth-grpc/internal/closer"
	"github.com/Arkosh744/auth-grpc/internal/config"
	userRepo "github.com/Arkosh744/auth-grpc/internal/repo/user"
	userService "github.com/Arkosh744/auth-grpc/internal/service/user"
	"github.com/jackc/pgx/v4/pgxpool"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	grpcConfig config.GRPCConfig

	pgClient       pg.Client
	userRepository userRepo.Repository
	userService    userService.Service

	userImpl *userV1.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) GetPGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := config.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) GetGRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) GetPGClient(ctx context.Context) pg.Client {
	if s.pgClient == nil {
		pgCfg, err := pgxpool.ParseConfig(s.GetPGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to get db config: %s", err.Error())
		}

		cl, err := pg.NewClient(ctx, pgCfg)
		if err != nil {
			log.Fatalf("failed to get pg client: %s", err.Error())

		}

		if cl.PG().Ping(ctx) != nil {
			log.Fatalf("ping error: %s", err)
		}

		closer.Add(cl.Close)

		s.pgClient = cl
	}

	return s.pgClient
}

func (s *serviceProvider) GetUserRepo(ctx context.Context) userRepo.Repository {
	if s.userRepository == nil {
		s.userRepository = userRepo.NewRepository(s.GetPGClient(ctx))
	}

	return s.userRepository
}

func (s *serviceProvider) GetUserService(ctx context.Context) userService.Service {
	if s.userService == nil {
		s.userService = userService.NewService(s.GetUserRepo(ctx))
	}

	return s.userService
}

func (s *serviceProvider) GetUserImpl(ctx context.Context) *userV1.Implementation {
	if s.userImpl == nil {
		s.userImpl = userV1.NewImplementation(s.GetUserService(ctx))
	}

	return s.userImpl
}
