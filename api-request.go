package phisherman_api

// BulkReportDomainBody is used for the bulk report endpoint
//
// Each key is the end-user's API Key, and the value is the map of reports for that user (see: BulkReportDomainList)
type BulkReportDomainBody map[string]BulkReportDomainList

// BulkReportDomainList is a map of each domain to report, and contains an array of the timestamp of each usage
type BulkReportDomainList map[string][]int

// ReportDomainBody is used for reporting a phishing domain being sent
type ReportDomainBody struct {
	Domain string `json:"domain"`
}

// ReportNewPhishBody is used for reporting a new phishing domain being located
type ReportNewPhishBody struct {
	URL string `json:"url"`
}
