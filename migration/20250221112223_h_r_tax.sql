-- +goose Up
-- +goose StatementBegin
CREATE TABLE HR_Tax(
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    tax_from DECIMAL(10,2) NOT NULL,
    tax_to DECIMAL(10,2) NOT NULL,
    tax_percentage DECIMAL(10,2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE HR_Tax;
-- +goose StatementEnd
