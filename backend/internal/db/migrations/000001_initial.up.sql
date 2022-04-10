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
create table IF NOT EXISTS users (
  id            serial primary key not null,
  name          varchar(255)       not null,
  password      varchar(255)       not null,
  email         varchar(255)       not null,
  email_verified boolean            not null default false,
  created_at    timestamp with time zone not null default now(),
  updated_at    timestamp with time zone not null default now()
);

-- Updated at trigger for users
CREATE TRIGGER set_updated_at_users
    BEFORE UPDATE
    ON users
    FOR EACH ROW
EXECUTE FUNCTION set_updated_at();

