-- +goose Up
-- +goose StatementBegin
ALTER TABLE HR_EMP_Certificates DROP COLUMN image_path;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE HR_EMP_Certificates ADD COLUMN image_path TEXT NOT NULL;
-- +goose StatementEnd
