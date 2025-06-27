-- +goose Up
-- +goose StatementBegin
ALTER TABLE HR_Trainer_Emp
ADD COLUMN updated_by BIGINT,
ADD FOREIGN KEY (updated_by) REFERENCES HR_Admin(id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE HR_Trainer_Emp
DROP COLUMN updated_by,
DROP FOREIGN KEY HR_Trainer_Com_updated_by_fk;
-- +goose StatementEnd
