-- +goose Up
-- +goose StatementBegin
ALTER TABLE V2Preset_Workouts
MODIFY COLUMN reps BIGINT NOT NULL DEFAULT 1,
MODIFY COLUMN sets BIGINT NOT NULL DEFAULT 1
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE V2Preset_Workouts
MODIFY COLUMN reps INT NOT NULL DEFAULT 1,
MODIFY COLUMN sets INT NOT NULL DEFAULT 1
-- +goose StatementEnd
