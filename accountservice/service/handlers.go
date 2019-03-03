/**
 * Create by Le Trong on 17/Feb/2019
 */

package service

import (
	"encoding/json"
	"net/http"

	"github.com/letrong/masonjar-service/accountservice/model"
)

// CreateAccountHandler creta a person
func CreateAccountHandler(w http.ResponseWriter, r *http.Request) {
	var account model.Account
	_ = json.NewDecoder(r.Body).Decode(&account)
	// dao.InsertOneValue(account)
	json.NewEncoder(w).Encode(account)
}
