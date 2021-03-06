package app_test

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/cloudfoundry/metric-store-release/src/cmd/metric-store/app"
	"github.com/cloudfoundry/metric-store-release/src/internal/metrics"
	"github.com/cloudfoundry/metric-store-release/src/internal/testing"
	"github.com/cloudfoundry/metric-store-release/src/internal/tls"
	"github.com/cloudfoundry/metric-store-release/src/pkg/logger"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Metric Store App", func() {
	var (
		metricStore *app.MetricStoreApp
	)

	BeforeEach(func() {
		metricStore = app.NewMetricStoreApp(&app.Config{
			TLS: tls.TLS{
				CAPath:   testing.Cert("metric-store-ca.crt"),
				CertPath: testing.Cert("metric-store.crt"),
				KeyPath:  testing.Cert("metric-store.key"),
			},
			MetricStoreServerTLS: app.MetricStoreServerTLS{
				CAPath:   testing.Cert("metric-store-ca.crt"),
				CertPath: testing.Cert("metric-store.crt"),
				KeyPath:  testing.Cert("metric-store.key"),
			},
			MetricStoreInternodeTLS: app.MetricStoreInternodeTLS{
				CAPath:   testing.Cert("metric-store-ca.crt"),
				CertPath: testing.Cert("metric-store.crt"),
				KeyPath:  testing.Cert("metric-store.key"),
			},
			MetricStoreMetricsTLS: app.MetricStoreMetricsTLS{
				CAPath:   testing.Cert("metric-store-ca.crt"),
				CertPath: testing.Cert("metric-store.crt"),
				KeyPath:  testing.Cert("metric-store.key"),
			},
			StoragePath:          "/tmp/metric-store",
			ReplicationFactor:    1,
			NodeAddrs:            []string{"localhost:8080"},
			MaxConcurrentQueries: 20,
		}, logger.NewTestLogger(GinkgoWriter))
		go metricStore.Run()

		Eventually(metricStore.ProfilingAddr, 15).ShouldNot(BeEmpty())
	})

	AfterEach(func() {
		metricStore.Stop()
		metricStore = nil
	})

	It("serves metrics on a metrics endpoint", func() {
		var body string

		tlsConfig, err := tls.NewMutualTLSClientConfig(
			testing.Cert("metric-store-ca.crt"),
			testing.Cert("metric-store.crt"),
			testing.Cert("metric-store.key"),
			"metric-store",
		)
		Expect(err).ToNot(HaveOccurred())

		httpClient := &http.Client{
			Transport: &http.Transport{TLSClientConfig: tlsConfig},
		}

		fn := func() string {
			resp, err := httpClient.Get("https://" + metricStore.MetricsAddr() + "/metrics")
			if err != nil {
				return ""
			}
			defer func() { _ = resp.Body.Close() }()

			bytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return ""
			}

			body = string(bytes)

			return body
		}
		Eventually(fn).ShouldNot(BeEmpty())
		Expect(body).To(ContainSubstring(metrics.MetricStoreWrittenPointsTotal))
		Expect(body).To(ContainSubstring("go_threads"))
	})

	It("listens with pprof", func() {
		callPprof := func() int {

			resp, err := http.Get("http://" + metricStore.ProfilingAddr() + "/debug/pprof")
			if err != nil {
				fmt.Printf("calling pprof: %s\n", err)
				return -1
			}
			defer func() { _ = resp.Body.Close() }()

			return resp.StatusCode
		}
		Eventually(callPprof).Should(Equal(200))
	})
})
