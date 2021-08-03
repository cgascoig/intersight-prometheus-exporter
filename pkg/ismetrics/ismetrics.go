package ismetrics

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cgascoig/intersight-metrics/pkg/intersight"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
)

type IntersightMetrics struct {

	// Intersight API Credentials
	keyID   string
	keyFile string

	ctx    context.Context
	client *intersight.APIClient

	updateInterval int

	// Port Stats metric
	portStatsMetric *prometheus.GaugeVec
}

const defaultUpdateInterval = 60

func NewIntersightMetrics(keyID string, keyFile string) *IntersightMetrics {
	im := IntersightMetrics{
		keyID:          keyID,
		keyFile:        keyFile,
		updateInterval: defaultUpdateInterval,

		portStatsMetric: promauto.NewGaugeVec(prometheus.GaugeOpts{
			Name: "port_stats_tx_rate",
			Help: "Port TX rate",
		}, []string{"port"}),
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

	client := intersight.NewAPIClient(config)

	im.ctx = ctx
	im.client = client

	logrus.Info("Intersight metrics setup complete, starting collector thread")
	go im.run()

	logrus.Info("Starting metrics server")
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)

}

func (im *IntersightMetrics) run() {
	for {
		im.collectPortStats()

		logrus.Infof("Sleeping for %d seconds", im.updateInterval)
		time.Sleep(time.Duration(im.updateInterval) * time.Second)
	}
}
