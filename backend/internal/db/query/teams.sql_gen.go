// Code generated by sqlc. DO NOT EDIT.
// source: teams.sql

package query

import (
	"context"
)

const addTeamMember = `-- name: AddTeamMember :one
INSERT INTO team_members (user_id, team_id)
VALUES ($1, $2)
RETURNING user_id, team_id, created_at, updated_at
`

type AddTeamMemberParams struct {
	UserID int32 `db:"user_id"`
	TeamID int32 `db:"team_id"`
}

func (q *Queries) AddTeamMember(ctx context.Context, arg AddTeamMemberParams) (TeamMember, error) {
	row := q.db.QueryRow(ctx, addTeamMember, arg.UserID, arg.TeamID)
	var i TeamMember
	err := row.Scan(
		&i.UserID,
		&i.TeamID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createTeam = `-- name: CreateTeam :one
INSERT INTO teams (name, owner_id)
VALUES ($1, $2)
RETURNING id, name, owner_id, created_at, updated_at
`

type CreateTeamParams struct {
	Name    string `db:"name"`
	OwnerID int32  `db:"owner_id"`
}

func (q *Queries) CreateTeam(ctx context.Context, arg CreateTeamParams) (Team, error) {
	row := q.db.QueryRow(ctx, createTeam, arg.Name, arg.OwnerID)
	var i Team
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.OwnerID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getTeam = `-- name: GetTeam :one
SELECT id, name, owner_id, created_at, updated_at FROM teams
WHERE id = $1
`

func (q *Queries) GetTeam(ctx context.Context, id int32) (Team, error) {
	row := q.db.QueryRow(ctx, getTeam, id)
	var i Team
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.OwnerID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getTeamMembers = `-- name: GetTeamMembers :many
SELECT user_id, team_id, created_at, updated_at FROM team_members WHERE team_id = $1
`

func (q *Queries) GetTeamMembers(ctx context.Context, teamID int32) ([]TeamMember, error) {
	rows, err := q.db.Query(ctx, getTeamMembers, teamID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TeamMember
	for rows.Next() {
		var i TeamMember
		if err := rows.Scan(
			&i.UserID,
			&i.TeamID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const removeTeamMember = `-- name: RemoveTeamMember :exec
DELETE FROM team_members WHERE user_id = $1 AND team_id = $2
`

type RemoveTeamMemberParams struct {
	UserID int32 `db:"user_id"`
	TeamID int32 `db:"team_id"`
}

func (q *Queries) RemoveTeamMember(ctx context.Context, arg RemoveTeamMemberParams) error {
	_, err := q.db.Exec(ctx, removeTeamMember, arg.UserID, arg.TeamID)
	return err
}
