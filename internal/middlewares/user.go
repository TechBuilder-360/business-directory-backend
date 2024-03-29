package middlewares

import (
	"errors"
	"github.com/TechBuilder-360/business-directory-backend/internal/model"
	"github.com/gofiber/fiber/v2"
)

// AuthorizeUserJWT authorise user JWT
//func AuthorizeUserJWT() Adapter {
//
//	return func(next http.Handler) http.Handler {
//		return http.HandlerFunc(func(ctx *fiber.Ctx) error {
//
//			var user *model.User
//			var ctx context.Context
//			tokenString := ExtractBearerToken(r)
//			if tokenString == "" {
//				w.WriteHeader(http.StatusUnauthorized)
//				json.NewEncoder(w).Encode(utils.ErrorResponse{
//					Status:  false,
//					Message: "missing authentication token",
//				})
//				return
//			}
//			token, err := services.NewAuthService().ValidateToken(tokenString)
//			if err != nil {
//				log.Error(err)
//				w.WriteHeader(http.StatusUnauthorized)
//				json.NewEncoder(w).Encode(utils.ErrorResponse{
//					Status:  false,
//					Message: "authentication failed",
//				})
//				return
//			}
//			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
//				userId := claims["user_id"].(string)
//				user, err = repository.NewUserRepository().GetUserByID(userId)
//				if err != nil {
//					log.Error(err)
//					w.WriteHeader(http.StatusUnauthorized)
//					json.NewEncoder(w).Encode(utils.ErrorResponse{
//						Status:  false,
//						Message: "account not found",
//					})
//					return
//				}
//
//				ctx = context.WithValue(r.Context(), AuthUserContextKey, user)
//
//			} else {
//				log.Error(err)
//				w.WriteHeader(http.StatusUnauthorized)
//				json.NewEncoder(w).Encode(utils.ErrorResponse{
//					Status:  false,
//					Message: "unauthorized",
//				})
//				return
//			}
//
//			// Serve the next handler
//			next.ServeHTTP(w, r.WithContext(ctx))
//		})
//	}
//}

func ExtractBearerToken(r *fiber.Ctx) string {
	const BearerSchema = "Bearer"
	authHeader := r.Get("Authorization")
	if authHeader == "" {
		return ""
	}
	tokenString := authHeader[len(BearerSchema)+1:]
	return tokenString
}

func UserFromContext(r *fiber.Ctx) (*model.User, error) {
	u := r.Context().Value(AuthUserContextKey)

	if u == nil {
		return nil, errors.New("no user in context")
	}

	user := u.(*model.User)

	return user, nil
}
