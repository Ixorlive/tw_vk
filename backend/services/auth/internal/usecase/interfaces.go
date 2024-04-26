package usecase

import (
	"context"

	"github.com/Ixorlive/tw_vk/backend/services/auth/internal/entity"
	"github.com/Ixorlive/tw_vk/backend/services/auth/internal/usecase/jwt"
)

type RegistrationStatus int

const (
	Success RegistrationStatus = iota
	UserExists
	IncorrectLoginOrPassword
	Error
)

type (
	AuthService interface {
		Auth(context.Context, entity.User) (*jwt.AccessToken, error)
		AuthByToken(context.Context, string) (entity.User, error)
		Register(context.Context, entity.User) (RegistrationStatus, error)
	}

	UserRepo interface {
		UserExists(context.Context, entity.User /* check only login */, bool) (bool, error)
		AddUser(context.Context, entity.User) (bool, error)
	}
)
