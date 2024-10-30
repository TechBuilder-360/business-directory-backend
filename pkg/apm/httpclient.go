package apm

import (
	"github.com/TechBuilder-360/business-directory-backend/pkg/apm/sentry"
	"net/http"
)

func InstrumentedRoundTripper() http.RoundTripper {
	return sentry.ClientOpt().HTTPTransport
}
