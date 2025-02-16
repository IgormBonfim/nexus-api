CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  public_key UUID,
  username VARCHAR(50) UNIQUE NOT NULL CHECK (char_length(username) >= 3),
  hashed_password TEXT NOT NULL,
  email VARCHAR(255) UNIQUE NOT NULL,
  created_at TIMESTAMP DEFAULT NOW() NOT NULL,
  updated_at TIMESTAMP DEFAULT NOW() NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_users_publickey ON users (public_key);
CREATE INDEX IF NOT EXISTS idx_users_email ON users (email);
CREATE INDEX IF NOT EXISTS idx_users_username ON users (username);

CREATE TABLE IF NOT EXISTS posts (
  id SERIAL PRIMARY KEY,
  public_key UUID,
  user_id INTEGER NOT NULL,
  content TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT NOW() NOT NULL,
  updated_at TIMESTAMP DEFAULT NOW() NOT NULL,
  CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id)
)

CREATE INDEX IF NOT EXISTS idx_posts_publickey ON posts (public_key);