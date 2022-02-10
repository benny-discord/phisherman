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
func MakeClient() *Client {
	return &Client{
		client: resty.New(),
	}
}

func processResponse(resp *resty.Response, err error) (*resty.Response, error) {
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, errors.New("unexpected response code (" + resp.Status() + "), body:" + resp.String())
	}

	return resp, nil
}

// CheckDomain returns the CheckDomainResponse for a given domain
func (c *Client) CheckDomain(domain string, token string) (*CheckDomainResponse, error) {
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
	_, err = processResponse(c.client.R().
		SetHeader("Authorization", "Bearer "+token).
		SetResult(&r).
		Get(strings.Replace(CheckDomainRoute, "{domain}", u.Host, 1)))

	return &r, err
}

func (c *Client) FetchDomainInfo(domain string, token string) (*FetchDomainInfoResponse, error) {
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
	_, err = processResponse(c.client.R().
		SetHeader("Authorization", "Bearer "+token).
		SetResult(&d).
		Get(strings.Replace(FetchDomainRoute, "{domain}", u.Host, 1)))
	if err != nil {
		return nil, err
	}

	r := d[u.Host]
	return &r, err
}
