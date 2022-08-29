package accounts

import (
	"encoding/json"
	"github.com/francorosatti/form3-api-client/pkg/form3/models"
)

func serialiseAccountData(model models.AccountData) ([]byte, error) {
	return json.Marshal(model)
}

func parseAccountData(bytes []byte) (models.AccountData, error) {
	var model models.AccountData
	err := json.Unmarshal(bytes, &model)
	return model, err
}
