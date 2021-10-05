package main

import (
	"context"
	"flag"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/xuzhenglun/apple-store-exporter/pkg/api"
	"github.com/xuzhenglun/apple-store-exporter/pkg/config"
	"github.com/xuzhenglun/apple-store-exporter/pkg/metric"
)

var (
	bind = flag.String("bind", ":2112", "address to bind exporter server")
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalln(err)
	}

	client := api.NewClient(cfg.Endpoint, cfg.Timeout)

	metric.StartMonitFulfillment(context.Background(), client, cfg)

	http.Handle("/metrics", promhttp.Handler())
	if err := http.ListenAndServe(*bind, nil); err != nil {
		log.Fatalln(err)
	}
}
