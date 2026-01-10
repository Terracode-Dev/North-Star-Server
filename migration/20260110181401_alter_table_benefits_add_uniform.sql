-- +goose Up
-- +goose StatementBegin
ALTER TABLE HR_EMP_Benifits 
ADD COLUMN uniform boolean,
ADD COLUMN uniform_quantity INT, 
ADD COLUMN uniform_renew_months INT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE HR_EMP_Benifits 
DROP COLUMN uniform,
DROP COLUMN uniform_quantity,
DROP COLUMN uniform_renew_months;
-- +goose StatementEnd
