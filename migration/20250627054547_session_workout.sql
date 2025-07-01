-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS Session_Workout (
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    preset_session_id BIGINT NOT NULL,
    active_day BIGINT NOT NULL,
    workout_id BIGINT NOT NULL,
    status VARCHAR(50) ,
    session_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (preset_session_id) REFERENCES V2Preset_Session(id) ON DELETE CASCADE,
    FOREIGN KEY (workout_id) REFERENCES Workouts(id) ON DELETE CASCADE,
    FOREIGN KEY (session_id) REFERENCES V2Session(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS Session_Workout;
-- +goose StatementEnd
