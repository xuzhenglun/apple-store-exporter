package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	baseURL string
	http.Client
}

func NewClient(endpoint string, timeout time.Duration) *Client {
	return &Client{
		baseURL: endpoint,
		Client:  http.Client{Timeout: timeout},
	}
}

func (c *Client) do(ctx context.Context, method, url string, params url.Values, in, out interface{}) error {
	var reqBody io.Reader
	if in != nil {
		reqBodyStr, err := json.Marshal(reqBody)
		if err != nil {
			return err
		}

		reqBody = bytes.NewReader(reqBodyStr)
	}

	reqURL := fmt.Sprintf("%s/%s?%s", c.baseURL, url, params.Encode())

	req, err := http.NewRequestWithContext(ctx, method, reqURL, reqBody)
	if err != nil {
		return err
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	respBodyStr, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode/100 != 2 {
		return fmt.Errorf("response code is not 2xx, detail: %q", string(respBodyStr))
	}

	if out != nil {
		if err := json.Unmarshal(respBodyStr, out); err != nil {
			return err
		}
	}

	return nil
}

type Response struct {
	Head struct {
		Status string      `json:"status"`
		Data   interface{} `json:"data"`
	} `json:"head"`

	Body interface{} `json:"body"`
}
