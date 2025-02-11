-- +goose Up
-- +goose StatementBegin
CREATE TABLE HR_Admin(
  id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  Name TEXT NOT NULL,
  Email TEXT NOT NULL,
  Role TEXT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE HR_Admin
-- +goose StatementEnd
