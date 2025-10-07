-- +goose Up
-- +goose StatementBegin
ALTER TABLE emp_link DROP FOREIGN KEY emp_link_ibfk_2;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE emp_link ADD CONSTRAINT emp_link_ibfk_2 FOREIGN KEY (updated_by) REFERENCES HR_Employee (id);
-- +goose StatementEnd
