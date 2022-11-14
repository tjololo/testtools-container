package nettools

import (
	"fmt"
	"net"
	"net/http"
	"time"
)

// TestConnectPossible checks if it is possible to connect to server
func TestConnectPossible(host, network string, port int, timeout time.Duration) (time.Duration, error) {
	hoststring := fmt.Sprintf("%s:%d", host, port)
	start := time.Now()
	conn, err := net.DialTimeout(network, hoststring, timeout)
	elaped := time.Since(start)
	if err != nil {
		return elaped, err
	}
	err = conn.Close()
	if err != nil {
		fmt.Printf("Failed to close connection %v", err)
	}
	return elaped, nil
}

func TestDNSLookup(host string) ([]net.IP, time.Duration, error) {
	start := time.Now()
	ips, err := net.LookupIP(host)
	elapsed := time.Since(start)
	return ips, elapsed, err
}

func TestApiRequest(uri string, timeout time.Duration) (respTime time.Duration, statusCode int, err error) {
	client := http.Client{Timeout: timeout}
	start := time.Now()
	resp, err := client.Get(uri)
	respTime = time.Since(start)
	statusCode = resp.StatusCode
	return
}
