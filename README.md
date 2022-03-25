# Phisherman API - Golang

## Installation
Requires Go modules to function

`go get github.com/benny-discord/phisherman`

## Usage
```go
package main

import (
    "github.com/benny-discord/phisherman"
)

func main() {
    client := phisherman.MakeClient()
}
```

#### Check Domain
[Check Domain Docs](https://docs.phisherman.gg/api/v2/check-a-domain.html)
```go
// This should be an end user API key, or any key with API.READ scope
apiKey := "MY_API_KEY"
data, err := client.CheckDomain("suspicious.test.phisherman.gg", apiKey)
if err != nil {
    fmt.Println(err)
}
fmt.Printf("%+v\n", data)
// Output:
// &{Classification:suspicious VerifiedPhish:false}
```

#### Fetch Domain Info
[Fetch Domain Info Docs](https://docs.phisherman.gg/api/v2/fetch-domain-info.html)
```go
// This should be an end user API key, or any key with API.READ scope
apiKey := "MY_API_KEY"
data, err := client.FetchDomainInfo("suspicious.test.phisherman.gg", apiKey)
if err != nil {
    fmt.Println(err)
}
fmt.Printf("%+v\n", data)
// Output:
/* &{Status:0xc00022d980 LastChecked:2021-12-29 21:42:36 +0000 UTC VerifiedPhish:0xc00000ef6f Classification:0xc00022d990 Created:2021-12-29 21:42:29 +0000 UTC FirstSeen:0001-01-01 00:00:0 0 +0000 UTC LastSeen:0001-01-01 00:00:00 +0000 UTC TargetedBrand:0xc00022d9a0 Details:0xc00004c680}*/
```

#### Reporting a Caught Phish
This is entirely optional, but is used to improve the metrics of the Phisherman API.
[More Information](https://docs.phisherman.gg/api/v2/catching-a-phish.html)
```go
// This should be your Bot's API key, or a user API key with the API.UPDATE scope
apiKey := "MY_API_KEY"
// The guild ID the event occurred in - can be nil
guildID := 1234567890123456
err := client.ReportCaughtPhish("suspicious.test.phisherman.gg", apiKey, guildID)
if err != nil {
    fmt.Println(err)
}
fmt.Println("Success")
// Output:
// Success
```

#### Reporting a New Phish
This should be done when you suspect a domain is a phishing domain, but is not listed on Phisherman (this is optional).
[More Information](https://docs.phisherman.gg/api/v2/report-a-phish.html)
```go
// This should be an end user API key, or any key with API.READ scope
apiKey := "MY_API_KEY"
// Note: the domain is a full URL here, rather than just the hostname
err := client.ReportCaughtPhish("https://suspicious.test.phisherman.gg/my-full-URL-path", apiKey, guildID)
if err != nil {
    fmt.Println(err)
}
fmt.Println("Success")
// Output:
// Success
```

#### Bulk Reporting Domains
[Docs](https://docs.phisherman.gg/api/v2/catching-a-phish.html#bulk-reporting)
```go
body := phisherman_api.BulkReportDomainBody{
    "MY_USER_API_KEY": phisherman_api.BulkReportDomainList{
        "DOMAIN_TO_REPORT":   []int{/*my int timestamp of each usage goes here*/},
        "DOMAIN_TO_REPORT_2": []int{1635591333, 1635591334, 1635591335 /*3 reported uses*/},
    },
    "MY_USER_API_KEY_2": phisherman_api.BulkReportDomainList{
        "DOMAIN_TO_REPORT_3": []int{1635591332 /* 1 reported use */},
        "DOMAIN_TO_REPORT_4": []int{1635591333, 1635591334 /* 2 reported uses */},
        //...
    },
    //...
}
// This should be a bot API key, or any key with API.UPDATE_BULK scope
apiKey := "MY_API_KEY"
err := client.BulkReportPhish(body, apiKey)
if err != nil {
    fmt.Println(err)
}
fmt.Println("Success")
```
This endpoint is intended for larger bots who would like to report several domains at once.