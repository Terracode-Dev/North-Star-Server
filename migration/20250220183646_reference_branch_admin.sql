-- +goose Up
-- +goose StatementBegin
ALTER TABLE HR_Admin
ADD CONSTRAINT fk_admin_branch
FOREIGN KEY (branch_id) REFERENCES HR_Branch(id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE HR_Admin
DROP FOREIGN KEY fk_admin_branch;
-- +goose StatementEnd
