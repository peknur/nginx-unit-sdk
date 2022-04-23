package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/peknur/nginx-unit-sdk/unit"
)

type ResponseError struct {
	StatusCode int
	StatusText string
	URL        string
	Text       string `json:"error,omitempty"`
	Detail     string `json:"detail,omitempty"`
}

func (e *ResponseError) Error() string {
	return fmt.Sprintf("got unexpected response code %d (%s) from %s [%s: %s]", e.StatusCode, e.StatusText, e.URL, e.Text, e.Detail)
}

type Client struct {
	baseURL    *url.URL
	httpClient *http.Client
}

var _ unit.Client = (*Client)(nil)

// Get returns the entity at the request URI.
func (c *Client) Get(ctx context.Context, path string, v interface{}) error {
	body, err := c.request(ctx, http.MethodGet, c.entityURL(path), nil)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, v)
}

// Put replaces the entity at the request URI.
func (c *Client) Put(ctx context.Context, path string, v interface{}) error {
	js, err := json.Marshal(v)
	if err != nil {
		return err
	}
	_, err = c.request(ctx, http.MethodPut, c.entityURL(path), bytes.NewReader(js))
	return err
}

// PutBinary replaces the entity at the request URI with data.
func (c *Client) PutBinary(ctx context.Context, path string, data []byte) error {
	_, err := c.request(ctx, http.MethodPut, c.entityURL(path), bytes.NewReader(data))
	return err
}

// Post updates the array at the request URI.
func (c *Client) Post(ctx context.Context, path string, v interface{}) error {
	js, err := json.Marshal(v)
	if err != nil {
		return err
	}
	_, err = c.request(ctx, http.MethodPost, c.entityURL(path), bytes.NewReader(js))
	return err
}

// Delete deletes the entity at the request URI.
func (c *Client) Delete(ctx context.Context, path string) error {
	_, err := c.request(ctx, http.MethodDelete, c.entityURL(path), nil)
	return err
}

func (c *Client) request(ctx context.Context, method string, uri *url.URL, payload io.Reader) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, method, uri.String(), payload)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return body, err
	}
	if resp.StatusCode != http.StatusOK {
		return body, newResponseError(uri.String(), resp, body)
	}
	return body, nil
}

func (c *Client) entityURL(path string) *url.URL {
	u := *c.baseURL
	u.Path = fmt.Sprintf("%s/%s", c.baseURL.Path, path)
	return &u
}

func New(URL string, httpClient *http.Client) (*Client, error) {
	u, err := url.Parse(strings.TrimSuffix(URL, "/"))
	if err != nil {
		return nil, err
	}
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	return &Client{baseURL: u, httpClient: httpClient}, nil
}

func NewClient(URL string) (*Client, error) {
	u, err := url.Parse(strings.TrimSuffix(URL, "/"))
	if err != nil {
		return nil, err
	}
	return &Client{baseURL: u, httpClient: http.DefaultClient}, nil
}

func newResponseError(URL string, r *http.Response, body []byte) error {
	err := ResponseError{
		StatusCode: r.StatusCode,
		StatusText: r.Status,
		URL:        URL,
	}
	if e := json.Unmarshal(body, &err); e != nil {
		err.Text = e.Error()
	}
	return &err
}
