-- +goose Up
-- +goose StatementBegin
CREATE TABLE HR_Admin(
  id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  user_name VARCHAR(150) NOT NULL,
  email VARCHAR(150) NOT NULL,
  password VARCHAR(100) NOT NULL,
  role VARCHAR(100) NOT NULL,
  status VARCHAR(100) NOT NULL,
  branch_id BIGINT NOT NULL,
  created_by BIGINT,
  updated_by BIGINT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE HR_Admin
-- +goose StatementEnd
