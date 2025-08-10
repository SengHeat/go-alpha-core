package service

import (
	"alpha-core/internal/config"
	"alpha-core/internal/model"
	"alpha-core/internal/repository"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepository *repository.UserRepository
	configure      *config.Config
}

func NewAuthService(userRepository *repository.UserRepository, configure *config.Config) *AuthService {
	return &AuthService{
		userRepository: userRepository,
		configure:      configure,
	}
}

func (service *AuthService) Register(email, password, name string) (*model.User, error) {
	existing, _ := service.userRepository.FindByEmail(email)

	if existing != nil {
		return nil, errors.New("email already registered")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	user := &model.User{
		Email:    email,
		Password: string(hashedPassword),
		Name:     name,
	}

	if err := service.userRepository.Create(user); err != nil {
		return nil, err
	}

	return user, nil

}

func (service *AuthService) Login(email, password string) (*model.User, error) {
	existing, err := service.userRepository.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	if existing == nil {
		return nil, errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(existing.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid password")
	}

	return existing, nil
}

func (service *AuthService) Authenticate(email, password string) (*model.User, error) {
	user, err := service.userRepository.FindByEmail(email)

	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}

func (service *AuthService) GenerateJWT(user *model.User) (string, error) {
	roles := []string{}
	permissions := []string{}

	for _, role := range user.Roles {
		roles = append(roles, role.Name)
		for _, permission := range role.Permissions {
			permissions = append(permissions, permission.Name)
		}
	}

	cliams := jwt.MapClaims{
		"email":       user.Email,
		"name":        user.Name,
		"roles":       roles,
		"permissions": permissions,
		"exp":         time.Now().Add(time.Hour * 72).Unix(),
		"iat":         time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, cliams)
	return token.SignedString([]byte(service.configure.JWTSecret))
}
