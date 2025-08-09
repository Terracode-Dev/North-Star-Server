-- +goose Up
-- +goose StatementBegin
ALTER TABLE HR_EMP_Benifits 
    MODIFY COLUMN leave_type VARCHAR(100),
    MODIFY COLUMN leave_count INT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE HR_EMP_Benifits 
    MODIFY COLUMN leave_type VARCHAR(100) NOT NULL,
    MODIFY COLUMN leave_count INT NOT NULL;
-- +goose StatementEnd