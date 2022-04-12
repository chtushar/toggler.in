drop trigger if exists set_updated_at_users on users cascade;

drop table if exists users cascade;

drop trigger if exists set_updated_at_teams on team cascade;

drop table if exists team cascade;

drop table if exists user_team;

