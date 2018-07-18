package routers

import (
	"net/http"
	"github.com/xeenhl/spender/backend/models"
	"encoding/json"
	"context"
	appCtx 	"github.com/xeenhl/spender/backend/context"
)

func Login(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	var userLogin models.User
	decodeError := json.NewDecoder(r.Body).Decode(&userLogin)

	if decodeError != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	ctx = context.WithValue(ctx, appCtx.LoginData, userLogin)

	next(rw, r.WithContext(ctx))


}
