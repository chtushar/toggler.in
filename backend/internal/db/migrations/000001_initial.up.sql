/**
  Trigger function to set the updated_at field to current time
  when a row is updated.
 */
CREATE FUNCTION set_updated_at() RETURNS TRIGGER
    LANGUAGE plpgsql
AS
$BODY$
BEGIN
    IF TG_OP = 'UPDATE' THEN
        NEW."updated_at" = NOW();
    END IF;
    RETURN NEW;
END;
$BODY$;


-- Users table
CREATE TABLE IF NOT EXISTS users (
  id            SERIAL PRIMARY KEY NOT NULL,
  name          VARCHAR(255)       NOT NULL,
  password      VARCHAR(255)       NOT NULL,
  email         VARCHAR(255)       NOT NULL UNIQUE,
  email_verified BOOLEAN            NOT NULL DEFAULT false,
  created_at    TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
  updated_at    TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);

-- Updated at trigger for users
CREATE TRIGGER set_updated_at_users
    BEFORE UPDATE
    ON users
    FOR EACH ROW
EXECUTE FUNCTION set_updated_at();

-- Teams table
CREATE TABLE IF NOT EXISTS teams (
    id            SERIAL PRIMARY KEY NOT NULL,
    name          VARCHAR(255)       NOT NULL,
    created_at    TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at    TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);

-- Updated at trigger for team
CREATE TRIGGER set_updated_at_teams
    BEFORE UPDATE
    ON teams
    FOR EACH ROW
EXECUTE FUNCTION set_updated_at();

-- Users teams junction table
CREATE TABLE IF NOT EXISTS user_team (
    user_id       INTEGER            NOT NULL,
    team_id       INTEGER            NOT NULL,
    created_at    TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at    TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);