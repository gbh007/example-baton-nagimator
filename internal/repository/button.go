package repository

import (
	"app/internal/domain"
	"context"
)

func (repo Repository) ButtonCreate(ctx context.Context, b domain.ButtonPressed) error {
	res := repo.db.Create(&b)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (repo Repository) ButtonUpdate(ctx context.Context, b domain.ButtonPressed) error {
	res := repo.db.Save(&b)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (repo Repository) GetButton(ctx context.Context, b domain.ButtonPressed) (domain.ButtonPressed, error) {
	res := repo.db.Model(&b).Where(&b).First(&b)
	if res.Error != nil {
		return domain.ButtonPressed{}, res.Error
	}

	return b, nil
}

// func (repo Repository) GetButton(ctx context.Context, b domain.ButtonPressed) (domain.ButtonPressed, error) {
// 	var b2 domain.ButtonPressed
// 	res := repo.db.Model(&domain.ButtonPressed{}).Where(&domain.ButtonPressed{
// 		Year:  b.Year,
// 		Month: b.Month,
// 		Day:   b.Day,
// 	}).First(&b2)
// 	if res.Error != nil {
// 		return domain.ButtonPressed{}, res.Error
// 	}

// 	return b2, nil
// }
