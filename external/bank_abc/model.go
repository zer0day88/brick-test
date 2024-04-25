package bank_abc

type InquiryAccountRequest struct {
	BeneficiaryAccountNumber string          `json:"beneficiaryAccountNumber"`
	BeneficiaryBankCode      string          `json:"beneficiaryBankCode"`
	AdditionalInfo           *AdditionalInfo `json:"additionalInfo,omitempty"`
}

type InquiryAccountResponse struct {
	BeneficiaryAccountName   string          `json:"beneficiaryAccountName"`
	BeneficiaryAccountNumber string          `json:"beneficiaryAccountNumber"`
	BeneficiaryBankCode      string          `json:"beneficiaryBankCode"`
	BeneficiaryBankName      string          `json:"beneficiaryBankName"`
	Currency                 string          `json:"currency"`
	ReferenceNo              string          `json:"referenceNo"`
	ResponseCode             string          `json:"responseCode"`
	ResponseMessage          string          `json:"responseMessage"`
	AdditionalInfo           *AdditionalInfo `json:"additionalInfo,omitempty"`
}

type TransferRequest struct {
	Amount                   Amount          `json:"amount"`
	BeneficiaryAccountName   string          `json:"beneficiaryAccountName"`
	BeneficiaryAccountNumber string          `json:"beneficiaryAccountNumber"`
	BeneficiaryBankCode      string          `json:"beneficiaryBankCode"`
	BeneficiaryBankName      *string         `json:"beneficiaryBankName,omitempty"`
	PartnerReferenceNo       string          `json:"partnerReferenceNo"`
	SourceAccountNumber      string          `json:"sourceAccountNumber"`
	TransactionDate          string          `json:"transactionDate"`
	AdditionalInfo           *AdditionalInfo `json:"additionalInfo,omitempty"`
}

type TransferResponse struct {
	PartnerReferenceNo       string          `json:"partnerReferenceNo"`
	ReferenceNo              string          `json:"referenceNo"`
	Amount                   Amount          `json:"amount"`
	BeneficiaryAccountNumber string          `json:"beneficiaryAccountNumber"`
	BeneficiaryBankCode      string          `json:"beneficiaryBankCode"`
	ResponseCode             string          `json:"responseCode"`
	ResponseMessage          string          `json:"responseMessage"`
	SourceAccountNo          string          `json:"sourceAccountNo"`
	AdditionalInfo           *AdditionalInfo `json:"additionalInfo,omitempty"`
}

type Amount struct {
	Currency string `json:"currency"`
	Value    string `json:"value"`
}

type AdditionalInfo struct {
	Channel  *string `json:"channel,omitempty"`
	DeviceID *string `json:"deviceId,omitempty"`
}
