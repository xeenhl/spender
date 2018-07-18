package authentication

import (
	"context"
	"crypto/rsa"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

type JWTAuthenticationSettings struct {
	privateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

type UserIdKey string

const (
	tokenDuration = 72
	expireOffset  = 3600
	privateKey    = "./authentication/keys/private_key"
	publicKey     = "./authentication/keys/public_key.pub"
	//Key to get user id populated from token in AuhtWithToken
	UserID = UserIdKey("userId")
)

var authenticationSettings *JWTAuthenticationSettings

func AuhtWithToken(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	auth := initJWTAuthenticationSettings()

	token, err := request.ParseFromRequest(r, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return auth.PublicKey, nil

	})

	if err == nil {
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			ctx := r.Context()
			ctx = context.WithValue(ctx, UserID, claims["userId"])
			next(rw, r.WithContext(ctx))
		}
	}

	rw.WriteHeader(http.StatusUnauthorized)
}

func initJWTAuthenticationSettings() *JWTAuthenticationSettings {
	if authenticationSettings == nil {
		authenticationSettings = &JWTAuthenticationSettings{
			privateKey: getPrivateKey(),
			PublicKey:  getPublicKey(),
		}
	}

	return authenticationSettings
}

func getPrivateKey() *rsa.PrivateKey {

	path, err := filepath.Abs(privateKey)

	if err != nil {
		panic(err)
	}

	prKey, err := os.Open(path)
	defer prKey.Close()

	if err != nil {
		panic(err)
	}

	fmt.Println(prKey)

	return nil
}

func getPublicKey() *rsa.PublicKey {
	path, err := filepath.Abs(publicKey)

	if err != nil {
		panic(err)
	}

	pubKey, err := os.Open(path)

	defer pubKey.Close()

	if err != nil {
		panic(err)
	}

	fmt.Println(pubKey)

	return nil
}
