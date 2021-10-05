package metric

import (
	"context"
	"log"
	"time"

	"github.com/xuzhenglun/apple-store-exporter/pkg/api"
	"github.com/xuzhenglun/apple-store-exporter/pkg/config"
)

func MonitorFulfillmentFunc(ctx context.Context, client *api.Client, shops, models []string) {
	for _, shop := range shops {
		for _, model := range models {
			res, err := client.StoreFulfillmentMessages(ctx, shop, model)
			if err != nil {
				log.Printf("fetch fulfillment for %s in %s failed, detail: %s", model, shop, err)
				continue
			}

			var number float64 = 0
			if res.StoreSelectionEnabled {
				number = 1
			}

			stockMetricMetric.WithLabelValues(shop, model).Set(number)
		}
	}
}

func StartMonitFulfillment(ctx context.Context, client *api.Client, config *config.Config) {
	timer := time.NewTimer(0)

	go func() {
		for {
			select {
			case <-timer.C:
				MonitorFulfillmentFunc(ctx, client, config.Shops, config.Models)
				timer.Reset(config.Interval)
			case <-ctx.Done():
				return
			}
		}
	}()
}
