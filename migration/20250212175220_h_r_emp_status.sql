-- +goose Up
-- +goose StatementBegin
CREATE TABLE HR_EMP_Status(
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    status VARCHAR(50) NOT NULL,
    department VARCHAR(100) NOT NULL,
    designation VARCHAR(100) NOT NULL,
    valid_from DATE NOT NULL,
    valid_till DATE NOT NULL,
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
DROP TABLE HR_EMP_Status;
-- +goose StatementEnd
