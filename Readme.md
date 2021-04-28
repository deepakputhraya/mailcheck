# Mailcheck
Go library to verify, detect invalid, spam and junk email id.              


#### Features:
- Identify disposable and free email providers
- The API also automatically checks for role-based emails (such as sales@domain.com or team@domain.com),

## Usage

```go
package main

import (
	"fmt"
	mailcheck "mailcheck/libs"
)

func main() {
	for _, email := range []string{"elonmusk@tesla.com", "invalid", "elon@gmail.com", "hello@mailinator.com"} {
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