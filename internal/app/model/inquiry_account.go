package model_pg

type InquiryAccountRequest struct {
	BeneficiaryAccountNumber string          `json:"beneficiaryAccountNumber"`
	BeneficiaryBankCode      string          `json:"beneficiaryBankCode"`
	AdditionalInfo           *AdditionalInfo `json:"additionalInfo,omitempty"`
}

func (v InquiryAccountRequest) Validate() bool {
	if v.BeneficiaryBankCode == "" || v.BeneficiaryAccountNumber == "" {
		return false
	}

	return true
}

type InquiryAccountResponse struct {
	BeneficiaryAccountName   string          `json:"beneficiaryAccountName"`
	BeneficiaryAccountNumber string          `json:"beneficiaryAccountNumber"`
	BeneficiaryBankCode      string          `json:"beneficiaryBankCode"`
	BeneficiaryBankName      string          `json:"beneficiaryBankName"`
	Currency                 string          `json:"currency"`
	ReferenceNo              string          `json:"referenceNo"`
	AdditionalInfo           *AdditionalInfo `json:"additionalInfo,omitempty"`
}
