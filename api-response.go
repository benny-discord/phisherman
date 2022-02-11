package phisherman

import "time"

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

// FetchDomainInfoResponse is the body of the fetch domain info endpoint
type FetchDomainInfoResponse struct {
	Status         *string         `json:"status"`
	LastChecked    *Time           `json:"lastChecked"`
	VerifiedPhish  *bool           `json:"verifiedPhish"`
	Classification *Classification `json:"classification"`
	Created        *Time           `json:"created"`
	FirstSeen      *Time           `json:"firstSeen"`
	LastSeen       *Time           `json:"lastSeen"`
	TargetedBrand  *string         `json:"targetedBrand"`
	Details        *Details        `json:"details"`
}

// Details are the deeper details of the domain's information
type Details struct {
	PhishCaught       *int     `json:"phishCaught"`
	PhishTankID       *string  `json:"phishTankId"`
	URLScanID         *string  `json:"urlScanId"`
	WebsiteScreenshot *string  `json:"websiteScreenshot"`
	IPAddress         *string  `json:"ip_address"`
	ASN               *ASN     `json:"asn"`
	Registry          *string  `json:"registry"`
	Country           *Country `json:"country"`
}

// ASN is the ASN information for the domain
type ASN struct {
	ASN     *string `json:"asn"`
	ASNName *string `json:"asn_name"`
	Route   *string `json:"route"`
}

// Country is the country information for the domain
type Country struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

// Time is used for handling the time formats returned by the Phisherman API
type Time struct {
	time.Time
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (m *Time) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	if string(data) == "null" || string(data) == `""` {
		return nil
	}
	// Fractional seconds are handled implicitly by Parse.
	tt, err := time.Parse(`"`+time.RFC3339+`"`, string(data))
	*m = Time{tt}
	return err
}
