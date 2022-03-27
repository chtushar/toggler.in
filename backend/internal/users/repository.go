package users

import (
	"context"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"toggler.in/internal/models"
)

//Repository has CRUD functions for users
type Repository struct {
	db *gorm.DB
	log *zap.Logger
}

//NewRepository creates a new instance of Repository
func NewRepository(db *gorm.DB, log *zap.Logger) *Repository {
	return &Repository{
		db: db,
		log: log,
	}
}

//AddUser adds a new User
func (r *Repository) AddUser(ctx context.Context, user models.AddUserParams) (*models.User, error) {
	u := &models.User{
		Name: user.Name,
		Email: user.Email,
		Password: user.Password,
	}
	result := r.db.Create(u)

	if result.Error != nil {
		r.log.Error("failed to add a new user", zap.Error(result.Error))
		return nil, result.Error
	}

	return u, nil
}