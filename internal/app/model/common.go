package model_pg

type Amount struct {
	Currency string `json:"currency"`
	Value    string `json:"value"`
}

type AdditionalInfo struct {
	Channel  *string `json:"channel,omitempty"`
	DeviceID *string `json:"deviceId,omitempty"`
}
