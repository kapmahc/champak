-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE leave_words (
  id         SERIAL PRIMARY KEY,
  body       TEXT                        NOT NULL,
  type       VARCHAR(8)                  NOT NULL DEFAULT 'markdown',
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now()
);

CREATE TABLE notices (
  id         SERIAL PRIMARY KEY,
  body       TEXT                        NOT NULL,
  type       VARCHAR(8)                  NOT NULL DEFAULT 'markdown',
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);

CREATE TABLE links (
  id SERIAL PRIMARY KEY,
  href VARCHAR(255) NOT NULL,
  label VARCHAR(255) NOT NULL,
  loc VARCHAR(16) NOT NULL,
  sort_order INT NOT NULL DEFAULT 0,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);
CREATE UNIQUE INDEX idx_links_loc ON links (loc);

CREATE TABLE cards (
  id SERIAL PRIMARY KEY,
  href VARCHAR(255) NOT NULL,
  logo VARCHAR(255) NOT NULL,
  title VARCHAR(64) NOT NULL,
  summary  VARCHAR(500) NOT NULL,
  loc VARCHAR(16) NOT NULL,
  sort_order INT NOT NULL DEFAULT 0,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);
CREATE UNIQUE INDEX idx_cards_loc ON cards (loc);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE cards;
DROP TABLE links;
DROP TABLE notices;
DROP TABLE leave_words;
