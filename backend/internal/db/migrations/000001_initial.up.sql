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

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

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
    owner_id 	  INTEGER REFERENCES users (id) ON UPDATE CASCADE NOT NULL,
    created_at    TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at    TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);

-- Updated at trigger for team
CREATE TRIGGER set_updated_at_teams
    BEFORE UPDATE
    ON teams
    FOR EACH ROW
EXECUTE FUNCTION set_updated_at();

-- Team members junction table
CREATE TABLE IF NOT EXISTS team_members (
    user_id       INTEGER REFERENCES users (id) ON UPDATE CASCADE ON DELETE CASCADE NOT NULL,
    team_id       INTEGER            NOT NULL,
    created_at    TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at    TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);

-- Updated at trigger for team members
CREATE TRIGGER set_updated_at_team_members
    BEFORE UPDATE
    ON team_members
    FOR EACH ROW
EXECUTE FUNCTION set_updated_at();

-- Feature Flag Types table
CREATE TABLE IF NOT EXISTS feature_flag_types (
    id            SERIAL PRIMARY KEY NOT NULL,
    type          VARCHAR(255)       NOT NULL
);

-- Feature Flags table
CREATE TABLE IF NOT EXISTS feature_flags (
    id            SERIAL PRIMARY KEY NOT NULL,
    uuid          UUID   DEFAULT uuid_generate_v1 (),
    name          VARCHAR(255)       NOT NULL,
    type          INTEGER            NOT NULL,
    team_id       INTEGER REFERENCES teams (id) ON UPDATE CASCADE ON DELETE CASCADE NOT NULL,
    created_at    TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at    TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);

-- Updated at trigger for team members
CREATE TRIGGER set_updated_at_feature_flags
    BEFORE UPDATE
    ON feature_flags
    FOR EACH ROW
EXECUTE FUNCTION set_updated_at();

-- Feature Flag resolution for boolean
CREATE TABLE IF NOT EXISTS ff_resolution_boolean (
	id 			  BIGSERIAL PRIMARY KEY NOT NULL,
  	key 		  VARCHAR(255) NOT NULL,
  	ff_id 		  INTEGER REFERENCES feature_flags (id) ON UPDATE CASCADE ON DELETE CASCADE NOT NULL,
  	enabled       BOOLEAN DEFAULT false,
    UNIQUE(KEY, ff_id)
);
