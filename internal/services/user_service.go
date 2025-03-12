package services

import (
	"fmt"
	"itv_movies_web_server/internal/models"
	"itv_movies_web_server/internal/repositories"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	GetUsers() ([]models.User, error)
	RegisterUser(user *models.User) error
	Login(email string, password string) (string, error)
}

type userService struct {
	userRepo repositories.UserRepository
}

var jwtSecret = []byte(os.Getenv("JWT_SECRET_KEY"))

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{userRepo}
}

func (s *userService) GetUsers() ([]models.User, error) {
	return s.userRepo.FindAll()
}

func (s *userService) RegisterUser(user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.ID = uuid.New()
	user.Password = string(hashedPassword)
	return s.userRepo.Register(user)
}
func (s *userService) Login(email string, password string) (string, error) {
	var user *models.User
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return "", fmt.Errorf("user not found")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", fmt.Errorf("incorrect password")
	}

	token, err := s.generateJWT(user.ID, user.Email, string(user.Role))
	if err != nil {
		return "", err
	}

	return token, nil
}
func (s userService) generateJWT(userID uuid.UUID, email string, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID.String(),
		"email":   email,
		"role":    role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", fmt.Errorf("could not sign the token: %v", err)
	}

	return signedToken, nil
}
