-- +goose Up
-- +goose StatementBegin
ALTER TABLE HR_EMP_Expatriate DROP COLUMN visa_image_path;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE HR_EMP_Expatriate ADD COLUMN visa_image_path TEXT NOT NULL;
-- +goose StatementEnd
