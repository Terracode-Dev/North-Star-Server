-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS V2Preset_Session (
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    preset_id BIGINT NOT NULL,
    client_id BIGINT NOT NULL,
    assign_session BIGINT NOT NULL,
    active_day BIGINT NOT NULL,
    max_day BIGINT NOT NULL,
    state VARCHAR(50) NOT NULL DEFAULT 'active',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (preset_id) REFERENCES V2Presets(id) ON DELETE CASCADE
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS V2Preset_Session;
-- +goose StatementEnd
