package constant

import "github.com/TechBuilder-360/business-directory-backend/internal/common/types"

const (
	// RequestIdentifier is the name of the request ID header
	RequestIdentifier   = "Request-Id"
	InternalServerError = "internal server error"

	Directory types.Directory = "Directory"

	AuthToken types.Hash = "Auth-Token"
)
