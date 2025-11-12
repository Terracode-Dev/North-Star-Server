-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS emp_confirmation_letter_table(
    id BIGINT(20) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    emp_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE TABLE IF EXISTS emp_confirmation_letter_table;
-- +goose StatementEnd
