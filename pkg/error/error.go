package pkgerr

import "net/http"

type CustomErr struct {
	Msg      string
	Code     string
	HttpCode int
}

var (
	InquiryAccountNoError       = CustomErr{Code: "200001", Msg: "Successful", HttpCode: http.StatusOK}
	InquiryAccountInternalError = CustomErr{Code: "500001", Msg: "Internal Server Error", HttpCode: http.StatusInternalServerError}
	InquiryAccountBadParam      = CustomErr{Code: "400001", Msg: "Bad Parameter", HttpCode: http.StatusBadRequest}

	TransferNoError            = CustomErr{Code: "200101", Msg: "Successful", HttpCode: http.StatusOK}
	TransferBadParam           = CustomErr{Code: "400101", Msg: "Bad Parameter", HttpCode: http.StatusBadRequest}
	TransferInvalidFieldFormat = CustomErr{Code: "400102", Msg: "Invalid Field Format", HttpCode: http.StatusBadRequest}

	NotifyTransferNoError            = CustomErr{Code: "200201", Msg: "Successful", HttpCode: http.StatusOK}
	NotifyTransferBadParam           = CustomErr{Code: "400201", Msg: "Bad Parameter", HttpCode: http.StatusBadRequest}
	NotifyTransferInvalidFieldFormat = CustomErr{Code: "400202", Msg: "Invalid Field Format", HttpCode: http.StatusBadRequest}
	NotifyTransferRefNotFound        = CustomErr{Code: "404203", Msg: "Reference No. not found", HttpCode: http.StatusNotFound}
)
