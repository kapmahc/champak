-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE t1 (
  ID SERIAL PRIMARY KEY
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE t1;