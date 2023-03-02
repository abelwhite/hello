CREATE TABLE IF NOT EXISTS users (
  user_id serial PRIMARY KEY,
  name text NOT NULL,
  email citext UNIQUE NOT NULL,
  created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW()
);
