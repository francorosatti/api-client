package accounts

import (
	"encoding/json"
	"fmt"
	"github.com/francorosatti/form3-api-client/pkg/form3/models"
)

func accountDataToJson(model models.Account) ([]byte, error) {
	return json.Marshal(model)
}

func jsonToAccountData(bytes []byte) (models.Account, error) {
	var model models.Account
	err := json.Unmarshal(bytes, &model)
	if err != nil {
		return model, fmt.Errorf("%w: %s", errResponseUnmarshal, err)
	}
	return model, err
}
