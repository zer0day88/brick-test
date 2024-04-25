package bank_abc

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

const baseURL = "https://mock.apidog.com/m1/520155-0-default"

func TestInquiryAccountAPI(t *testing.T) {
	srv := New(baseURL)
	result, err := srv.InquiryAccount(InquiryAccountRequest{
		BeneficiaryAccountNumber: "1111",
		BeneficiaryBankCode:      "111",
	})
	assert.Nil(t, err)
	assert.Equal(t, "200500", result.ResponseCode)
}

func TestTransferAPI(t *testing.T) {
	srv := New(baseURL)
	result, err := srv.Transfer(TransferRequest{
		Amount: Amount{
			Currency: "IDR",
			Value:    "200000.00",
		},
		BeneficiaryAccountName:   "111",
		BeneficiaryAccountNumber: "1111",
		BeneficiaryBankCode:      "111",
		PartnerReferenceNo:       "111",
		SourceAccountNumber:      "1111",
		TransactionDate:          time.Now().Format(time.RFC3339),
	})
	assert.Nil(t, err)
	assert.Equal(t, "200600", result.ResponseCode)
}
