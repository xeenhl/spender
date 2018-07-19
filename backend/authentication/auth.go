package authentication

import (
	"context"
	"crypto/rsa"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	appCtx "github.com/xeenhl/spender/backend/context"
	"github.com/xeenhl/spender/backend/models"

	"bufio"
	"crypto/x509"
	"encoding/pem"
	"errors"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

type JWTAuthenticationSettings struct {
	privateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

const (
	tokenDuration = 72
	expireOffset  = 3600
	privateKey    = "./authentication/keys/private_key"
	publicKey     = "./authentication/keys/public_key.pub"
)

var authenticationSettings *JWTAuthenticationSettings

func initJWTAuthenticationSettings() *JWTAuthenticationSettings {
	if authenticationSettings == nil {
		authenticationSettings = &JWTAuthenticationSettings{
			privateKey: getPrivateKey(),
			PublicKey:  getPublicKey(),
		}
	}

	return authenticationSettings
}

func AuhtWithToken(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	auth := initJWTAuthenticationSettings()

	token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSAPSS); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return auth.PublicKey, nil

	})

	if err != nil {
		rw.WriteHeader(http.StatusUnauthorized)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		ctx := r.Context()
		ctx = context.WithValue(ctx, appCtx.UserID, claims["userId"])
		next(rw, r.WithContext(ctx))
	}

}

func SignWithNewToken(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	auth := initJWTAuthenticationSettings()
	login, err := getLogin(r.Context().Value(appCtx.LoginData))
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("no login data provided"))
	}

	token := jwt.NewWithClaims(jwt.SigningMethodPS512, jwt.MapClaims{
		"userId": login.ID,
	})

	tokenString, err := token.SignedString(auth.privateKey)

	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("unprocessable login data"))
		return
	}

	http.SetCookie(rw, &http.Cookie{
		Name:  "token",
		Value: tokenString,
	})

	next(rw, r)
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

	keyStat, _ := prKey.Stat()

	var s int64 = keyStat.Size()
	pembytes := make([]byte, s)

	buffer := bufio.NewReader(prKey)
	buffer.Read(pembytes)

	data, _ := pem.Decode([]byte(pembytes))

	pubKeyImported, err := x509.ParsePKCS1PrivateKey(data.Bytes)

	if err != nil {
		panic(err)
	}

	return pubKeyImported
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

	keyStat, _ := pubKey.Stat()

	s := keyStat.Size()
	pembytes := make([]byte, s)

	buffer := bufio.NewReader(pubKey)
	buffer.Read(pembytes)

	data, _ := pem.Decode([]byte(pembytes))

	pubKeyImported, err := x509.ParsePKIXPublicKey(data.Bytes)

	if err != nil {
		panic(err)
	}

	rsaPub, ok := pubKeyImported.(*rsa.PublicKey)

	if !ok {
		panic("cant validate public key")
	}

	return rsaPub
}

func getLogin(l interface{}) (*models.User, error) {
	switch login := l.(type) {
	case models.User:
		login = models.User(login)
		return &login, nil
	default:
		return nil, errors.New("no login data provided")
	}
}
