package main

import (
	"github.com/ogier/pflag"

	"github.com/cgascoig/intersight-metrics/pkg/ismetrics"
	"github.com/sirupsen/logrus"
)

var (
	keyID   = "59c84e4a16267c0001c23428/59cc595416267c0001a0dfc7/5ea26d577564612d3025b892"
	keyFile = "/Users/cgascoig/dev/api-keys/intersight-cgascoig-20200423.pem"

	// command line flags
	configFileName string
	verbose        bool
)

func main() {
	pflag.Parse()

	if verbose {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.Debug("Debug logging enabled")
	}

	im := ismetrics.NewIntersightMetrics(configFileName, verbose)

	im.Start()

	logrus.Info("Shutting down")
}

func init() {
	pflag.StringVarP(&configFileName, "config", "c", "", "Configuration file")
	pflag.Bool(&verbose, "verbose", "v", false, "Verbose logging")
}
