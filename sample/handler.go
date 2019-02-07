package function

import (
	"fmt"
)

// Handle a serverless request
func Handle(req []byte) string {
	return fmt.Sprintf("Hello, raspberry pi. Commit 15. Should update in dashboard.")
}
