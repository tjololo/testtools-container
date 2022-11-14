package main

import (
	"fmt"
	"github.com/tjololo/testtools-container/pkg/serve"
	"os"
)

func main() {
	configfile := "/mnt/test-config.yaml"
	port := 2112
	err := serve.StartTestsAndPrometheusEnpoint(configfile, port)
	if err != nil {
		fmt.Printf("An error occurred: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Server shutdown")
	os.Exit(0)
}
