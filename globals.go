package identitymind

import (
	"fmt"
	"os"
	"sync"

	"github.com/kthomas/go-logger"
)

const identitymindDefaultEnvironment = "sandbox" // use 'edna' for production; see https://edoc.identitymind.com/reference#section-integration-environments

var (
	log           *logger.Logger
	bootstrapOnce sync.Once

	identitymindAPIBaseURL string
	identitymindAPIUser    string
	identitymindAPIToken   string
)

func init() {
	bootstrapOnce.Do(func() {
		lvl := os.Getenv("IDENTITYMIND_LOG_LEVEL")
		if lvl == "" {
			lvl = "INFO"
		}
		var endpoint *string
		if os.Getenv("SYSLOG_ENDPOINT") != "" {
			endpt := os.Getenv("SYSLOG_ENDPOINT")
			endpoint = &endpt
		}
		log = logger.NewLogger("identitymind", lvl, endpoint)

		if os.Getenv("IDENTITYMIND_API_ENVIRONMENT") != "" {
			identitymindAPIBaseURL = fmt.Sprintf("https://%s.identitymind.com", os.Getenv("IDENTITYMIND_API_ENVIRONMENT"))
		} else {
			identitymindAPIBaseURL = fmt.Sprintf("https://%s.identitymind.com", identitymindDefaultEnvironment)
		}

		if os.Getenv("IDENTITYMIND_API_USER") != "" {
			identitymindAPIUser = os.Getenv("IDENTITYMIND_API_USER")
		}

		if os.Getenv("IDENTITYMIND_API_TOKEN") != "" {
			identitymindAPIToken = os.Getenv("IDENTITYMIND_API_TOKEN")
		}
	})
}
