-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS Exchange_Rate (
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    exchange_rate DECIMAL(10, 4) NOT NULL,
    currency_type VARCHAR(10) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS Exchange_Rate;
-- +goose StatementEnd
