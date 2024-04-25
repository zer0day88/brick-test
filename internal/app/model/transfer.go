package model_pg

type TransferRequest struct {
	Amount                   Amount          `json:"amount"`
	BeneficiaryAccountName   string          `json:"beneficiaryAccountName"`
	BeneficiaryAccountNumber string          `json:"beneficiaryAccountNumber"`
	BeneficiaryBankCode      string          `json:"beneficiaryBankCode"`
	BeneficiaryBankName      *string         `json:"beneficiaryBankName,omitempty"`
	PartnerReferenceNo       string          `json:"partnerReferenceNo"`
	SourceAccountNumber      string          `json:"sourceAccountNumber"`
	AdditionalInfo           *AdditionalInfo `json:"additionalInfo,omitempty"`
}

func (v TransferRequest) Validate() bool {
	if v.Amount.Value == "" || v.Amount.Currency == "" || v.BeneficiaryAccountName == "" ||
		v.BeneficiaryAccountNumber == "" || v.BeneficiaryBankCode == "" ||
		v.PartnerReferenceNo == "" || v.SourceAccountNumber == "" {

		return false
	}

	return true
}

type TransferResponse struct {
	PartnerReferenceNo       string          `json:"partnerReferenceNo"`
	ReferenceNo              string          `json:"referenceNo"`
	Amount                   Amount          `json:"amount"`
	BeneficiaryAccountNumber string          `json:"beneficiaryAccountNumber"`
	BeneficiaryBankCode      string          `json:"beneficiaryBankCode"`
	SourceAccountNo          string          `json:"sourceAccountNo"`
	AdditionalInfo           *AdditionalInfo `json:"additionalInfo,omitempty"`
}
