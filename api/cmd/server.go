package cmd

import (
	"VizMigrateX/configs"
	"VizMigrateX/internal/models"
	"VizMigrateX/internal/pkg/lg"
	"VizMigrateX/internal/router"
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"
)

var (
	ServerStartCmd = &cobra.Command{
		Use:   "server",
		Short: `Start the server`,
		Run: func(cmd *cobra.Command, args []string) {
			start()
		},
	}
)

func init() {}

func start() {

	// init router
	router.Init()
	r := router.Router

	// init lg
	lg.Init()
	logger := lg.Logger

	// load env config
	configs.Init()
	EnvConfig := configs.EnvConfig

	// connect database
	models.Connect(EnvConfig.DataSource.Type, EnvConfig.DataSource.Dsn)

	// graceful shutdown
	server := &http.Server{
		Addr:    EnvConfig.Server.Port,
		Handler: r,
	}

	logger.Printf("👻 Server is now listening at port:  %s\n", EnvConfig.Server.Port)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("server listen error: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	i := <-quit
	logger.Println("server receive a signal: ", i.String())

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logger.Fatalf("server shutdown error: %s\n", err)
	}
	logger.Println("Server exiting")

}
