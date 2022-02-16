package auth

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	log "github.com/sirupsen/logrus"
)

// A private key for context that only this package can access. This is important
// to prevent collisions between different context uses
var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

// A stand-in for our database backed user object
type User struct {
	Name    string
	IsAdmin bool
}

// Middleware decodes the share session cookie and packs the session into context
func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authToken := r.Header.Get("authorization")
			log.WithFields(log.Fields{"authorization": authToken}).Infof("Bearer Token on request: ")
			validateJWT(strings.Split(authToken, " ")[1])

			next.ServeHTTP(w, r)
		})
	}
}

type MyCustomClaims struct {
	Foo string `json:"sub"`
	jwt.StandardClaims
}

func validateJWT(inputToken string) {
	fmt.Println(inputToken)
	token, err := jwt.Parse(inputToken, func(token *jwt.Token) (interface{}, error) {
		return []byte("AllYourBase"), nil
	})
	if err != nil {
		fmt.Println(err)
	}
	claims, _ := token.Claims.(*MyCustomClaims)
	fmt.Printf("%v %v", claims.Foo, claims.StandardClaims.ExpiresAt)
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *User {
	raw, _ := ctx.Value(userCtxKey).(*User)
	return raw
}
