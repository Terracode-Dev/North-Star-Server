-- +goose Up
-- +goose StatementBegin
ALTER TABLE HR_EMP_Benifits
ADD COLUMN ticket boolean,
ADD COLUMN ticket_quantity int DEFAULT 0;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE HR_EMP_Benifits
DROP COLUMN ticket,
DROP COLUMN ticket_quantity;
-- +goose StatementEnd
