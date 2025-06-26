-- +goose Up
-- +goose StatementBegin
ALTER TABLE HR_Payroll
ADD COLUMN er_id BIGINT,
ADD CONSTRAINT fk_hr_payroll_er_id FOREIGN KEY (er_id) REFERENCES Exchange_Rate(id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE HR_Payroll
DROP COLUMN er_id,
DROP FOREIGN KEY fk_hr_payroll_er_id;
-- +goose StatementEnd
