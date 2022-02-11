package phisherman_api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/go-resty/resty/v2"
)

// Client is the base client for the Phisherman API.
type Client struct {
	client *resty.Client
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
func (c *Client) CheckDomain(domain string, endUserToken string) (*CheckDomainResponse, error) {
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
		SetHeader("Authorization", "Bearer "+endUserToken).
		SetResult(&r).
		Get(strings.Replace(CheckDomainRoute, "{domain}", u.Host, 1)))

	return &r, err
}

// FetchDomainInfo returns the FetchDomainInfoResponse for a given domain
func (c *Client) FetchDomainInfo(domain string, endUserToken string) (*FetchDomainInfoResponse, error) {
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
	if _, err = processResponse(c.client.R().
		SetHeader("Authorization", "Bearer "+endUserToken).
		SetResult(&d).
		Get(strings.Replace(FetchDomainRoute, "{domain}", u.Host, 1))); err != nil {
		return nil, err
	}
	r := d[u.Host]
	return &r, err
}

// ReportCaughtPhish allows you to submit metrics for a given domain. This is not required for the API to work, but is
// useful for tracking metrics.
func (c *Client) ReportCaughtPhish(domain string, botAPIToken string, guildID *int) error {
	if domain == "" {
		return errors.New("missing domain")
	}
	if strings.Index(domain, "://") == -1 {
		domain = "https://" + domain
	}
	u, err := url.Parse(domain)
	if err != nil {
		return err
	}
	req := c.client.R().
		SetHeader("Authorization", "Bearer "+botAPIToken)
	if guildID != nil {
		req.SetHeader("Content-Type", "application/json").
			SetBody(reportDomainBody{GuildID: *guildID})
	}

	if _, err := processResponse(req.Post(strings.Replace(ReportCaughtPhish, "{domain}", u.Host, 1))); err != nil {
		return err
	}
	return err
}

// BulkReportCaughtPhish allows you to report several domains at once, across several guilds.
// This is intended for larger bots who do not want to post each metric individually.
// See: https://docs.phisherman.gg/api/v2/catching-a-phish.html#bulk-reporting
func (c *Client) BulkReportCaughtPhish(body BulkReportDomainBody, botAPIToken string) error {
	for key := range body {
		for domain := range body[key] {
			if strings.Index(domain, "://") == -1 {
				domain = "https://" + domain
			}
			u, err := url.Parse(domain)
			if err != nil {
				return err
			}
			body[key][u.Host] = body[key][domain]
			delete(body[key], domain)
		}
	}
	fmt.Println(body)
	d, err1 := json.Marshal(body)
	fmt.Println(string(d))
	fmt.Println(err1)
	_, err := processResponse(c.client.R().
		SetHeader("Authorization", "Bearer "+botAPIToken).
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(BulkReportCaughtPhish),
	)
	return err
}

func (c *Client) ReportNewPhish(domain string, endUserToken string) error {
	if domain == "" {
		return errors.New("missing domain")
	}
	if strings.Index(domain, "://") == -1 {
		domain = "https://" + domain
	}
	u, err := url.Parse(domain)
	if err != nil {
		return err
	}
	_, err = processResponse(c.client.R().
		SetHeader("Authorization", "Bearer "+endUserToken).
		SetHeader("Content-Type", "application/json").
		SetBody(reportNewPhishBody{URL: u.String()}).
		Put(ReportNewPhish),
	)
	return err
}
