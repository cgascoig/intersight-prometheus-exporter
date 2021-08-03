package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/cgascoig/intersight-metrics/pkg/ismetrics"
	"github.com/sirupsen/logrus"
)

var (
	keyID   = "59c84e4a16267c0001c23428/59cc595416267c0001a0dfc7/5ea26d577564612d3025b892"
	keyFile = "/Users/cgascoig/dev/api-keys/intersight-cgascoig-20200423.pem"
)

func main() {
	// client.GetConfig().Debug = true
	logrus.SetLevel(logrus.DebugLevel)
	logrus.Debug("Debug logging enabled")

	// ismetrics.StartPortStats(ctx, client)

	im := ismetrics.NewIntersightMetrics(keyID, keyFile)

	im.Start()

	//
	//Wait forever
	exitSignal := make(chan os.Signal)
	signal.Notify(exitSignal, syscall.SIGINT, syscall.SIGTERM)
	<-exitSignal

	logrus.Info("Shutting down")
}
