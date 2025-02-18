-- +goose Up
-- +goose StatementBegin
CREATE TABLE HR_EMP_Expatriate(
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    expatriate BOOLEAN NOT NULL,
    nationality VARCHAR(200) NOT NULL,
    visa_type VARCHAR(100) NOT NULL,
    visa_from DATE NOT NULL,
    visa_till DATE NOT NULL,
    visa_number VARCHAR(100) NOT NULL,
    visa_fee DECIMAL(10,2) NOT NULL,
    visa_image_path TEXT NOT NULL,
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
DROP TABLE HR_EMP_Expatriate;
-- +goose StatementEnd
