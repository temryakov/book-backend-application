package usecase

import (
	"context"
	"time"

	"github.com/temryakov/go-backend-book-app/user-service/domain"
)

type profileUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewProfileUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.ProfileUsecase {
	return &profileUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (pu *profileUsecase) FetchByID(c context.Context, id uint) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()
	if user, err := pu.userRepository.FetchByID(ctx, id); user == nil {
		return nil, err
	}
	return pu.userRepository.FetchByID(ctx, id)
}
