package serve

import (
	"context"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/tjololo/testtools-container/cmd/structs"
	"github.com/tjololo/testtools-container/pkg/nettools"
	"strconv"
	"time"
)

var (
	opsDnsTest = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:      "dns_test_duration",
		Help:      "DNS lookup duration",
		Namespace: "network_test",
	}, []string{"success", "name"})
	opsConTest = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:      "con_test_duration",
		Help:      "Establish connection duration",
		Namespace: "network_test",
	}, []string{"success", "name"})
	opsApiTest = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:      "api_test_duration",
		Help:      "Responsetime api request",
		Namespace: "network_test",
	}, []string{"success", "status_code", "name"})
)

func StartTestsInParalell(ctx context.Context, config structs.TestsConfig) {
	for name, dnsTest := range config.DnsTests {
		go doLookup(ctx, name, dnsTest)
	}
	for name, connectTest := range config.ConnectTests {
		go doConnect(ctx, name, connectTest)
	}

	for name, apiTest := range config.ApiTests {
		go doAPI(ctx, name, apiTest)
	}
}

func doLookup(ctx context.Context, name string, test structs.DnsTest) {
	fmt.Printf("Starting dns test %s agains host %s every %s\n", name, test.Hostname, test.Interval)
	ticker := time.NewTicker(*test.Interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Stopping test %s\n", name)
			return
		case <-ticker.C:
			_, elapsedTime, err := nettools.TestDNSLookup(test.Hostname)
			opsDnsTest.With(map[string]string{
				"success": strconv.FormatBool(err == nil),
				"name":    name,
			}).Observe(elapsedTime.Seconds())
		}
	}
}

func doConnect(ctx context.Context, name string, test structs.ConnectTest) {
	fmt.Printf("Starting connect test %s agains host %s:%d every %s\n", name, test.Hostname, test.Port, test.Interval)
	ticker := time.NewTicker(*test.Interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Stopping test %s\n", name)
			return
		case <-ticker.C:
			elapsedTime, err := nettools.TestConnectPossible(test.Hostname, "tcp", *test.Port, *test.Timeout)
			opsConTest.With(map[string]string{
				"success": strconv.FormatBool(err == nil),
				"name":    name,
			}).Observe(elapsedTime.Seconds())
		}
	}
}

func doAPI(ctx context.Context, name string, test structs.ApiTest) {
	fmt.Printf("Starting connect test %s agains URI %s every %s\n", name, test.URI, test.Interval)
	ticker := time.NewTicker(*test.Interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Stopping test %s\n", name)
			return
		case <-ticker.C:
			elapsedTime, statuscode, err := nettools.TestApiRequest(test.URI, *test.Timeout)
			opsApiTest.With(map[string]string{
				"success":     strconv.FormatBool(err == nil),
				"status_code": strconv.Itoa(statuscode),
				"name":        name,
			}).Observe(elapsedTime.Seconds())
		}
	}
}
