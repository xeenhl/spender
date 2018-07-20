package routers

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/xeenhl/spender/backend/authentication"
	appCtx "github.com/xeenhl/spender/backend/context"
	"github.com/xeenhl/spender/backend/models"
)

func (env *Env) Login(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	var userLogin authentication.Login
	decodeError := json.NewDecoder(r.Body).Decode(&userLogin)

	if decodeError != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	pass := base64.StdEncoding.EncodeToString([]byte(userLogin.Password))
	userLogin.Password = pass

	user, verified := allowed(&userLogin, env)
	if !verified {
		rw.WriteHeader(http.StatusUnauthorized)
		rw.Write([]byte("Not registered user"))
		return
	}

	ctx := r.Context()
	ctx = context.WithValue(ctx, appCtx.LoginData, *user)

	next(rw, r.WithContext(ctx))

}

func allowed(login *authentication.Login, env *Env) (*models.User, bool) {

	_, err := env.getLoginData(login.Email)

	if err != nil {
		return nil, false
	}

	expectedP := base64.StdEncoding.EncodeToString([]byte("qwerty"))

	if login.Email == "test@test.com" && login.Password == expectedP {
		return &models.User{ID: 10}, true
	}
	return nil, false
}

func (enc *Env) getLoginData(email string) (*authentication.Login, error) {
	return nil, nil
}
