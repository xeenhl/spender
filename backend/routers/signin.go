package routers

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"net/http"

	appCtx "github.com/xeenhl/spender/backend/context"
	"github.com/xeenhl/spender/backend/models"
)

func (env *Env) SignIn(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	var creds models.Credentials
	decodeError := json.NewDecoder(r.Body).Decode(&creds)

	if decodeError != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	pass := base64.StdEncoding.EncodeToString([]byte(creds.Password))
	creds.Password = pass



	ctx := r.Context()
	ctx = context.WithValue(ctx, appCtx.Credentils, creds)

	next(rw, r.WithContext(ctx))
}

func sendVerificationEmail(signin models.Credentials) {

}
