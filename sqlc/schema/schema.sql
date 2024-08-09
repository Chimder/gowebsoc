CREATE TABLE
  users (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT now (),
    updated_at TIMESTAMP NOT NULL DEFAULT now (),
    name VARCHAR(255) UNIQUE NOT NULL
    
  );

CREATE TABLE
  channels (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT now (),
    updated_at TIMESTAMP NOT NULL DEFAULT now (),
    name VARCHAR(255) UNIQUE NOT NULL
  );

CREATE TABLE
  podchannels (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT now (),
    updated_at TIMESTAMP NOT NULL DEFAULT now (),
    name VARCHAR(255) NOT NULL,
    types VARCHAR(50) NOT NULL,
    channel_id INT NOT NULL,
    FOREIGN KEY (channel_id) REFERENCES channels (id)
  );

CREATE TABLE
  messages (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT now (),
    updated_at TIMESTAMP NOT NULL DEFAULT now (),
    content TEXT NOT NULL,
    author_id VARCHAR(50) NOT NULL,
    podchannel_id INT NOT NULL,
    FOREIGN KEY (podchannel_id) REFERENCES podchannels (id)
  );

CREATE INDEX idx_messages_podchannel_id ON messages (podchannel_id);

CREATE INDEX idx_messages_author_id ON messages (author_id);