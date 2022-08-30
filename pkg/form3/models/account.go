package models

type Account struct {
	Data *AccountData `json:"data"`
}

func (a *Account) WithData(data AccountData) *Account {
	a.Data = &data
	return a
}

type AccountData struct {
	Attributes     *AccountAttributes `json:"attributes,omitempty"`
	ID             string             `json:"id,omitempty"`
	OrganisationID string             `json:"organisation_id,omitempty"`
	Type           string             `json:"type,omitempty"`
	Version        *int64             `json:"version,omitempty"`
}

func (ad *AccountData) WithID(id string) *AccountData {
	ad.ID = id
	return ad
}

func (ad *AccountData) WithOrganisationID(organisationID string) *AccountData {
	ad.OrganisationID = organisationID
	return ad
}

func (ad *AccountData) WithType(_type string) *AccountData {
	ad.Type = _type
	return ad
}

func (ad *AccountData) WithVersion(version int64) *AccountData {
	ad.Version = &version
	return ad
}

func (ad *AccountData) WithAttributes(attributes AccountAttributes) *AccountData {
	ad.Attributes = &attributes
	return ad
}
