package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/xeenhl/rest-api-demo/backend/models"
)

var spends = []models.Spend{
	models.Spend{ID: 10, User: models.User{ID: 101}, Amount: &models.Amount{Amount: 13.11, Currency: "GBP"}},
	models.Spend{ID: 10, User: models.User{ID: 101}, Amount: &models.Amount{Amount: 13.11, Currency: "GBP"}},
	models.Spend{ID: 10, User: models.User{ID: 101}, Amount: &models.Amount{Amount: 13.11, Currency: "GBP"}},
	models.Spend{ID: 10, User: models.User{ID: 101}, Amount: &models.Amount{Amount: 13.11, Currency: "GBP"}},
}

func GetSpendsHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	json.NewEncoder(rw).Encode(spends)
}
