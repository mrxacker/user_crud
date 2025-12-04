package services

import "github.com/mrxacker/user_service/internal/models"

type UserRepository interface {
	Create(user models.User) (models.User, error)
	GetByID(id int) (*models.User, error)
	GetAll() ([]models.User, error)
	Update(user models.User) (*models.User, error)
	Delete(id int) error
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
	return s.repo.Create(user)
}

func (s *UserService) GetUser(id int) (*models.User, error) {
	return s.repo.GetByID(id)
}

func (s *UserService) GetUsers() ([]models.User, error) {
	return s.repo.GetAll()
}

func (s *UserService) UpdateUser(id int, name, email string) (*models.User, error) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	user.Name = name
	user.Email = email

	return s.repo.Update(*user)
}

func (s *UserService) DeleteUser(id int) error {
	return s.repo.Delete(id)
}
