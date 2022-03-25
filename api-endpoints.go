package phisherman

// APIBase is the base URL for the API.
const APIBase = "https://api.phisherman.gg/v2"

// Constants for API Routes
const (
	CheckDomainRoute      = APIBase + "/domains/check/{domain}"
	FetchDomainRoute      = APIBase + "/domains/info/{domain}"
	ReportCaughtPhish     = APIBase + "/phish/caught/{domain}"
	BulkReportCaughtPhish = APIBase + "/phish/caught/bulk"
	ReportNewPhish        = APIBase + "/phish/report"
)

// Domain constants for testing purposes
const (
	SuspiciousDomain string = "suspicious.test.phisherman.gg"
	VerifiedDomain   string = "verified.test.phisherman.gg"
	UnknownDomain    string = "unknown.test.phisherman.gg"
)

var _ = SuspiciousDomain
var _ = VerifiedDomain
var _ = UnknownDomain
