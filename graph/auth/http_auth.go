package auth

import (
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"math/big"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/golang-jwt/jwt"
)

var rsakeys map[string]*rsa.PublicKey

// Middleware decodes the share session cookie and packs the session into context
func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenString := r.Header.Get("Authorization")
			isValid := false
			errorMessage := ""

			if strings.HasPrefix(tokenString, "Bearer ") {
				tokenString = strings.TrimPrefix(tokenString, "Bearer ")
				token, err := jwt.ParseWithClaims(tokenString, &EveCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
					return rsakeys[token.Header["kid"].(string)], nil
				})
				if err != nil {
					errorMessage = err.Error()
				} else if !token.Valid {
					errorMessage = "Invalid token"
				} else if token.Header["alg"] == nil {
					errorMessage = "alg must be defined"
				} else {
					isValid = true
				}
				claims, ok := token.Claims.(*EveCustomClaims)
				if ok && isValid {
					log.WithFields(log.Fields{"audience": claims.Audience, "expiration": claims.ExpiresAt, "id": claims.Id, "issuedAt": claims.IssuedAt, "issuer": claims.Issuer, "notBefore": claims.NotBefore, "subject": claims.Subject, "scopes": claims.Scopes}).Info("JWT recieved and decoded.")
				} else {
					log.Errorf("Invalid jwt: %s", errorMessage)
				}
			}
			next.ServeHTTP(w, r)
		})
	}
}

type EveCustomClaims struct {
	Scopes          []string `json:"scp"`
	JWTID           string   `json:"jti"`
	AuthorizedParty string   `json:"azp"`
	Tenant          string   `json:"tenat"`
	Tier            string   `json:"tier"`
	Region          string   `json:"region"`
	Name            string   `json:"name"`
	Owner           string   `json:"owner"`
	jwt.StandardClaims
}

func GetPublicKeys() {
	rsakeys = make(map[string]*rsa.PublicKey)
	var body map[string]interface{}
	uri := "https://login.eveonline.com/oauth/jwks"
	resp, _ := http.Get(uri)
	json.NewDecoder(resp.Body).Decode(&body)
	for _, bodykey := range body["keys"].([]interface{}) {
		key := bodykey.(map[string]interface{})
		if key["alg"].(string) == "RS256" {
			kid := key["kid"].(string)
			rsakey := new(rsa.PublicKey)
			number, _ := base64.RawURLEncoding.DecodeString(key["n"].(string))
			rsakey.N = new(big.Int).SetBytes(number)
			rsakey.E = 65537
			rsakeys[kid] = rsakey
		}
	}
}
