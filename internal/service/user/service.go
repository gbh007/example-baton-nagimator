package user

import (
	"app/internal/domain"
	"app/internal/repository"
	"context"
	"crypto/md5"
	"encoding/hex"
	"time"
)

type Service struct {
	repo *repository.Repository
}

func New(repo *repository.Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreateUser(ctx context.Context) (domain.User, error) {
	token := md5.Sum([]byte(time.Now().String()))

	u := domain.User{
		Name:  "Test", // FIXME: добавить генерацию
		Token: hex.EncodeToString(token[:]),
	}

	err := s.repo.SetUser(ctx, u)
	if err != nil {
		return domain.User{}, err
	}

	return u, nil
}

func (s *Service) GetUser(ctx context.Context, token string) (domain.User, error) {
	u := domain.User{
		Token: token,
	}

	err := s.repo.GetUserByToken(ctx, &u)
	if err != nil {
		return domain.User{}, err
	}

	return u, nil
}
