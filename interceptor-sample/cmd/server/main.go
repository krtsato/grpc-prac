package main

import (
	"context"
	"fmt"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"net"
	"os/signal"
	"syscall"
)

func main() {
	// generate logger
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	grpc_zap.ReplaceGrpcLoggerV2(logger)

	// context and err channel
	ctx, cancel := context.WithCancel(context.Background())
	errCh := make(chan error)

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				grpc_zap.UnaryServerInterceptor(logger),
				grpc_auth.UnaryServerInterceptor(auth),
			),
		),
	)

	go func() {
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9999))
		if err != nil {
			errCh <- err
			return
		}
		if err := grpcServer.Serve(lis); err != nil {
			logger.Error("server closed with error", zap.Error(err))
			errCh <- err
		}
	}()

	// Wait signals
	signalHandler := signal.NewHandler(syscall.SIGINT, syscall.SIGTERM)
	defer func() {
		signalHandler.Close()
		cancel()
		grpcServer.GracefulStop()
	}()

	select {
	case err := <-errCh:
		logger.Warn("Terminated with error", zap.Error(err))
	case <-signalHandler.SignalCh:
		logger.Info("Shutdown signal is received")
	}
}

func auth(ctx context.Context) (context.Context, error) {
	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, err
	}

	if token != "hi-mi-tsu" {
		return nil, grpc.Errorf(codes.Unauthenticated, "invalid bearer token")
	}

	return context.WithValue(ctx, "userName", "God"), nil
}
