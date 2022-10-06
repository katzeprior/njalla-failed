package njalla

import (
	"context"
	"fmt"
)

func main() {
	p := Provider{
		APIToken: "TOKEN",
	}

	records, err := p.GetRecords(context.Background(), "domain.tld")
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		return
	}

	fmt.Println(records)
}
