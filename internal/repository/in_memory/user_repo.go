package in_memory

import (
	"errors"

	"github.com/mrxacker/user_service/internal/models"
)

type InMemoryUserRepo struct {
	users  []models.User
	nextID int
}

func NewInMemoryUserRepo() *InMemoryUserRepo {
	return &InMemoryUserRepo{
		users:  make([]models.User, 0),
		nextID: 1,
	}
}

func (repo *InMemoryUserRepo) CreateUser(user models.User) (models.User, error) {
	user.ID = repo.nextID
	repo.nextID++
	repo.users = append(repo.users, user)
	return user, nil
}

func (repo *InMemoryUserRepo) GetUserByID(id int) (models.User, error) {
	for _, user := range repo.users {
		if user.ID == id {
			return user, nil
		}
	}
	return models.User{}, errors.New("user not found")
}
