package authenticator

import (
	"context"
	"net/http"

	"github.com/GbSouza15/apiToDoGo/internal/app/response"
	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

var userIdCtxKey contextKey = "user_id"

func UserIDFromContext(ctx context.Context) string {
	userId, ok := ctx.Value(userIdCtxKey).(string)
	if !ok {
		return ""
	}

	return userId
}

func CheckTokenIsValid(n http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				response.SendResponse(401, []byte("Não autorizado"), w)
				return
			}
			response.SendResponse(400, []byte("Erro no servidor"), w)
			return
		}

		tokenString := c.Value
		token, err := ValidatorToken(tokenString)

		if err != nil {
			response.SendResponse(400, []byte("Erro no servidor"), w)
			return
		}

		if !token.Valid {
			response.SendResponse(401, []byte("Não autorizado"), w)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response.SendResponse(401, []byte("Não autorizado"), w)
			return
		}

		userId := claims["user_id"].(string)

		r = r.WithContext(context.WithValue(r.Context(), userIdCtxKey, userId))

		n(w, r)
	}
}

func SendResponse(code int, data []byte, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
