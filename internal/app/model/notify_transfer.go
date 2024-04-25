package model_pg

type NotifyTransferRequest struct {
	ReferenceNo       string `json:"referenceNo"`
	TransactionStatus string `json:"transactionStatus"`
}

func (v NotifyTransferRequest) Validate() bool {
	if v.ReferenceNo == "" || v.TransactionStatus == "" {

		return false
	}

	return true
}
