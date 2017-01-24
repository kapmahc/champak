CREATE SEQUENCE hibernate_sequences;

CREATE TABLE users (
  id                 SERIAL PRIMARY KEY,
  full_name          VARCHAR(32)                 NOT NULL,
  email              VARCHAR(255)                NOT NULL,
  uid                VARCHAR(36)                 NOT NULL,
  password           BYTEA,
  provider_id        VARCHAR(255)                NOT NULL,
  provider_type      VARCHAR(32)                 NOT NULL,
  home               VARCHAR(255),
  logo               VARCHAR(255),
  sign_in_count      INT                         NOT NULL DEFAULT 0,
  current_sign_in_at TIMESTAMP WITHOUT TIME ZONE,
  current_sign_in_ip INET,
  last_sign_in_at    TIMESTAMP WITHOUT TIME ZONE,
  last_sign_in_ip    INET,
  confirmed_at       TIMESTAMP WITHOUT TIME ZONE,
  locked_at          TIMESTAMP WITHOUT TIME ZONE,
  created_at         TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
  updated_at         TIMESTAMP WITHOUT TIME ZONE NOT NULL
);
CREATE UNIQUE INDEX idx_users_uid
  ON users (uid);
CREATE UNIQUE INDEX idx_users_email
  ON users (email);
CREATE UNIQUE INDEX idx_users_provider_id_type
  ON users (provider_id, provider_type);
CREATE INDEX idx_users_full_name
  ON users (full_name);
CREATE INDEX idx_users_provider_id
  ON users (provider_id);
CREATE INDEX idx_users_provider_type
  ON users (provider_type);


CREATE TABLE contacts (
  id         SERIAL PRIMARY KEY,
  user_id    BIGINT                      NOT NULL,
  key        VARCHAR(32)                 NOT NULL,
  val        VARCHAR(255)                NOT NULL,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);
CREATE UNIQUE INDEX idx_contacts
  ON contacts (user_id, key);
CREATE INDEX idx_contacts_key
  ON contacts (key);

CREATE TABLE logs (
  id         SERIAL PRIMARY KEY,
  user_id    BIGINT                      NOT NULL,
  type       VARCHAR(8)                  NOT NULL DEFAULT 'info',
  message    VARCHAR(255)                NOT NULL,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now()
);
CREATE INDEX idx_logs_type
  ON logs (type);

CREATE TABLE roles (
  id            SERIAL PRIMARY KEY,
  name          VARCHAR(32)                 NOT NULL,
  resource_id   BIGINT,
  resource_type VARCHAR(255),
  created_at    TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
  updated_at    TIMESTAMP WITHOUT TIME ZONE NOT NULL
);
CREATE UNIQUE INDEX idx_roles_name_resource_type_id
  ON roles (name, resource_type, resource_id);
CREATE INDEX idx_roles_name
  ON roles (name);
CREATE INDEX idx_roles_resource_type
  ON roles (resource_type);

CREATE TABLE policies (
  id         SERIAL PRIMARY KEY,
  user_id    BIGINT                      NOT NULL,
  role_id    BIGINT                      NOT NULL,
  start_up   DATE                        NOT NULL DEFAULT current_date,
  shut_down  DATE                        NOT NULL DEFAULT '2016-12-13',
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);
CREATE UNIQUE INDEX idx_policies
  ON policies (user_id, role_id);

CREATE TABLE settings (
  id         SERIAL PRIMARY KEY,
  key        VARCHAR(255)                NOT NULL,
  user_id    BIGINT,
  val        BYTEA                       NOT NULL,
  flag       BOOLEAN                     NOT NULL DEFAULT FALSE,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);
CREATE UNIQUE INDEX idx_settings
  ON settings (key, user_id);
CREATE INDEX idx_settings_key
  ON settings (key);

CREATE TABLE votes (
  id            SERIAL PRIMARY KEY,
  resource_type VARCHAR(255)                NOT NULL,
  resource_id   BIGINT                      NOT NULL,
  point         INT                         NOT NULL DEFAULT 0,
  created_at    TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
  updated_at    TIMESTAMP WITHOUT TIME ZONE NOT NULL
);
CREATE UNIQUE INDEX idx_votes_resources
  ON votes (resource_type, resource_id);
CREATE INDEX idx_votes_resource_type
  ON votes (resource_type);

CREATE TABLE locales (
  id         SERIAL PRIMARY KEY,
  code       VARCHAR(255)                NOT NULL,
  lang       VARCHAR(8)                  NOT NULL DEFAULT 'en-US',
  message    TEXT                        NOT NULL,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);
CREATE UNIQUE INDEX idx_locales_code_lang
  ON locales (code, lang);
CREATE INDEX idx_locales_code
  ON locales (code);
CREATE INDEX idx_locales_lang
  ON locales (lang);

CREATE TABLE notices (
  id         SERIAL PRIMARY KEY,
  body       TEXT                        NOT NULL,
  type       VARCHAR(8)                  NOT NULL DEFAULT 'markdown',
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);

CREATE TABLE leave_words (
  id         SERIAL PRIMARY KEY,
  body       TEXT                        NOT NULL,
  type       VARCHAR(8)                  NOT NULL DEFAULT 'markdown',
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now()
);

CREATE TABLE attachments (
  id            SERIAL PRIMARY KEY,
  title         VARCHAR(255)                NOT NULL,
  url           VARCHAR(255)                NOT NULL,
  length        INT                         NOT NULL,
  media_type    VARCHAR(32)                 NOT NULL,
  resource_type VARCHAR(255)                NOT NULL,
  resource_id   BIGINT                      NOT NULL,
  user_id       BIGINT                      NOT NULL,
  sort_order    INT                         NOT NULL DEFAULT 0,
  created_at    TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now()
);
CREATE UNIQUE INDEX idx_attachments_url
  ON attachments (url);
CREATE UNIQUE INDEX idx_attachments_resources
  ON attachments (resource_type, resource_id);
CREATE INDEX idx_attachments_title
  ON attachments (title);
CREATE INDEX idx_attachments_media_type
  ON attachments (media_type);