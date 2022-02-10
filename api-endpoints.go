package phisherman_api

// APIBase is the base URL for the API.
const APIBase = "https://api.phisherman.gg/v2"

// Constants for API Routes
const (
	CheckDomainRoute      = APIBase + "/domains/check/:domain"
	FetchDomainRoute      = APIBase + "/domains/info/:domain"
	ReportCaughtPhish     = APIBase + "/phish/caught/:domain"
	BulkReportCaughtPhish = APIBase + "/phish/caught/bulk"
	ReportNewPhish        = APIBase + "/phish/report"
)
