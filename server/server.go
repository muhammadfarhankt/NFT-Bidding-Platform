package server

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/robfig/cron/v3"

	"github.com/muhammadfarhankt/NFT-Bidding-Platform/config"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/middleware/middlewareHandler"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/middleware/middlewareRepository"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/modules/middleware/middlewareUsecase"
	"github.com/muhammadfarhankt/NFT-Bidding-Platform/pkg/jwtAuth"

	"go.mongodb.org/mongo-driver/mongo"
)

type (
	server struct {
		app        *echo.Echo
		db         *mongo.Client
		cfg        *config.Config
		middleware middlewareHandler.MiddlewareHandlerService
		cron       *cron.Cron
	}
)

func newMiddleware(cfg *config.Config) middlewareHandler.MiddlewareHandlerService {
	repo := middlewareRepository.NewMiddlewareRepository()
	usecase := middlewareUsecase.NewMiddlewareUsecase(repo)
	return middlewareHandler.NewMiddlewareHandler(cfg, usecase)
}

func (s *server) gracefulShutdown(pctx context.Context, quit <-chan os.Signal) {
	log.Printf("Start service: %s", s.cfg.App.Name)

	<-quit
	log.Printf("Shutting down service: %s", s.cfg.App.Name)

	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	if err := s.app.Shutdown(ctx); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func (s *server) httpListening() {
	if err := s.app.Start(s.cfg.App.Url); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Error: %v", err)
	}
}

func Start(pctx context.Context, cfg *config.Config, db *mongo.Client) {
	s := &server{
		app:        echo.New(),
		db:         db,
		cfg:        cfg,
		middleware: newMiddleware(cfg),
		cron:       cron.New(),
	}

	jwtAuth.SetApiKey(cfg.Jwt.ApiSecretKey)

	// Basic Middleware
	// Request Timeout
	s.app.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Skipper:      middleware.DefaultSkipper,
		ErrorMessage: "Error: Request Timeout",
		Timeout:      30 * time.Second,
	}))

	// CORS
	s.app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.PATCH, echo.DELETE},
	}))

	// Body Limit
	s.app.Use(middleware.BodyLimit("10M"))

	switch s.cfg.App.Name {
	case "auth":
		s.authService()
	case "user":
		s.userService()
	case "nft":
		s.nftService()
		// Create a new cron job
		_, err := s.cron.AddFunc("* * * * *", s.executeCronJob)
		if err != nil {
			log.Fatal(err)
		}

		// Start the cron scheduler
		s.cron.Start()
	case "inventory":
		s.inventoryService()
	case "payment":
		s.paymentService()
	}

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	s.app.Use(middleware.Logger())

	go s.gracefulShutdown(pctx, quit)

	// Listening
	s.httpListening()
}

func (s *server) executeCronJob() {
	// Your cron job logic here
	fmt.Println("Executing cron job", time.Now())

	// Make GET request - http://localhost:1324/nft_v1/nft/bidding/execute-bids
	resp, err := http.Get("http://" + s.cfg.App.Url + "/nft_v1/nft/bidding/execute-bids")
	if err != nil {
		log.Println("Error making GET request:", err)
		return
	}
	defer resp.Body.Close()

	// Handle the response
	if resp.StatusCode != http.StatusOK {
		log.Printf("Unexpected status code: %d", resp.StatusCode)
		return
	}

	// Read the response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return
	}
	log.Println("Response body:", string(respBody))
	// Process the response body
	// ...

}
