package repository

import (
	"context"
	"github.com/zer0day88/brick-test/internal/app/domain/entities"
	"gorm.io/gorm"
)

type PgRepository struct {
	db *gorm.DB
}

func NewPgRepository(db *gorm.DB) *PgRepository {
	return &PgRepository{db: db}
}

func (r *PgRepository) TransferInsert(ctx context.Context, transfer *entities.Transfer) error {

	res := r.db.WithContext(ctx).Create(transfer)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (r *PgRepository) UpdateStatus(ctx context.Context, status string, refNo string) error {

	return r.db.WithContext(ctx).
		Model(&entities.Transfer{}).
		Where("ref_no = ?", refNo).
		Update("status", status).Error
}

func (r *PgRepository) FindTransferOneBy(ctx context.Context, criteria map[string]interface{}) (*entities.Transfer, error) {
	var tf entities.Transfer

	res := r.db.WithContext(ctx).
		Model(&entities.Transfer{}).
		Where(criteria).
		First(&tf)
	if res.Error != nil {
		return nil, res.Error
	}

	return &tf, nil
}
