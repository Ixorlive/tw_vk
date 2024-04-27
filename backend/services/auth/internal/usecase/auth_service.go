package usecase

import (
	"fmt"

	"github.com/Ixorlive/tw_vk/backend/services/auth/internal/entity"
	"github.com/Ixorlive/tw_vk/backend/services/auth/internal/usecase/crypto"
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
	foundUser, err := a.repo.FindByLogin(ctx, user.Login)
	if err != nil {
		return nil, err
	}
	if foundUser.Login == "" || crypto.CheckPasswordHash(user.Password, foundUser.Password) {
		return nil, fmt.Errorf("Incorrect Login or Password")
	}
	jwtToken, err := a.jwtGenerator.Generate(user)
	return jwtToken, nil
}

func (a *BasicAuthService) AuthByToken(ctx context.Context, jwtToken string) (entity.User, error) {
	user, err := a.jwtValidator.ValidateToken(jwtToken)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (a *BasicAuthService) Register(ctx context.Context, newUser entity.User) (RegistrationStatus, error) {
	foundUser, err := a.repo.FindByLogin(ctx, newUser.Login)
	if err != nil {
		return Error, err
	}
	if foundUser.Login != "" {
		return UserExists, nil
	}
	hashPwd, err := crypto.HashPassword(newUser.Password)
	if err != nil {
		return 0, fmt.Errorf("error generate hash for password %s: %w", newUser.Password, err)
	}
	inserted, err := a.repo.AddUser(ctx, entity.User{Login: newUser.Login, Password: hashPwd})
	if err != nil {
		return Error, err
	}
	if inserted {
		return Success, nil
	}
	return IncorrectLoginOrPassword, nil
}
