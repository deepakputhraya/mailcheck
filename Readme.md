# Mailcheck
Go library to verify, detect invalid, spam and junk email id.              


#### Features:
- Identify disposable and free email providers
- The API also automatically checks for role-based emails (such as sales@domain.com or team@domain.com),

## Usage

```shell script
go get github.com/deepakputhraya/mailcheck
```

## Example

```go
package main

import (
	"fmt"
	"github.com/deepakputhraya/mailcheck"
)

var emails = []string{"elon@tesla.com", "elon@gmail.com", "hello@mailinator.com"}

func main() {
	for _, email := range emails {
		// Skipped error handling
		details, _ := mailcheck.GetEmailDetails(email)
		fmt.Println(email)
		fmt.Printf("Valid : %v; Disposable : %v; Free : %v; Role Based : %v\n",
			details.IsValid,
			details.IsDisposable,
			details.IsFree,
			details.IsRoleBased)
		fmt.Println("-----")
	}
}
```