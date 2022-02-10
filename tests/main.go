package main

import (
	"fmt"

	"github.com/benny-discord/phisherman-api"
)

func main() {
	c, _ := phisherman_api.MakeClient("a490401c-5582-4866-b3c4-99785abbf845")
	d, err := c.FetchDomainInfo("testing-discord.com")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("%+v\n", d)
	}
}
