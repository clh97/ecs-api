CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS ecs_app (
  id serial PRIMARY KEY,
  owner_id int NOT NULL,
  url_id uuid DEFAULT uuid_generate_v4(),
  name text NOT NULL,
  pages int[],
  url text,
  created_at timestamp NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS ecs_page (
  id serial PRIMARY KEY,
  page_id uuid DEFAULT uuid_generate_v4(),
  app_id uuid NOT NULL,
  title text NOT NULL,
  slug text,
  url text,
  created_at timestamp NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS ecs_user (
  id serial PRIMARY KEY,
  first_name text NOT NULL,
  last_name text,
  email text NOT NULL UNIQUE,
  password text NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS ecs_anon (
  id serial PRIMARY KEY,
  username text NOT NULL,
  ip text NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS ecs_comment (
  id serial PRIMARY KEY, /* ecs_user.id or ecs_anon.id */
  app_id serial, /* ecs_app.id */
  page_id serial, /* ecs_app.pages[x].id */
  user_id serial,
  content text NOT NULL,
  content_format text NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW()
);