-- +goose Up
-- +goose StatementBegin
CREATE TABLE HR_EMP_Bank_Details(
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    bank_name VARCHAR(200) NOT NULL,
    branch_name VARCHAR(200) NOT NULL,
    account_number VARCHAR(200) NOT NULL,
    account_holder VARCHAR(200) NOT NULL,
    updated_by BIGINT,
    employee_id BIGINT NOT NULL,
    FOREIGN KEY (employee_id) REFERENCES HR_Employee(id),
    FOREIGN KEY (updated_by) REFERENCES HR_Admin(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE HR_EMP_Bank_Details;
-- +goose StatementEnd
