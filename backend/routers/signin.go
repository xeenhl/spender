package routers

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/xeenhl/spender/backend/authentication"
	appCtx "github.com/xeenhl/spender/backend/context"
)

func (env *Env) SignIn(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	var userSignIn authentication.Signin
	decodeError := json.NewDecoder(r.Body).Decode(&userSignIn)

	if decodeError != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	pass := base64.StdEncoding.EncodeToString([]byte(userSignIn.Password))
	userSignIn.Password = pass

	ctx := r.Context()
	ctx = context.WithValue(ctx, appCtx.SigninData, *userSignIn)

	next(rw, r.WithContext(ctx))
}
