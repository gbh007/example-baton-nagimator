package button

import (
	"app/internal/domain"
	"app/internal/repository"
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
)

type Service struct {
	repo *repository.Repository
}

func New(repo *repository.Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) PressButton(ctx context.Context, user domain.User) (domain.Button, error) {
	y, m, d := time.Now().Date()

	b := domain.Button{
		UserID: user.ID,
		Year:   y,
		Month:  int(m),
		Day:    d,
	}

	err := s.repo.GetButton(ctx, &b)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return domain.Button{}, err
	}

	b.Count++

	err = s.repo.SetButton(ctx, b)
	if err != nil {
		return domain.Button{}, err
	}

	return b, nil
}
