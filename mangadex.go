// Package mangadex is a client for the MangaDex API.
package mangadex

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

// Client implements a way to talk to MangaDex' API.
type Client struct {
	base, path string
	client     *http.Client
}

// An OptionFunc can be used to modify the Tapas client.
type OptionFunc func(*Client)

// WithBase sets the MangaDex base.
func WithBase(base string) OptionFunc {
	return func(md *Client) { md.base = base }
}

// WithPath replaces the default path. Might be used on a new API version.
func WithPath(path string) OptionFunc {
	return func(md *Client) { md.path = path }
}

// WithHTTPClient makes the manga client use a given http.Client to make
// requests.
func WithHTTPClient(c *http.Client) OptionFunc {
	return func(md *Client) { md.client = c }
}

// New returns a new MangaDex Client.
func New(options ...OptionFunc) *Client {
	c := &Client{
		base:   "https://mangadex.org",
		path:   "/api/v2",
		client: http.DefaultClient,
	}
	for _, option := range options {
		option(c)
	}
	return c
}

// get sends a HTTP GET request.
func (c *Client) get(ctx context.Context, path string, query url.Values) (json.RawMessage, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.base+c.path+path, nil)
	if err != nil {
		return nil, errors.Wrap(err, "could not create get request")
	}
	req.URL.RawQuery = query.Encode()
	res, err := c.client.Do(req)
	if err != nil {
		return nil, errors.Wrapf(err, "could not get %s", req.URL)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, errors.Errorf("could not get %s: %s", req.URL, res.Status)
	}
	var payload struct {
		Code    int             `json:"code"`
		Status  string          `json:"status"`
		Message string          `json:"message"`
		Data    json.RawMessage `json:"data"`
	}
	if err := json.NewDecoder(res.Body).Decode(&payload); err != nil {
		return nil, errors.Wrap(err, "could not decode response")
	}
	if payload.Code != http.StatusOK {
		return nil, errors.Errorf("could not get %s: %s", req.URL, payload.Message)
	}
	return payload.Data, nil
}
