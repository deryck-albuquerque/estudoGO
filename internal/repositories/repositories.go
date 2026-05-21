package repositories

import (
	"github.com/deryck-albuquerque/estudoGO/internal/models"
	"github.com/deryck-albuquerque/estudoGO/internal/repositories/users"
)

type Repositories struct {
	User interface {
		GetAll() []models.User
		Add(newUser models.User)
		EmailExists(email string) bool
	}
}

func New() *Repositories {
	return &Repositories{
		User: users.New(),
	}
}
