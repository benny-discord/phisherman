package phisherman_api

import "time"

// APIHTTPCode is used for common http codes returned by the API
type APIHTTPCode int

// Constants for APIHTTPCode
const (
	OK              APIHTTPCode = 200
	OKCreated       APIHTTPCode = 201
	OkNoContent     APIHTTPCode = 204
	NotModified     APIHTTPCode = 304
	BadRequest      APIHTTPCode = 400
	Unauthorized    APIHTTPCode = 401
	Forbidden       APIHTTPCode = 403
	TooManyRequests APIHTTPCode = 429
	InternalServer  APIHTTPCode = 500
)

// Classification is used for the classification of a phishing site
type Classification string

// Constants for Classification
const (
	Suspicious Classification = "suspicious"
	Malicious  Classification = "malicious"
	Safe       Classification = "safe"
)

// CheckDomainResponse is the response from the check domain endpoint
type CheckDomainResponse struct {
	Classification Classification `json:"classification"`
	VerifiedPhish  bool           `json:"verified_phish"`
}

// FetchDomainInfoResponse is the response from the fetch domain endpoint
type FetchDomainInfoResponse struct {
	Status         string         `json:"status"`
	LastChecked    time.Time      `json:"lastChecked"`
	VerifiedPhish  bool           `json:"verifiedPhish"`
	Classification Classification `json:"classification"`
	Created        time.Time      `json:"created"`
	FirstSeen      time.Time      `json:"firstSeen"`
	LastSeen       time.Time      `json:"lastSeen"`
	TargetedBrand  string         `json:"targetedBrand"`
	Details        Details        `json:"details"`
}

// Details are the deeper details of the domain's information
type Details struct {
	PhishCaught       int     `json:"phishCaught"`
	PhishTankID       int     `json:"phishTankId"`
	URLScanID         string  `json:"urlScanId"`
	WebsiteScreenshot string  `json:"websiteScreenshot"`
	IPAddress         string  `json:"ip_address"`
	ASN               ASN     `json:"asn"`
	Registry          string  `json:"registry"`
	Country           Country `json:"country"`
}

// ASN is the ASN information for the domain
type ASN struct {
	ASN     string `json:"asn"`
	ASNName string `json:"asn_name"`
	Route   string `json:"route"`
}

// Country is the country information for the domain
type Country struct {
	Code string `json:"code"`
	Name string `json:"name"`
}
