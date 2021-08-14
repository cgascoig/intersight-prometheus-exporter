package ismetrics

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"github.com/cgascoig/intersight-metrics/pkg/intersight"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
)

var collectors = []IntersightCollector{}

func RegisterCollector(c IntersightCollector) {
	collectors = append(collectors, c)
}

type IntersightMetrics struct {
	configFileName string
	verboseLogging bool

	ctx    context.Context
	client *intersight.APIClient

	keyID   string
	keyFile string

	mutex sync.RWMutex
}

const defaultUpdateInterval = 60

func NewIntersightMetrics(configFileName string, verboseLogging bool) *IntersightMetrics {
	im := IntersightMetrics{
		configFileName: configFileName,
		verboseLogging: verboseLogging,
	}

	return &im
}

func (im *IntersightMetrics) setup() (context.Context, error) {
	authConfig := intersight.HttpSignatureAuth{
		KeyId:          im.keyID,
		PrivateKeyPath: im.keyFile,

		// Passphrase:           "my-passphrase",
		SigningScheme: intersight.HttpSigningSchemeRsaSha256,
		SignedHeaders: []string{
			intersight.HttpSignatureParameterRequestTarget, // The special (request-target) parameter expresses the HTTP request target.
			"Host",   // The Host request header specifies the domain name of the server, and optionally the TCP port number.
			"Date",   // The date and time at which the message was originated.
			"Digest", // A cryptographic digest of the request body.
		},
		SigningAlgorithm: intersight.HttpSigningAlgorithmRsaPKCS1v15,
	}

	authCtx, err := authConfig.ContextWithValue(context.Background())
	if err != nil {
		return nil, fmt.Errorf("Unable to create request context with authentication: %v", err)
	}

	// authCtx = context.WithValue(authCtx, intersight.ContextServerVariables, map[string]string{
	// 	"server": viper.GetString(serverConfigKey),
	// })

	return authCtx, nil
}

func (im *IntersightMetrics) Start() {
	ctx, err := im.setup()
	if err != nil {
		logrus.Fatalf("Setup error: %v", err)
	}

	config := intersight.NewConfiguration()
	config.Debug = im.verboseLogging

	client := intersight.NewAPIClient(config)

	im.ctx = ctx
	im.client = client

	prometheus.MustRegister(im)

	logrus.Info("Starting metrics server")
	http.Handle("/metrics", promhttp.Handler())
	if err = http.ListenAndServe(":2112", nil); err != nil {
		logrus.Fatal("Error starting HTTP server")
	}
}

func (im *IntersightMetrics) Collect(ch chan<- prometheus.Metric) {
	im.mutex.Lock()
	defer im.mutex.Unlock()

	logrus.Info("Collect called")

	for _, c := range collectors {
		c.Collect(im, ch)
	}
}

func (im *IntersightMetrics) Describe(ch chan<- *prometheus.Desc) {
	logrus.Info("Describe called")

	for _, c := range collectors {
		c.Describe(ch)
	}
}
