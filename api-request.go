package phisherman_api

// BulkReportDomainBody is used for the bulk report endpoint
//
// Each key is the end-user's API Key, and the value is the map of reports for that user (see: BulkReportDomainList)
type BulkReportDomainBody map[string]BulkReportDomainList

// BulkReportDomainList is a map of each domain to report, and contains an array of the timestamp of each usage
type BulkReportDomainList map[string][]int

type reportNewPhishBody struct {
	URL string `json:"url"`
}

type reportDomainBody struct {
	GuildID int `json:"guild"`
}
