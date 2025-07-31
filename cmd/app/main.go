package main

import (
	"context"
	"go_pet_shop/internal/config"
	"go_pet_shop/internal/handlers"
	"go_pet_shop/internal/lib/logger"
	"go_pet_shop/internal/storage/postgres"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {
	cfg := config.MustLoad()

	log := logger.SetupLogger(cfg.Env)
	log.Info("starting go_pet_shop server",
		slog.String("env", cfg.Env),
		slog.String("address", cfg.Address))

	storage, err := postgres.New(cfg.DatabaseURL)
	if err != nil {
		log.Error("failed to connect to database", slog.Any("error", err))
		os.Exit(1)
	}
	defer storage.Close()

	log.Info("database connection established")

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(30 * time.Second))

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	router.Get("/health", handlers.HealthCheck(log))
	router.Get("/ready", handlers.ReadinessCheck(log))
	router.Get("/status", handlers.StatusInfo(log))

	router.Get("/api/test", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "API is working", "timestamp": "` + time.Now().Format(time.RFC3339) + `"}`))
	})

	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      router,
		ReadTimeout:  cfg.Timeout,
		WriteTimeout: cfg.Timeout,
		IdleTimeout:  cfg.IdleTimeout,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Info("server starting", slog.String("address", cfg.Address))

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error("server failed to start", slog.Any("error", err))
			os.Exit(1)
		}
	}()

	log.Info("server started successfully")

	<-done
	log.Info("server stopping...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Error("server forced to shutdown", slog.Any("error", err))
		os.Exit(1)
	}

	log.Info("server stopped gracefully")
}
