package routers

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"net/http"

	appCtx "github.com/xeenhl/spender/backend/context"
	"github.com/xeenhl/spender/backend/models"
)

func (env *Env) Login(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	var creds models.Credentials
	decodeError := json.NewDecoder(r.Body).Decode(&creds)

	if decodeError != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	pass := base64.StdEncoding.EncodeToString([]byte(creds.Password))
	creds.Password = pass

	user, verified := allowed(&creds, env)
	if !verified {
		rw.WriteHeader(http.StatusUnauthorized)
		rw.Write([]byte("Not registered user"))
		return
	}

	ctx := r.Context()
	ctx = context.WithValue(ctx, appCtx.Credentils, *user)

	next(rw, r.WithContext(ctx))

}

func allowed(login *models.Credentials, env *Env) (*models.User, bool) {

	userData, err := env.getLoginData(login.Email)

	if err != nil {
		return nil, false
	}



	if login.Email == userData.Email && login.Password == userData.Password {
		return userData, true
	}
	return nil, false
}

func (env *Env) getLoginData(email string) (*models.User, error) {
	usr, err := env.DB.GetUserByEmail(email)
	if err != nil {
		return  nil, err
	}
	return usr, nil
}
