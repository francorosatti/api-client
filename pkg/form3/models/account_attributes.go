package models

type AccountAttributes struct {
	AccountClassification   *string  `json:"account_classification,omitempty"`
	AccountMatchingOptOut   *bool    `json:"account_matching_opt_out,omitempty"`
	AccountNumber           string   `json:"account_number,omitempty"`
	AlternativeNames        []string `json:"alternative_names,omitempty"`
	BankID                  string   `json:"bank_id,omitempty"`
	BankIDCode              string   `json:"bank_id_code,omitempty"`
	BaseCurrency            string   `json:"base_currency,omitempty"`
	Bic                     string   `json:"bic,omitempty"`
	Country                 *string  `json:"country,omitempty"`
	Iban                    string   `json:"iban,omitempty"`
	JointAccount            *bool    `json:"joint_account,omitempty"`
	Name                    []string `json:"name,omitempty"`
	SecondaryIdentification string   `json:"secondary_identification,omitempty"`
	Status                  *string  `json:"status,omitempty"`
	Switched                *bool    `json:"switched,omitempty"`
}

func (aa *AccountAttributes) WithAccountClassification(accountClassification string) *AccountAttributes {
	aa.AccountClassification = &accountClassification
	return aa
}

func (aa *AccountAttributes) WithAccountMatchingOptOut(accountMatchingOptOut bool) *AccountAttributes {
	aa.AccountMatchingOptOut = &accountMatchingOptOut
	return aa
}

func (aa *AccountAttributes) WithAccountNumber(accountNumber string) *AccountAttributes {
	aa.AccountNumber = accountNumber
	return aa
}

func (aa *AccountAttributes) WithAlternativeNames(alternativeNames []string) *AccountAttributes {
	aa.AlternativeNames = append([]string{}, alternativeNames...)
	return aa
}

func (aa *AccountAttributes) WithBankID(bankID string) *AccountAttributes {
	aa.BankID = bankID
	return aa
}

func (aa *AccountAttributes) WithBankIDCode(bankIDCode string) *AccountAttributes {
	aa.BankIDCode = bankIDCode
	return aa
}

func (aa *AccountAttributes) WithBaseCurrency(baseCurrency string) *AccountAttributes {
	aa.BaseCurrency = baseCurrency
	return aa
}

func (aa *AccountAttributes) WithBic(bic string) *AccountAttributes {
	aa.Bic = bic
	return aa
}

func (aa *AccountAttributes) WithCountry(country string) *AccountAttributes {
	aa.Country = &country
	return aa
}

func (aa *AccountAttributes) WithIban(iban string) *AccountAttributes {
	aa.Iban = iban
	return aa
}

func (aa *AccountAttributes) WithJointAccount(jointAccount bool) *AccountAttributes {
	aa.JointAccount = &jointAccount
	return aa
}

func (aa *AccountAttributes) WithName(name []string) *AccountAttributes {
	aa.Name = append([]string{}, name...)
	return aa
}

func (aa *AccountAttributes) WithSecondaryIdentification(secondaryIdentification string) *AccountAttributes {
	aa.SecondaryIdentification = secondaryIdentification
	return aa
}

func (aa *AccountAttributes) WithStatus(status string) *AccountAttributes {
	aa.Status = &status
	return aa
}

func (aa *AccountAttributes) WithSwitched(switched bool) *AccountAttributes {
	aa.Switched = &switched
	return aa
}
