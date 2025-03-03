-- +goose Up
-- +goose StatementBegin
ALTER TABLE HR_Tax
    ADD COLUMN updated_by BIGINT,
    ADD CONSTRAINT fk_updated_by FOREIGN KEY (updated_by) REFERENCES HR_Admin(id) ON DELETE SET NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE HR_Tax
    DROP FOREIGN KEY fk_updated_by,
    DROP COLUMN updated_by;
-- +goose StatementEnd
