package main

import (
	"github.com/ogier/pflag"

	"github.com/cgascoig/intersight-metrics/pkg/ismetrics"
	"github.com/sirupsen/logrus"
)

var (
	// command line flags
	sourceConfigFileName string
	sourceConfigURL      string
	verbose              bool
)

func main() {
	pflag.Parse()

	if verbose {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.Debug("Debug logging enabled")
	}

	if sourceConfigFileName == "" && sourceConfigURL == "" {
		logrus.Fatal("One of --source-file or --source-url is required")
	}

	im := ismetrics.NewIntersightMetrics(sourceConfigFileName, sourceConfigURL, verbose)

	im.Start()

	logrus.Info("Shutting down")
}

func init() {
	pflag.StringVar(&sourceConfigFileName, "source-file", "", "JSON file containing Intersight source configurations")
	pflag.StringVar(&sourceConfigURL, "source-url", "", "URL to retrieve Intersight source configurations")
	pflag.BoolVarP(&verbose, "verbose", "v", false, "Verbose logging")
}
