package resty

import (
	"github.com/go-resty/resty/v2"
	"time"
)

func GetHttpClient(baseURL string, timeOut time.Duration) *resty.Client {
	goRestyClient := resty.New()
	goRestyClient.SetDebug(false)
	goRestyClient.SetTimeout(timeOut)
	goRestyClient.SetBaseURL(baseURL)

	return goRestyClient
}
