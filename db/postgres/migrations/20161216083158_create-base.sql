-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE locales (
  id SERIAL PRIMARY KEY,
  code VARCHAR(255) NOT NULL,
  lang VARCHAR(8) NOT NULL DEFAULT 'en-US',
  message TEXT NOT NULL,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);
CREATE UNIQUE INDEX idx_locales_code_lang ON locales (code, lang);
CREATE INDEX idx_locales_code ON locales (code);
CREATE INDEX idx_locales_lang ON locales (lang);


CREATE TABLE settings (
  id SERIAL PRIMARY KEY,
  key VARCHAR(255) NOT NULL,
  val BYTEA NOT NULL,
  flag BOOLEAN NOT NULL DEFAULT FALSE,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);
CREATE UNIQUE INDEX idx_settings_key ON settings (key);

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
DROP TABLE settings;
DROP TABLE locales;
