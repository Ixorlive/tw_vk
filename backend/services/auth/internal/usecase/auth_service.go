package usecase

import (
	"github.com/Ixorlive/tw_vk/backend/services/auth/internal/entity"
	"github.com/Ixorlive/tw_vk/backend/services/auth/internal/usecase/jwt"
	jwtv4 "github.com/golang-jwt/jwt/v4"
	"golang.org/x/net/context"
)

type BasicAuthService struct {
	jwtGenerator jwt.JWTGenerator
	jwtValidator jwt.JWTValidator
	repo         UserRepo
}

func NewAuthService(repo UserRepo) AuthService {
	var aservice BasicAuthService
	aservice.jwtGenerator = jwt.NewJWTGenerator(jwtv4.SigningMethodHS256)
	aservice.jwtValidator = jwt.NewJWTValidator()
	aservice.repo = repo
	return &aservice
}

func (a *BasicAuthService) Auth(ctx context.Context, user entity.User) (*jwt.AccessToken, error) {
	userExists, err := a.repo.UserExists(ctx, user, false)
	if err != nil {
		return nil, err
	}
	if userExists {
		token, err := a.jwtGenerator.Generate(user.Login)
		if err != nil {
			return nil, err
		}
		return token, nil
	}
	return nil, err
}

func (a *BasicAuthService) AuthByToken(ctx context.Context, jwtToken string) (entity.User, error) {
	login, err := a.jwtValidator.ValidateToken(jwtToken)
	if err != nil {
		return entity.User{}, err
	}
	return entity.User{Login: login}, nil
}

func (a *BasicAuthService) Register(ctx context.Context, newUser entity.User) (RegistrationStatus, error) {
	userExists, err := a.repo.UserExists(ctx, newUser, true)
	if err != nil {
		return Error, err
	}
	if userExists {
		return UserExists, nil
	}
	inserted, err := a.repo.AddUser(ctx, newUser)
	if err != nil {
		return Error, err
	}
	if inserted {
		return Success, nil
	}
	return IncorrectLoginOrPassword, nil
}
