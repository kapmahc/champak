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

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE notices;
DROP TABLE leave_words;
