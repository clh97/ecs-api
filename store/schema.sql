CREATE TABLE IF NOT EXISTS ecs_user (
  id serial PRIMARY KEY,
  first_name text NOT NULL,
  last_name text,
  email text NOT NULL UNIQUE,
  password text NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW()
);