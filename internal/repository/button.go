package repository

import (
	"app/internal/domain"
	"context"
)

func (repo Repository) SetButton(ctx context.Context, b domain.Button) error {
	res := repo.db.Save(&b)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (repo Repository) GetButton(ctx context.Context, b *domain.Button) error {
	res := repo.db.Model(b).Take(b)
	if res.Error != nil {
		return res.Error
	}

	return nil
}
