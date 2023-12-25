package constant

import "github.com/TechBuilder-360/business-directory-backend/internal/common/types"

const (
	// RequestIdentifier is the name of the request ID header
	RequestIdentifier   = "Request-Id"
	InternalServerError = "internal server error"

	Directory types.Directory = "Directory"

	Verified   types.VerificationType = "VERIFIED"
	Partial    types.VerificationType = "PARTIAL"
	Unverified types.VerificationType = "UNVERIFIED"

	OnSite types.LocationType = "ON-SITE"
	Remote types.LocationType = "REMOTE"
)
