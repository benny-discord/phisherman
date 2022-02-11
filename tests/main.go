package main

import (
	"fmt"
	"github.com/benny-discord/phisherman-api"
)

func main() {
	key := "a490401c-5582-4866-b3c4-99785abbf845"
	c := phisherman_api.MakeClient()
	b := phisherman_api.BulkReportDomainBody{
		key: phisherman_api.BulkReportDomainList{
			phisherman_api.SuspiciousDomain:            []int{1635591332},
			"https://" + phisherman_api.VerifiedDomain: []int{1635591333},
		},
	}
	err := c.BulkReportCaughtPhish(b, key)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("success")
	}
}
