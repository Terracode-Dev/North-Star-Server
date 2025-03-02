-- +goose Up
-- +goose StatementBegin
ALTER TABLE HR_EMP_User
ADD CONSTRAINT fk_emp_user_branch
FOREIGN KEY (branch_id) REFERENCES HR_Branch(id);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE HR_EMP_User
DROP FOREIGN KEY fk_emp_user_branch;
-- +goose StatementEnd
