-- +goose Up
-- +goose StatementBegin
CREATE TABLE HR_EMP_Allowances(
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(200) NOT NULL,
    amount DECIMAL(10,2) NOT NULL,
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
DROP TABLE HR_EMP_Allowances;
-- +goose StatementEnd
