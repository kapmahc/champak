-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE cms_articles (
  id         SERIAL PRIMARY KEY,
  user_id    BIGINT                      NOT NULL,
  title      VARCHAR(255)                NOT NULL,
  summary    VARCHAR(800)                NOT NULL,
  body       TEXT                        NOT NULL,
  type       VARCHAR(8)                  NOT NULL DEFAULT 'markdown',
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);
CREATE INDEX idx_cms_articles
  ON cms_articles (title);
CREATE INDEX idx_cms_type
  ON cms_articles (type);

CREATE TABLE cms_tags (
  id         SERIAL PRIMARY KEY,
  name       VARCHAR(255)                NOT NULL,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);
CREATE UNIQUE INDEX idx_cms_tags_name
  ON cms_tags (name);

CREATE TABLE cms_articles_tags (
  id         SERIAL PRIMARY KEY,
  article_id BIGINT NOT NULL,
  tag_id     BIGINT NOT NULL
);
CREATE UNIQUE INDEX idx_cms_articles_tags
  ON cms_articles_tags (article_id, tag_id);

CREATE TABLE cms_comments (
  id         SERIAL PRIMARY KEY,
  article_id BIGINT                      NOT NULL,
  user_id    BIGINT                      NOT NULL,
  body       TEXT                        NOT NULL,
  type       VARCHAR(8)                  NOT NULL DEFAULT 'markdown',
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);
CREATE INDEX idx_cms_comments_type
  ON cms_comments (type);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE cms_comments;
DROP TABLE cms_articles_tags;
DROP TABLE cms_tags;
DROP TABLE cms_articles;
