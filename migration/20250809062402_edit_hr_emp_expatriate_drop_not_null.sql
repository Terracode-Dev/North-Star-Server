-- +goose Up
-- +goose StatementBegin
ALTER TABLE HR_EMP_Expatriate 
    MODIFY COLUMN nationality VARCHAR(200),
    MODIFY COLUMN visa_type VARCHAR(100),
    MODIFY COLUMN visa_from DATE,
    MODIFY COLUMN visa_till DATE,
    MODIFY COLUMN visa_number VARCHAR(100),
    MODIFY COLUMN visa_fee DECIMAL(10,2);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE HR_EMP_Expatriate 
    MODIFY COLUMN nationality VARCHAR(200) NOT NULL,
    MODIFY COLUMN visa_type VARCHAR(100) NOT NULL,
    MODIFY COLUMN visa_from DATE NOT NULL,
    MODIFY COLUMN visa_till DATE NOT NULL,
    MODIFY COLUMN visa_number VARCHAR(100) NOT NULL,
    MODIFY COLUMN visa_fee DECIMAL(10,2) NOT NULL;
-- +goose StatementEnd