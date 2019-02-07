package function

import (
	"fmt"
)

// Handle a serverless request
func Handle(req []byte) string {
	return fmt.Sprintf("Hello, raspberry pi. Commit 13. Checking log streaming.")
}
