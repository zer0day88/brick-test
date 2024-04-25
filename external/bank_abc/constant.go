package bank_abc

type BankABCAPI struct {
	Name    string
	UrlPath string
}

var (
	InquiryAccountAPI BankABCAPI = BankABCAPI{Name: "inquiry account", UrlPath: "/snap/v1.0/account-inquiry"}
	TransferAPI       BankABCAPI = BankABCAPI{Name: "transfer", UrlPath: "/snap/v1.0/transfer-bank"}
)
