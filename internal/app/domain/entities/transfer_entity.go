package entities

import "time"

type Transfer struct {
	ID                       string     `gorm:"column:id;type:uuid;primaryKey" json:"id"`
	RefNo                    string     `gorm:"column:ref_no;type:character varying;primaryKey" json:"ref_no"`
	PartnerRefNo             string     `gorm:"column:partner_ref_no;type:character varying;not null" json:"partner_ref_no"`
	AmountValue              float64    `gorm:"column:amount_value;type:numeric;not null" json:"amount_value"`
	AmountCurrency           string     `gorm:"column:amount_currency;type:character varying;not null" json:"amount_currency"`
	BeneficiaryAccountNumber string     `gorm:"column:beneficiary_account_number;type:character varying;not null" json:"beneficiary_account_number"`
	BeneficiaryAccountName   string     `gorm:"column:beneficiary_account_name;type:character varying;not null" json:"beneficiary_account_name"`
	BeneficiaryBankCode      string     `gorm:"column:beneficiary_bank_code;type:character varying;not null" json:"beneficiary_bank_code"`
	BeneficiaryBankName      string     `gorm:"column:beneficiary_bank_name;type:character varying;not null" json:"beneficiary_bank_name"`
	SourceAccountNumber      string     `gorm:"column:source_account_number;type:character varying;not null" json:"source_account_number"`
	TxTime                   time.Time  `gorm:"column:tx_time;type:date;not null" json:"tx_time"`
	Status                   string     `gorm:"column:status;type:character varying;not null;default:PENDING" json:"status"`
	DeviceID                 *string    `gorm:"column:device_id;type:character varying" json:"device_id"`
	Channel                  *string    `gorm:"column:channel;type:character varying" json:"channel"`
	CreatedAt                time.Time  `gorm:"column:created_at;type:timestamp with time zone;not null" json:"created_at"`
	UpdatedAt                *time.Time `gorm:"column:updated_at;type:timestamp with time zone" json:"updated_at"`
}
