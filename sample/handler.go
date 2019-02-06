package function

import (
	"fmt"
)

// Handle a serverless request
func Handle(req []byte) string {
	return fmt.Sprintf("Hello, raspberry p. This is the second push. Fingers crossed.")
}
