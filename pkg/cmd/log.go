package cmd

import (
	"BackEndDemo/pkg/http/rest"
	"BackEndDemo/pkg/logger"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func init() {
	rootCmd.AddCommand(logCmd)
}

var logCmd = &cobra.Command{
	Use:   "log",
	Short: "log",
	Long:  `log`,
	Run: func(cmd *cobra.Command, args []string) {
		zap.ReplaceGlobals(logger.NewRolling("backend_demo.log"))

		r := gin.Default()
		r.Use(gzip.Gzip(gzip.BestSpeed))

		logHandler := &rest.LogHandler{}
		logHandler.Router(r)

		go func() {
			if err := r.Run(); err != nil {
				zap.S().Fatal(err)
			}
		}()

		// listen os.Signal
		termChan := make(chan os.Signal, 1)
		signal.Notify(termChan, syscall.SIGTERM, syscall.SIGINT)

		// close goroutine when catch interrupt signal
		<-termChan
		zap.L().Info("SIGTERM received. close goroutine\n")
	},
}
