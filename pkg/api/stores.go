package api

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/xuzhenglun/apple-store-exporter/pkg/models"
)

func (c *Client) ListStores(ctx context.Context) (*models.Stores, error) {
	res := models.Stores{}
	err := c.do(ctx, http.MethodGet, "CN/zh_CN/reserve/A/stores.json", nil, nil, &res)
	return &res, err
}

func (c *Client) StoreFulfillmentMessages(ctx context.Context, shop, model string) (*models.StorePartsAvailability, error) {
	fulfillmentMessages := models.FulfillmentMessages{}
	responseStruct := Response{Body: &fulfillmentMessages}

	params := url.Values{
		"pl":      []string{"true"},
		"mt":      []string{"compact"},
		"parts.0": []string{model},
		"store":   []string{shop},
	}

	if err := c.do(ctx, http.MethodGet, "shop/fulfillment-messages", params, nil, &responseStruct); err != nil {
		return nil, err
	}

	if responseStruct.Head.Status != "200" {
		return nil, fmt.Errorf("response status is not ok, detail: %+v", responseStruct.Head.Data)
	}

	res := fulfillmentMessages.Content.PickupMessage.Stores[0].PartsAvailability[model]
	return &res, nil
}
