package usecase

import (
	"context"
	"time"

	"github.com/temryakov/go-backend-book-app/user-service/domain"
)

type userUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewUserUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.ProfileUsecase {
	return &userUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (uu *userUsecase) FetchByID(c context.Context, id uint) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	user, err := uu.userRepository.FetchByID(ctx, id)

	if user == nil {
		return nil, err
	}
	return user, nil
}
