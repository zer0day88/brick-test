package bank_abc

import (
	"github.com/rs/zerolog"
	"github.com/zer0day88/brick-test/pkg/resty"
	"time"
)

type BankAbcSrv struct {
	BaseURL string
	log     zerolog.Logger
}

func New(baseURL string) *BankAbcSrv {
	return &BankAbcSrv{BaseURL: baseURL}
}

func (s *BankAbcSrv) InquiryAccount(inquiryAccount InquiryAccountRequest) (*InquiryAccountResponse, error) {
	httpClient := resty.GetHttpClient(s.BaseURL, 10*time.Second)

	resultBody := new(InquiryAccountResponse)

	_, err := httpClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(&inquiryAccount).
		SetResult(resultBody).
		Post(InquiryAccountAPI.UrlPath)

	if err != nil {
		return nil, err
	}

	return resultBody, nil
}

func (s *BankAbcSrv) Transfer(transfer TransferRequest) (*TransferResponse, error) {
	httpClient := resty.GetHttpClient(s.BaseURL, 10*time.Second)

	resultBody := new(TransferResponse)

	_, err := httpClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(&transfer).
		SetResult(resultBody).
		Post(TransferAPI.UrlPath)

	if err != nil {
		return nil, err
	}

	return resultBody, nil
}
