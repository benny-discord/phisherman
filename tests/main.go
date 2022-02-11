package main

import (
	"fmt"
	"time"

	"github.com/benny-discord/phisherman-api"
)

func main() {
	key := "a490401c-5582-4866-b3c4-99785abbf845"
	c := phisherman_api.MakeClient()
	b := make(phisherman_api.BulkReportDomainBody)
	b[key] = phisherman_api.BulkReportDomainList{
		phisherman_api.SuspiciousDomain:            []int{int(time.Now().Unix())},
		"https://" + phisherman_api.VerifiedDomain: []int{int(time.Now().Unix())},
	}
	err := c.BulkReportCaughtPhish(b, key)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("success")
	}
}
