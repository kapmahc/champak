-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE forum_articles (
  id         SERIAL PRIMARY KEY,
  user_id    BIGINT                      NOT NULL,
  title      VARCHAR(255)                NOT NULL,
  summary    VARCHAR(800)                NOT NULL,
  body       TEXT                        NOT NULL,
  type       VARCHAR(8)                  NOT NULL DEFAULT 'markdown',
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);
CREATE INDEX idx_forum_articles
  ON forum_articles (title);
CREATE INDEX idx_forum_type
  ON forum_articles (type);

CREATE TABLE forum_tags (
  id         SERIAL PRIMARY KEY,
  name       VARCHAR(255)                NOT NULL,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);
CREATE UNIQUE INDEX idx_forum_tags_name
  ON forum_tags (name);

CREATE TABLE forum_articles_tags (
  article_id BIGINT NOT NULL,
  tag_id     BIGINT NOT NULL,
  PRIMARY KEY (article_id, tag_id)
);

CREATE TABLE forum_comments (
  id         SERIAL PRIMARY KEY,
  article_id BIGINT                      NOT NULL,
  user_id    BIGINT                      NOT NULL,
  body       TEXT                        NOT NULL,
  type       VARCHAR(8)                  NOT NULL DEFAULT 'markdown',
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);
CREATE INDEX idx_forum_comments_type
  ON forum_comments (type);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE forum_comments;
DROP TABLE forum_articles_tags;
DROP TABLE forum_tags;
DROP TABLE forum_articles;
