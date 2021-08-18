package ismetrics

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"github.com/cgascoig/intersight-prometheus-exporter/pkg/intersight"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
)

var collectors = []IntersightCollector{}

func RegisterCollector(c IntersightCollector) {
	collectors = append(collectors, c)
}

type IntersightMetrics struct {
	sourceConfigFileName string
	sourceConfigURL      string
	verboseLogging       bool

	// ctx    context.Context
	client *intersight.APIClient

	keyID   string
	keyFile string

	mutex sync.RWMutex
}

const defaultUpdateInterval = 60

func NewIntersightMetrics(sourceConfigFileName, sourceConfigURL string, verboseLogging bool) *IntersightMetrics {
	im := IntersightMetrics{
		sourceConfigFileName: sourceConfigFileName,
		sourceConfigURL:      sourceConfigURL,
		verboseLogging:       verboseLogging,
	}

	return &im
}

// func (im *IntersightMetrics) setup() (context.Context, error) {
// 	authConfig := intersight.HttpSignatureAuth{
// 		KeyId:          im.keyID,
// 		PrivateKeyPath: im.keyFile,

// 		// Passphrase:           "my-passphrase",
// 		SigningScheme: intersight.HttpSigningSchemeRsaSha256,
// 		SignedHeaders: []string{
// 			intersight.HttpSignatureParameterRequestTarget, // The special (request-target) parameter expresses the HTTP request target.
// 			"Host",   // The Host request header specifies the domain name of the server, and optionally the TCP port number.
// 			"Date",   // The date and time at which the message was originated.
// 			"Digest", // A cryptographic digest of the request body.
// 		},
// 		SigningAlgorithm: intersight.HttpSigningAlgorithmRsaPKCS1v15,
// 	}

// 	authCtx, err := authConfig.ContextWithValue(context.Background())
// 	if err != nil {
// 		return nil, fmt.Errorf("Unable to create request context with authentication: %v", err)
// 	}

// 	// authCtx = context.WithValue(authCtx, intersight.ContextServerVariables, map[string]string{
// 	// 	"server": viper.GetString(serverConfigKey),
// 	// })

// 	return authCtx, nil
// }

func (im *IntersightMetrics) Start() {
	// ctx, err := im.setup()
	// if err != nil {
	// 	logrus.Fatalf("Setup error: %v", err)
	// }

	config := intersight.NewConfiguration()
	config.Debug = im.verboseLogging

	client := intersight.NewAPIClient(config)

	// im.ctx = ctx
	im.client = client

	prometheus.MustRegister(im)

	logrus.Info("Starting metrics server")
	http.Handle("/metrics", promhttp.Handler())
	if err := http.ListenAndServe(":2112", nil); err != nil {
		logrus.Fatal("Error starting HTTP server")
	}
}

func getContextForSource(sc SourceConfig) (context.Context, error) {
	authConfig := intersight.HttpSignatureAuth{
		KeyId: sc.KeyID,

		SigningScheme: intersight.HttpSigningSchemeRsaSha256,
		SignedHeaders: []string{
			intersight.HttpSignatureParameterRequestTarget, // The special (request-target) parameter expresses the HTTP request target.
			"Host",   // The Host request header specifies the domain name of the server, and optionally the TCP port number.
			"Date",   // The date and time at which the message was originated.
			"Digest", // A cryptographic digest of the request body.
		},
		SigningAlgorithm: intersight.HttpSigningAlgorithmRsaPKCS1v15,
	}

	err := authConfig.SetPrivateKey(sc.GetKeyData())
	if err != nil {
		return nil, fmt.Errorf("Unable to set private key: %v", err)
	}

	authCtx, err := authConfig.ContextWithValue(context.Background())
	if err != nil {
		return nil, fmt.Errorf("Unable to create request context with authentication: %v", err)
	}

	return authCtx, nil
}

func (im *IntersightMetrics) Collect(ch chan<- prometheus.Metric) {
	im.mutex.Lock()
	defer im.mutex.Unlock()

	logrus.Info("Collect called")

	var scl SourceConfigList
	var err error

	if im.sourceConfigURL != "" {
		scl, err = GetSourceConfigListFromURL(im.sourceConfigURL)
		if err != nil {
			logrus.Errorf("Unable to get source configuration from URL: %v", err)
			return
		}
	} else if im.sourceConfigFileName != "" {
		scl, err = GetSourceConfigListFromFile(im.sourceConfigFileName)
		if err != nil {
			logrus.Errorf("Unable to get source configuration from file: %v", err)
			return
		}
	} else {
		logrus.Fatal("No source configuration file or URL set")
	}

	for _, sc := range scl {
		logrus.Infof("Starting collection for %s", sc)
		ctx, err := getContextForSource(sc)
		if err != nil {
			logrus.Errorf("Error getting request context with authentication: %v", err)
		}

		for _, c := range collectors {
			c.Collect(im, ctx, ch)
		}
	}
}

func (im *IntersightMetrics) Describe(ch chan<- *prometheus.Desc) {
	logrus.Info("Describe called")

	for _, c := range collectors {
		c.Describe(ch)
	}
}
