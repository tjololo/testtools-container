package serve

import (
	"context"
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func StartTestsAndPrometheusEnpoint(configfile string, promPort int) error {
	conf, err := ReadConfigFile(configfile)
	if err != nil {
		return fmt.Errorf("Unable to read configfile %s, error: %v\n", configfile, err)
	}
	fmt.Printf("Config: %+v\n", *conf)
	http.Handle("/metrics", promhttp.Handler())
	server := http.Server{Addr: fmt.Sprintf(":%d", promPort)}
	go func() {
		log.Fatal(server.ListenAndServe())
	}()

	ctx, cancel := context.WithCancel(context.Background())
	StartTestsInParalell(ctx, *conf)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	sig := <-signalChan
	cancel()
	fmt.Printf("Signal %s received, shuttinfdown...\n", sig)
	err = server.Shutdown(context.Background())
	if err != nil {
		return fmt.Errorf("Failed to shutdown server %v\n", err)
	}
	return nil
}
