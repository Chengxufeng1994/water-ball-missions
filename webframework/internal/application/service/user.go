package service

import (
	"fmt"
	"regexp"
	"strings"
	"sync"

	"github.com/Chengxufeng1994/water-ball-missions/webframework/internal/application/dto"
	"github.com/google/uuid"
)

const (
	minEmailLength    = 4
	maxEmailLength    = 32
	minNameLength     = 5
	maxNameLength     = 32
	minPasswordLength = 5
	maxPasswordLength = 32
)

var (
	emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
)

type (
	UserApplicationService interface {
		Register(email string, username string, password string) (dto.User, error)
		Login(email string, password string) (dto.User, dto.Token, error)
		Rename(id string, name string) (string, error)
		List(keyword string) ([]dto.User, error)
		ValidateToken(token string) (string, error)
	}

	userApplicationService struct {
		mu     sync.Mutex
		users  map[string]dto.User
		tokens map[string]string // token -> userID
	}
)

var _ UserApplicationService = (*userApplicationService)(nil)

func NewUserApplicationService() *userApplicationService {
	fmt.Println("userApplicationService Created")
	return &userApplicationService{
		users:  make(map[string]dto.User),
		tokens: make(map[string]string),
	}
}

func (svc *userApplicationService) Register(email string, username string, password string) (dto.User, error) {
	svc.mu.Lock()
	defer svc.mu.Unlock()

	if !emailRegex.MatchString(email) || len(email) < minEmailLength || len(email) > maxEmailLength ||
		len(username) < minNameLength || len(username) > maxNameLength ||
		len(password) < minPasswordLength || len(password) > maxPasswordLength {
		return dto.User{}, fmt.Errorf("Registration's format incorrect.")
	}

	if _, exists := svc.users[email]; exists {
		return dto.User{}, fmt.Errorf("Duplicate email")
	}

	newUser := dto.User{
		ID:       uuid.New().String(),
		Email:    email,
		Username: username,
		Password: password,
	}
	svc.users[email] = newUser
	return newUser, nil
}

func (svc *userApplicationService) Login(email string, password string) (dto.User, dto.Token, error) {
	svc.mu.Lock()
	defer svc.mu.Unlock()

	if !emailRegex.MatchString(email) || len(email) < minEmailLength || len(email) > maxEmailLength ||
		len(password) < minPasswordLength || len(password) > maxPasswordLength {
		return dto.User{}, dto.Token{}, fmt.Errorf("Login's format incorrect.")
	}

	user, exists := svc.users[email]
	if !exists || user.Password != password {
		return dto.User{}, dto.Token{}, fmt.Errorf("Credentials Invalid")
	}

	token := dto.Token{Token: uuid.New().String()}
	svc.tokens[token.Token] = user.ID // Store token and associated user ID

	return user, token, nil
}

func (svc *userApplicationService) ValidateToken(token string) (string, error) {
	svc.mu.Lock()
	defer svc.mu.Unlock()

	userID, exists := svc.tokens[token]
	if !exists {
		return "", fmt.Errorf("Can't authenticate who you are.")
	}
	return userID, nil
}

func (svc *userApplicationService) Rename(id string, name string) (string, error) {
	svc.mu.Lock()
	defer svc.mu.Unlock()

	if len(name) < minNameLength || len(name) > maxNameLength {
		return "", fmt.Errorf("Name's format invalid.")
	}

	for email, user := range svc.users {
		if user.ID == id {
			user.Username = name
			svc.users[email] = user
			return name, nil
		}
	}
	return "", fmt.Errorf("user with id %s not found", id)
}

func (svc *userApplicationService) List(keyword string) ([]dto.User, error) {
	svc.mu.Lock()
	defer svc.mu.Unlock()

	var ret []dto.User
	for _, user := range svc.users {
		if keyword == "" || strings.Contains(user.Username, keyword) {
			ret = append(ret, user)
		}
	}
	return ret, nil
}