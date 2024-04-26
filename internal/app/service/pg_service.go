package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/shopspring/decimal"
	"github.com/zer0day88/brick-test/external/bank_abc"
	"github.com/zer0day88/brick-test/internal/app/domain/entities"
	"github.com/zer0day88/brick-test/internal/app/domain/repository"
	modelpg "github.com/zer0day88/brick-test/internal/app/model"
	pkgerr "github.com/zer0day88/brick-test/pkg/error"
	"time"
)

type PG struct {
	log        zerolog.Logger
	bankABCSrv *bank_abc.BankAbcSrv
	pgRepo     repository.PgRepository
}

func NewPGService(log zerolog.Logger, srv *bank_abc.BankAbcSrv, pgRepo repository.PgRepository) *PG {
	return &PG{
		log:        log,
		bankABCSrv: srv,
		pgRepo:     pgRepo,
	}
}

func (s *PG) InquiryAccount(ctx context.Context, inquiryAccount modelpg.InquiryAccountRequest) (*modelpg.InquiryAccountResponse, pkgerr.CustomErr) {

	if !inquiryAccount.Validate() {
		return nil, pkgerr.InquiryAccountBadParam
	}

	req := bank_abc.InquiryAccountRequest{
		BeneficiaryAccountNumber: inquiryAccount.BeneficiaryAccountNumber,
		BeneficiaryBankCode:      inquiryAccount.BeneficiaryBankCode,
	}

	if inquiryAccount.AdditionalInfo != nil {
		req.AdditionalInfo = &bank_abc.AdditionalInfo{
			Channel:  inquiryAccount.AdditionalInfo.Channel,
			DeviceID: inquiryAccount.AdditionalInfo.DeviceID,
		}
	}

	result, err := s.bankABCSrv.InquiryAccount(req)
	if err != nil {
		s.log.Err(err).Send()
		return nil, pkgerr.InquiryAccountInternalError
	}

	response := modelpg.InquiryAccountResponse{
		BeneficiaryAccountName:   result.BeneficiaryAccountName,
		BeneficiaryAccountNumber: result.BeneficiaryAccountNumber,
		BeneficiaryBankCode:      result.BeneficiaryBankCode,
		BeneficiaryBankName:      result.BeneficiaryBankName,
		Currency:                 result.Currency,
		ReferenceNo:              result.ReferenceNo,
	}

	if result.AdditionalInfo != nil {
		response.AdditionalInfo = &modelpg.AdditionalInfo{
			Channel:  result.AdditionalInfo.Channel,
			DeviceID: result.AdditionalInfo.DeviceID,
		}
	}

	return &response, pkgerr.InquiryAccountNoError
}

func (s *PG) Transfer(ctx context.Context, transfer modelpg.TransferRequest) (*modelpg.TransferResponse, pkgerr.CustomErr) {

	if !transfer.Validate() {
		return nil, pkgerr.TransferBadParam
	}

	//check idempotency by partnerRefNo
	//if found than return message transaction in progress

	//TODO: check source account and balance

	RefNo := uuid.NewString()
	amountValue, err := decimal.NewFromString(transfer.Amount.Value)
	if err != nil {
		s.log.Err(err).Send()
		return nil, pkgerr.TransferInvalidFieldFormat
	}

	//insert with different goroutine
	go func(transfer modelpg.TransferRequest, refNo string, amount decimal.Decimal) {

		now := time.Now()

		tfRequest := bank_abc.TransferRequest{
			Amount: bank_abc.Amount{
				Currency: transfer.Amount.Currency,
				Value:    transfer.Amount.Value,
			},
			BeneficiaryAccountName:   transfer.BeneficiaryAccountName,
			BeneficiaryAccountNumber: transfer.BeneficiaryAccountNumber,
			BeneficiaryBankCode:      transfer.BeneficiaryBankCode,
			BeneficiaryBankName:      transfer.BeneficiaryBankName,
			PartnerReferenceNo:       transfer.PartnerReferenceNo,
			SourceAccountNumber:      transfer.SourceAccountNumber,
			TransactionDate:          now.Format(time.RFC3339),
		}

		if transfer.AdditionalInfo != nil {
			tfRequest.AdditionalInfo.DeviceID = transfer.AdditionalInfo.DeviceID
			tfRequest.AdditionalInfo.Channel = transfer.AdditionalInfo.Channel
		}

		_, errTransfer := s.bankABCSrv.Transfer(tfRequest)
		if errTransfer != nil {
			s.log.Err(errTransfer).Send()
		}

		transferEntity := entities.Transfer{
			ID:                       uuid.NewString(),
			RefNo:                    RefNo,
			PartnerRefNo:             transfer.PartnerReferenceNo,
			AmountValue:              amount.InexactFloat64(),
			AmountCurrency:           transfer.Amount.Currency,
			BeneficiaryAccountNumber: transfer.BeneficiaryAccountNumber,
			BeneficiaryAccountName:   transfer.BeneficiaryAccountName,
			BeneficiaryBankCode:      transfer.BeneficiaryBankCode,

			SourceAccountNumber: transfer.SourceAccountNumber,
			TxTime:              now,
			CreatedAt:           now,
		}

		if transfer.BeneficiaryBankName != nil {
			transferEntity.BeneficiaryBankName = *transfer.BeneficiaryBankName
		}

		if transfer.AdditionalInfo != nil {
			transferEntity.DeviceID = transfer.AdditionalInfo.DeviceID
			transferEntity.Channel = transfer.AdditionalInfo.Channel
		}

		errInsert := s.pgRepo.TransferInsert(context.Background(), &transferEntity)
		if errInsert != nil {
			//this error should be retryable
			s.log.Err(errInsert).Send()
		}
	}(transfer, RefNo, amountValue)

	result := modelpg.TransferResponse{
		PartnerReferenceNo: transfer.PartnerReferenceNo,
		ReferenceNo:        RefNo,
		Amount: modelpg.Amount{
			Currency: transfer.Amount.Currency,
			Value:    transfer.Amount.Value,
		},
		BeneficiaryAccountNumber: transfer.BeneficiaryAccountNumber,
		BeneficiaryBankCode:      transfer.BeneficiaryBankCode,
		SourceAccountNo:          transfer.SourceAccountNumber,
	}

	if transfer.AdditionalInfo != nil {
		result.AdditionalInfo.DeviceID = transfer.AdditionalInfo.DeviceID
		result.AdditionalInfo.Channel = transfer.AdditionalInfo.Channel
	}

	return &result, pkgerr.TransferNoError
}

func (s *PG) NotifyTransfer(ctx context.Context, notifyTransfer modelpg.NotifyTransferRequest) pkgerr.CustomErr {

	if !notifyTransfer.Validate() {
		return pkgerr.NotifyTransferBadParam
	}

	_, err := s.pgRepo.FindTransferOneBy(ctx, map[string]interface{}{
		"ref_no": notifyTransfer.ReferenceNo,
	})
	if err != nil {
		return pkgerr.NotifyTransferRefNotFound
	}

	err = s.pgRepo.UpdateStatus(ctx, notifyTransfer.TransactionStatus, notifyTransfer.ReferenceNo)
	if err != nil {
		s.log.Err(err).Send()
	}

	return pkgerr.NotifyTransferNoError
}
