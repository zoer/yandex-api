## Usage

```go
package main

import (
	"github.com/zoer/yandex-api/direct"
)

func main() {
	token := os.Getenv("YANDEX_API_TOKEN")
	client := direct.NewClient(token)

	id := 123

	// Campaigns list
	campaigns, err := client.Campaigns.GetList()

	// Campaign archive
	err := client.Campaigns.Archive(id)

	// Campaign unarchive
	err := client.Campaigns.UnArchive(id)

	// Campaign stop
	err := client.Campaigns.Stop(id)

	// Campaign resume
	err := client.Campaigns.Resume(id)

	// Campaign delete
	err := client.Campaigns.Delete(id)

	// Campaign create or update
	c := CampaignParams{
		Login:                      "my-login",
		Name:                       "Eric",
		FIO:                        "Eric Cartman",
		StartDate:                  "2012-04-12",
		//...
	}
	id, err := client.Campaigns.CreateOrUpdate(&c)
```


## Tests

```
$ go test -v ./...
```
