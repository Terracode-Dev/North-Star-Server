-- +goose Up
-- +goose StatementBegin
ALTER TABLE HR_Trainer_Emp
MODIFY COLUMN commission DECIMAL(10, 2) NOT NULL DEFAULT 0.00;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE HR_Trainer_Emp
DROP COLUMN IF EXISTS commission;
-- +goose StatementEnd
