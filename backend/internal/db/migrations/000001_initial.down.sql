DROP TRIGGER IF EXISTS set_updated_at_users ON users CASCADE;

DROP TABLE IF EXISTS users CASCADE;

DROP TRIGGER IF EXISTS set_updated_at_teams ON team CASCADE;

DROP TABLE IF EXISTS teams CASCADE;

DROP TABLE IF EXISTS team_members;

DROP TABLE IF EXISTS ff_resolution_boolean;

DROP TABLE IF EXISTS feature_flags;

DROP TABLE IF EXISTS feature_flag_types;

