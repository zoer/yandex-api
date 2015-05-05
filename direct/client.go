package direct

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	defaultBaseURL = "https://api-sandbox.direct.yandex.ru/v4/json/"
	userAgent      = "API Client"
)

type Client struct {
	client    *http.Client
	BaseURL   *url.URL
	UserAgent string
	Token     string

	Campaigns *CampaignsService
}

// NewClient returns a new Yandex Direct API client.
func NewClient(token string) *Client {
	BaseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{
		client:    http.DefaultClient,
		BaseURL:   BaseURL,
		UserAgent: userAgent,
		Token:     token,
	}

	c.Campaigns = &CampaignsService{client: c}

	return c
}

func (c *Client) NewRequest(body interface{}) (*http.Request, error) {
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest("POST", c.BaseURL.String(), buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	if c.UserAgent != "" {
		req.Header.Add("User-Agent", c.UserAgent)
	}
	return req, nil
}

func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	err = checkResponse(res)
	if err != nil {
		return res, err
	}

	if v != nil {
		err = json.NewDecoder(res.Body).Decode(v)
	}

	return res, err
}

func checkResponse(res *http.Response) error {
	c := res.StatusCode
	if c >= 200 && c < 300 {
		return nil
	}
	return fmt.Errorf("HTTP response code is %d", c)
}
