-- +goose Up
-- +goose StatementBegin
CREATE TABLE HR_Employee(
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    gender VARCHAR(20) NOT NULL,
    dob DATE NOT NULL,
    religion VARCHAR(100) NOT NULL,
    primary_number VARCHAR(20) NOT NULL,
    secondary_number VARCHAR(20) NOT NULL,
    passport_id VARCHAR(50) NOT NULL,
    nationality VARCHAR(100) NOT NULL,
    passport_valid_till DATE NOT NULL,
    nic VARCHAR(20) NOT NULL,
    country VARCHAR(100) NOT NULL,
    nic_valid_till DATE NOT NULL,
    address TEXT NOT NULL,
    current_country VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL,
    updated_by BIGINT,
    FOREIGN KEY (updated_by) REFERENCES HR_Admin(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE HR_Employee;
-- +goose StatementEnd
