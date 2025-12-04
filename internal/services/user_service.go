package services

import "github.com/mrxacker/user_service/internal/models"

type UserRepository interface {
	CreateUser(user models.User) (models.User, error)
	GetUserByID(id int) (models.User, error)
	GetUsers() ([]models.User, error)
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) RegisterUser(name, email string) (models.User, error) {
	user := models.User{
		Name:  name,
		Email: email,
	}
	return s.repo.CreateUser(user)
}

func (s *UserService) GetUser(id int) (models.User, error) {
	return s.repo.GetUserByID(id)
}

func (s *UserService) GetUsers() ([]models.User, error) {
	return s.repo.GetUsers()
}
