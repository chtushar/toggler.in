package teams

import (
	"context"

	"go.uber.org/zap"
	"toggler.in/internal/db"
	"toggler.in/internal/db/query"
)

type Repository struct {
	q query.Querier
	log *zap.Logger
}

func NewRepository (db *db.DB, log *zap.Logger) *Repository {
	return &Repository{
		q: query.New(db),
		log: log,
	}
}

// AddTeam adds a new Team
func (r *Repository) CreateTeam(ctx context.Context, team query.CreateTeamParams) (*query.CreateTeamRow, error) {
	t, err := r.q.CreateTeam(ctx, team)

	if err != nil {
		r.log.Error("failed to add a new team", zap.Error(err))
		return nil, err
	}

	return &t, nil
}

// // Add team member
// func (r *Repository) AddTeamMember(ctx context.Context, team query.AddTeamMemberParams) (*query.TeamMember, error) {
// 	t, err := r.q.
// }