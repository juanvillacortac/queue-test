package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/juanvillacortac/bank-queue/pkg/models"
	"github.com/juanvillacortac/bank-queue/pkg/utils"
)

var UserContextKey int
var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

type UserClaims struct {
	jwt.StandardClaims
	ID    int64  `json:"id"`
	Email string `json:"email"`
}

func (claims *UserClaims) Refresh() {
	claims.ExpiresAt = time.Now().Add(time.Hour * 1).Unix()
}

func GetUserFromContext(ctx context.Context) (models.User, bool) {
	user, ok := ctx.Value(&UserContextKey).(models.User)
	return user, ok
}

func NewUserClaims(user models.User) UserClaims {
	return UserClaims{
		ID:    user.ID,
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		},
	}
}

func GenerateJWT(payload jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func AuthMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		cookie, _ := r.Cookie("token")
		authHeader := r.Header.Get("Authorization")
		var tokenString string
		if cookie != nil {
			tokenString = cookie.Value
		}
		if authHeader != "" {
			tokenString = strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer"))
		}
		if tokenString == "" {
			http.Error(w, "Authorization header or X-Token cookie header are required", http.StatusUnauthorized)
			return
		}
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return jwtSecret, nil
		})
		if err != nil || !token.Valid {
			if err != nil {
				log.Println(err)
			}
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			var user UserClaims
			if err := utils.UnmarshalMap(claims, &user); err != nil {
				log.Print(err)
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}
			user.Refresh()
			token, _ := GenerateJWT(user)
			http.SetCookie(w, &http.Cookie{
				Name:     "token",
				Value:    token,
				HttpOnly: true,
				Secure:   true,
				Path:     "/",
				Expires:  time.Unix(user.ExpiresAt, 0),
			})
			w.Header().Set("x-token", token)
			ctx = context.WithValue(ctx, &UserContextKey, models.User{
				ID:    user.ID,
				Email: user.Email,
			})
		} else {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}
