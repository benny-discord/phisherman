package phisherman_api

import (
	"errors"
	"net/url"
	"strings"

	"github.com/go-resty/resty/v2"
)

// Client is the base client for the Phisherman API.
type Client struct {
	client     *resty.Client
	baseAPIKey string
}

// MakeClient initialises a new Client.
func MakeClient(baseAPIKey string) (*Client, error) {
	if baseAPIKey == "" {
		return nil, errors.New("missing base API Key")
	}
	return &Client{
		client:     resty.New(),
		baseAPIKey: baseAPIKey,
	}, nil
}

// BaseAPIKey returns the base API Key of the Client.
func (c *Client) BaseAPIKey() string {
	return c.baseAPIKey
}

// CheckDomain returns the CheckDomainResponse for a given domain
func (c *Client) CheckDomain(domain string) (*CheckDomainResponse, error) {
	if domain == "" {
		return nil, errors.New("missing domain")
	}
	if strings.Index(domain, "://") == -1 {
		domain = "https://" + domain
	}
	u, err := url.Parse(domain)
	if err != nil {
		return nil, err
	}
	var r CheckDomainResponse
	_, err = c.client.R().
		SetHeader("Authorization", c.baseAPIKey).
		SetResult(&r).
		Get(strings.Replace(CheckDomainRoute, "{domain}", u.Host, 1))

	return &r, err
}

func (c *Client) FetchDomainInfo(domain string) (*FetchDomainInfoResponse, error) {
	if domain == "" {
		return nil, errors.New("missing domain")
	}
	if strings.Index(domain, "://") == -1 {
		domain = "https://" + domain
	}
	u, err := url.Parse(domain)
	if err != nil {
		return nil, err
	}
	var d map[string]FetchDomainInfoResponse
	_, err = c.client.R().
		SetHeader("Authorization", "Bearer "+c.baseAPIKey).
		SetResult(d).
		Get(strings.Replace(FetchDomainRoute, "{domain}", u.Host, 1))

	var r FetchDomainInfoResponse
	if err != nil {
		r = d[u.Host]
	}
	return &r, err
}
