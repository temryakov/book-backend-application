package usecase

import (
	"context"
	"time"

	"github.com/temryakov/go-backend-book-app/user-service/domain"
)

type signupUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewSignupUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.SignupUsecase {
	return &signupUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (su *signupUsecase) Create(c context.Context, user *domain.User) error {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	if user, _ := su.userRepository.FetchByEmail(ctx, user.Email); user != nil {
		return domain.ErrUserAlreadyExists
	}
	return su.userRepository.Create(ctx, user)
}
