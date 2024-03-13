package middleware

import (
	"net/http"
	"strconv"
	"strings"

	pkgctx "github.com/derangga/shopifyx/internal/pkg/context"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

const (
	AuthorizationHeader = "Authorization"
	AuthorizationPrefix = "Bearer "
)

type JWTAuth struct {
	SigningKey string
}

func ProvideJWTAuth(secret string) *JWTAuth {
	return &JWTAuth{
		SigningKey: secret,
	}
}

func (m *JWTAuth) ToMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Get the Authorization header
			authHeader := c.Request().Header.Get(AuthorizationHeader)

			// Check if the Authorization header is present and starts with "Bearer"
			if authHeader == "" || !strings.HasPrefix(authHeader, AuthorizationPrefix) {
				return echo.NewHTTPError(http.StatusUnauthorized, "missing or invalid authorization header")
			}

			// Extract the token from the header
			tokenString := strings.TrimPrefix(authHeader, AuthorizationPrefix)

			// Parse the token
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				// Check the signing method
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, echo.NewHTTPError(http.StatusUnauthorized, "invalid token signing method")
				}

				// Return the secret key
				return []byte(m.SigningKey), nil
			})
			if err != nil || !token.Valid {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
			}

			sub, _ := token.Claims.GetSubject()
			userID, _ := strconv.Atoi(sub)

			newCtx := pkgctx.SetUserIDContext(c.Request().Context(), userID)
			c.SetRequest(c.Request().WithContext(newCtx))

			// Call the next handler in the chain
			return next(c)
		}
	}
}
