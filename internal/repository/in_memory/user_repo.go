package in_memory

import (
	"errors"
	"sync"

	"github.com/mrxacker/user_service/internal/models"
)

type InMemoryUserRepo struct {
	users  []models.User
	nextID int
	mu     sync.RWMutex
}

func NewInMemoryUserRepo() *InMemoryUserRepo {
	return &InMemoryUserRepo{
		users:  make([]models.User, 0),
		nextID: 1,
	}
}

func (repo *InMemoryUserRepo) GetAll() ([]models.User, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	return append([]models.User(nil), repo.users...), nil
}

func (repo *InMemoryUserRepo) GetByID(id int) (*models.User, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	for _, user := range repo.users {
		if user.ID == id {

			c := user
			return &c, nil
		}
	}
	return nil, errors.New("user not found")
}

func (repo *InMemoryUserRepo) Create(user models.User) (models.User, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	user.ID = repo.nextID
	repo.nextID++
	repo.users = append(repo.users, user)
	return user, nil
}

func (repo *InMemoryUserRepo) Update(user models.User) (*models.User, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	for i, u := range repo.users {
		if u.ID == user.ID {
			repo.users[i] = user
			c := user
			return &c, nil
		}
	}
	return nil, errors.New("user not found")
}

func (repo *InMemoryUserRepo) Delete(id int) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	for i, user := range repo.users {
		if user.ID == id {
			repo.users = append(repo.users[:i], repo.users[i+1:]...)
			return nil
		}
	}

	return errors.New("user not found")
}
