package middlewares

import (
	"errors"
	"github.com/TechBuilder-360/business-directory-backend/internal/model"
	"net/http"
)

// AuthorizeOrganisationJWT handles organisation jwt validation
//func AuthorizeOrganisationJWT(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(ctx *fiber.Ctx) error {
//		var ctx context.Context
//		publicKey := extractOrganisationToken(r)
//		organisation, err := services.NewOrganisationService().GetOrganisationByPublicKey(publicKey)
//		if organisation != nil {
//			ctx = context.WithValue(r.Context(), AuthOrganisationContextKey, organisation)
//		} else {
//			log.Error(err)
//			w.WriteHeader(http.StatusUnauthorized)
//			json.NewEncoder(w).Encode(utils.ErrorResponse{
//				Status:  false,
//				Message: "unauthorized",
//			})
//			return
//		}
//
//		next.ServeHTTP(w, r.WithContext(ctx))
//	})
//}

func extractOrganisationToken(r *http.Request) string {
	const BearerSchema = "Bearer"
	authHeader := r.Header.Get("X-Token")
	publicKey := authHeader[len(BearerSchema)+1:]
	return publicKey
}

func OrganisationFromContext(r *http.Request) (*model.Organisation, error) {
	org := r.Context().Value(AuthOrganisationContextKey)

	if org == nil {
		return nil, errors.New("no organisation in context")
	}

	organisation := org.(*model.Organisation)

	return organisation, nil
}
