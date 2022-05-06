package users

import (
	"context"

	"go.uber.org/zap"
	"toggler.in/internal/db"
	"toggler.in/internal/db/query"
)

//Repository has CRUD functions for users
type Repository struct {
	q query.Querier
	log *zap.Logger
}

//NewRepository creates a new instance of Repository
func NewRepository(db *db.DB, log *zap.Logger) *Repository {
	return &Repository{
		q: query.New(db),
		log: log,
	}
}

//AddUser adds a new User
func (r *Repository) AddUser(ctx context.Context, user query.AddUserParams) (*query.User, error) {
	u, err := r.q.AddUser(ctx, user)

	if err != nil {
		r.log.Error("failed to add a new user", zap.Error(err))
		return nil, err
	}

	return &u, nil
}

func (r *Repository) GetUserByEmail(ctx context.Context, email string) (*query.User, error) {
	u, err := r.q.GetUserByEmail(ctx, email)

	if err != nil {
		r.log.Error("failed to get user by email", zap.Error(err))
		return nil, err
	}

	return &u, nil
}

//GetUser gets a User by ID
func (r *Repository) GetUser(ctx context.Context, id int32) (*query.User, error) {
	u, err := r.q.GetUser(ctx, id)

	if err != nil {
		r.log.Error("failed to get user", zap.Error(err))
		return nil, err
	}

	return &u, nil
}

//VerifyEmail verifies an email
func (r *Repository) VerifyEmail(ctx context.Context, id int32) (*query.User, error) {
	u, err := r.q.VerifyEmail(ctx, id)

	if err != nil {
		r.log.Error("failed to verify email", zap.Error(err))
		return nil, err
	}

	return &u, nil
}