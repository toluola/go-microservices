package main

import (
	"flag"
)

func main() {
	listenAddr := flag.String("listenAddr", ":3000", "The service is running")
	flag.Parse()

	svc := NewLoggingService(NewMetricService(&priceFetcher{}))

	server := NewJSONAPIServer(*listenAddr, svc)
	server.Run()
}
