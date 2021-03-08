CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS ecs_app (
  id serial PRIMARY KEY,
  owner_id int NOT NULL,
  url_id uuid DEFAULT uuid_generate_v4() UNIQUE,
  name text NOT NULL UNIQUE,
  pages int[],
  url text,
  created_at timestamp NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS ecs_page (
  id serial PRIMARY KEY,
  page_id uuid DEFAULT uuid_generate_v4() UNIQUE,
  app_id uuid NOT NULL,
  title text NOT NULL,
  slug text,
  url text,
  created_at timestamp NOT NULL DEFAULT NOW(),
  CONSTRAINT fk_page_app
    FOREIGN KEY (app_id)
      REFERENCES ecs_app(url_id)
);

CREATE TABLE IF NOT EXISTS ecs_user (
  id serial PRIMARY KEY,
  username text NOT NULL UNIQUE,
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
  app_id uuid NOT NULL, /* ecs_app.id */
  page_id uuid NOT NULL, /* ecs_app.pages[x].id */
  user_id int DEFAULT 0,
  content text NOT NULL,
  content_format text NOT NULL DEFAULT 'markdown',
  anon boolean DEFAULT true, 
  anon_username text,
  username text,
  created_at timestamp NOT NULL DEFAULT NOW(),
  CONSTRAINT fk_comment_app
    FOREIGN KEY (app_id)
      REFERENCES ecs_app(url_id),
  CONSTRAINT fk_comment_page
    FOREIGN KEY (page_id)
      REFERENCES ecs_page(page_id),
  CONSTRAINT fk_comment_user
    FOREIGN KEY (user_id)
      REFERENCES ecs_user(id)
);

INSERT INTO ecs_user (username, email, password) VALUES ('<anonymous>', 'anon@calheiros.dev', '<><><>') ON CONFLICT DO NOTHING