package auth

import (
	"context"

	"go.uber.org/zap"
	"toggler.in/internal/db"
	"toggler.in/internal/db/query"
)

type Repository struct {
	q *query.Queries
	log *zap.Logger
}

//NewRepository creates a new instance of Repository
func NewRepository(db *db.DB, log *zap.Logger) *Repository {
	return &Repository{
		q: query.New(db),
		log: log,
	}
}

func (r *Repository) GetUserByEmail(ctx context.Context, email string) (*query.User, error) {
	u, err := r.q.GetUserByEmail(ctx, email)

	if err != nil {
		r.log.Error("failed to get user by email", zap.Error(err))
		return nil, err
	}

	return &u, nil
}