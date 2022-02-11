package phisherman

import (
	"errors"
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

// CheckDomain returns the CheckDomainResponse for a given domain.
// The domain should not include a http(s):// prefix, or a path.
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
// The domain should not include a http(s):// prefix, or a path.
func (c *Client) FetchDomainInfo(domain string, endUserToken string) (*FetchDomainInfoResponse, error) {
	if domain == "" {
		return nil, errors.New("missing domain")
	}
	var d map[string]FetchDomainInfoResponse
	if _, err := processResponse(c.client.R().
		SetHeader("Authorization", "Bearer "+endUserToken).
		SetResult(&d).
		Get(strings.Replace(FetchDomainRoute, "{domain}", domain, 1))); err != nil {
		return nil, err
	}
	r := d[domain]
	return &r, nil
}

// ReportCaughtPhish allows you to submit metrics for a given domain. This is not required for the API to work, but is
// useful for tracking metrics.
// The domain should not include a http(s):// prefix, or a path.
func (c *Client) ReportCaughtPhish(domain string, botAPIToken string, guildID *int) error {
	if domain == "" {
		return errors.New("missing domain")
	}
	req := c.client.R().
		SetHeader("Authorization", "Bearer "+botAPIToken)
	if guildID != nil {
		req.SetHeader("Content-Type", "application/json").
			SetBody(reportDomainBody{GuildID: *guildID})
	}

	_, err := processResponse(req.Post(strings.Replace(ReportCaughtPhish, "{domain}", domain, 1)))
	return err
}

// BulkReportCaughtPhish allows you to report several domains at once, across several guilds.
// This is intended for larger bots who do not want to post each metric individually.
// The domains should not include a http(s):// prefix, or a path.
// See: https://docs.phisherman.gg/api/v2/catching-a-phish.html#bulk-reporting
func (c *Client) BulkReportCaughtPhish(body BulkReportDomainBody, botAPIToken string) error {
	_, err := processResponse(c.client.R().
		SetHeader("Authorization", "Bearer "+botAPIToken).
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(BulkReportCaughtPhish),
	)
	return err
}

// ReportNewPhish allows you to submit metrics for a given domain. This is not required for the API to work, but
//is useful for tracking metrics.
// The domain should be a full URL, including the protocol and path
func (c *Client) ReportNewPhish(domain string, endUserToken string) error {
	if domain == "" {
		return errors.New("missing domain")
	}
	if strings.Index(domain, "://") == -1 {
		domain = "https://" + domain
	}
	_, err := processResponse(c.client.R().
		SetHeader("Authorization", "Bearer "+endUserToken).
		SetHeader("Content-Type", "application/json").
		SetBody(reportNewPhishBody{URL: domain}).
		Put(ReportNewPhish),
	)
	return err
}
