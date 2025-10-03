-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS Admin_Presets (
    id BIGINT(20) NOT NULL AUTO_INCREMENT,
    preset_name VARCHAR(255) NOT NULL,
    preset_value JSON NOT NULL,
    slug VARCHAR(100) NOT NULL UNIQUE,
    PRIMARY KEY (id)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS Admin_Presets;
-- +goose StatementEnd
