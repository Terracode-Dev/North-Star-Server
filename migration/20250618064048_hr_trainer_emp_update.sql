-- +goose Up
-- +goose StatementBegin
ALTER TABLE HR_Trainer_Emp
ADD COLUMN commission DECIMAL(10, 2) DEFAULT 0.00;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE HR_Trainer_Emp
DROP COLUMN IF EXISTS commission;
-- +goose StatementEnd
