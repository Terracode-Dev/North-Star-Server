-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS V2Preset_Workouts (
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    preset_id BIGINT NOT NULL,
    workout_id BIGINT NOT NULL,
    reps INT NOT NULL DEFAULT 1,
    weight  DECIMAL(10, 2),
    sets INT NOT NULL DEFAULT 1,
    notes TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (preset_id) REFERENCES V2Presets(id) ON DELETE CASCADE,
    FOREIGN KEY (workout_id) REFERENCES Workouts(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS Preset_Workouts;
-- +goose StatementEnd
