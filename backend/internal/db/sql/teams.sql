-- name: CreateTeam :one
INSERT INTO teams (name, owner_id)
VALUES ($1, $2)
RETURNING *;

-- name: GetTeam :one
SELECT * FROM teams
WHERE id = $1;

-- name: AddTeamMember :one
INSERT INTO team_members (user_id, team_id)
VALUES ($1, $2)
RETURNING *;

-- name: RemoveTeamMember :exec
DELETE FROM team_members WHERE user_id = $1 AND team_id = $2;

-- name: GetTeamMembers :many
SELECT * FROM team_members WHERE team_id = $1;

