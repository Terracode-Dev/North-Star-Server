-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS Workouts (
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    animation_url TEXT,
    preview_animation_url TEXT,
    categories LONGTEXT,
    optional LONGTEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS Workouts;
-- +goose StatementEnd
