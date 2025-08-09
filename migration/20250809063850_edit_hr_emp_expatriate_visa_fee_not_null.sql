-- +goose Up
-- +goose StatementBegin
ALTER TABLE HR_EMP_Expatriate 
    MODIFY COLUMN visa_fee DECIMAL(10,2) NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE HR_EMP_Expatriate 
    MODIFY COLUMN visa_fee DECIMAL(10,2);
-- +goose StatementEnd